package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

type RealLanguageService struct{
	db *sql.DB
	configService ConfigService
	config *DBConfig
}

func (s *RealLanguageService) DB() *sql.DB {
	if s.db == nil {
		if s.config == nil {
			if s.configService == nil {
				s.configService = &(FileConfigService{File: "config"})
			}
			s.config = s.configService.GetConfig()
		}

		db, err := sql.Open(s.config.Driver, s.config.DBString())

		if err != nil {
			panic(err)
		}

		s.db = db
	}
	return s.db
}

func (s *RealLanguageService) GetLangIdFromCode(code string) (int, error) {
	var err error

	rows, err := s.DB().Query("SELECT id FROM languages WHERE code = ?", code)
	defer rows.Close()
	if err != nil {
		return 0, err
	}

	id := 0

	if rows.Next() {
		err = rows.Scan(&id)
	} else {
		err = fmt.Errorf("language not found with code '%s'", code)
	}

	return id, err
}

func (s *RealLanguageService) GetVerbsSince(code string, since int) (int, VerbContainer, error) {
	id, err := s.GetLangIdFromCode(code)
	if err != nil {
		return 0, VerbContainer{}, err
	}

	verbs, err := s.getVerbsAndConjugationsSince(id, since)
	if err != nil {
		return id, VerbContainer{}, err
	}

	return id, VerbContainer{Data: verbs}, nil
}

func (s *RealLanguageService) GetVerbsOnly(code string) (int, VerbContainer, error) {
	id, err := s.GetLangIdFromCode(code)
	if err != nil {
		return 0, VerbContainer{}, err
	}

	verbs, err := s.getVerbsAndConjugations(id)
	if err != nil {
		return id, VerbContainer{}, err
	}

	return id, VerbContainer{Data: verbs}, nil
}

func (s *RealLanguageService) GetLang(code string) (Language, error) {
	language := Language{}

	rows, err := s.DB().Query(`
SELECT l.id, l.lang, l.` + "`code`" + `, l.locale, UNIX_TIMESTAMP(max(v.updated_at)) version, UNIX_TIMESTAMP(GREATEST(max(t.updated_at), max(p.updated_at))) schemaVersion, hasReflexives, hasHelpers
FROM languages l, verbs v, tenses t, pronouns p
WHERE ` + "`code`" + ` = ?
      AND v.lang_id = l.id
      AND t.lang_id = l.id
      AND p.lang_id = l.id
GROUP BY l.id`, code)
	if err != nil {
		return language, err
	}

	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(
			&language.Id,
			&language.Lang,
			&language.Code,
			&language.Locale,
			&language.Version,
			&language.SchemaVersion,
			&language.HasReflexives,
			&language.HasHelpers); err != nil {
			return language, err
		}
	} else {
		return language, fmt.Errorf("language not found with code '%s'", code)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	tenses, err := s.getTenses(language.Id)
	if err != nil {
		return language, err
	}
	language.Tenses = struct{Data []Tense}{Data: tenses}

	pronouns, err := s.getPronouns(language.Id)
	if err != nil {
		return language, err
	}
	language.Pronouns = struct{Data []Pronoun}{Data: pronouns}

	verbs, err := s.getVerbsAndConjugations(language.Id)
	if err != nil {
		return language, err
	}
	language.Verbs = struct{Data []Verb}{Data: verbs}

	return language, nil
}

func (s *RealLanguageService) getTenses(langId int) ([]Tense, error) {
	var tenses []Tense

	rows, err := s.DB().Query("SELECT id, identifier, displayName, `order` FROM tenses WHERE lang_id = ?", langId)
	if err != nil {
		return tenses, err
	}

	defer rows.Close()

	for rows.Next() {
		tense := Tense{}
		if err := rows.Scan(&tense.Id, &tense.Identifier, &tense.DisplayName, &tense.Order); err != nil {
			return []Tense{}, err
		} else {
			tenses = append(tenses, tense)
		}
	}

	return tenses, nil
}

func (s *RealLanguageService) getPronouns(langId int) ([]Pronoun, error) {
	var pronouns []Pronoun

	rows, err := s.DB().Query("SELECT id, identifier, displayName, `order`, reflexive FROM pronouns WHERE lang_id = ?", langId)
	if err != nil {
		return pronouns, err
	}

	defer rows.Close()

	for rows.Next() {
		pronoun := Pronoun{}
		if err := rows.Scan(
			&pronoun.Id,
			&pronoun.Identifier,
			&pronoun.DisplayName,
			&pronoun.Order,
			&pronoun.Reflexive); err != nil {
			return []Pronoun{}, err
		} else {
			pronouns = append(pronouns, pronoun)
		}
	}

	return pronouns, nil
}

func (s *RealLanguageService) scanVerbs(rows *sql.Rows) ([]Verb, error) {
	var verbs []Verb

	for rows.Next() {
		verb := Verb{}
		if err := rows.Scan(
			&verb.Id,
			&verb.Infinitive,
			&verb.NormalisedInfinitive,
			&verb.English,
			&verb.HelperID,
			&verb.IsHelper,
			&verb.IsReflexive); err != nil {
			return []Verb{}, err
		} else {
			verbs = append(verbs, verb)
		}
	}

	return verbs, nil
}

func (s *RealLanguageService) getVerbs(langId int) ([]Verb, error) {
	rows, err := s.DB().Query("SELECT id, infinitive, normalisedInfinitive, english, helperID, isHelper, isReflexive FROM verbs WHERE lang_id = ?", langId)
	if err != nil {
		return []Verb{}, err
	}
	defer rows.Close()

	return s.scanVerbs(rows)
}

func (s *RealLanguageService) getVerbsSince(langId int, since int) ([]Verb, error) {
	rows, err := s.DB().Query(
		"SELECT v.id, v.infinitive, v.normalisedInfinitive, v.english, v.helperID, v.isHelper, v.isReflexive FROM verbs AS v, conjugations AS c " +
			"WHERE v.lang_id = ? " +
			"AND c.verb_id = v.id " +
			"AND GREATEST(UNIX_TIMESTAMP(v.updated_at), UNIX_TIMESTAMP(c.updated_at)) > ? ",
		langId,
		since)
	if err != nil {
		return []Verb{}, err
	}
	defer rows.Close()

	return s.scanVerbs(rows)
}

func (s *RealLanguageService) getConjugations(verbId int) ([]Conjugation, error) {
	var conjs []Conjugation

	rows, err := s.DB().Query("SELECT conjugation, normalisedConjugation, pronoun_id, tense_id FROM conjugations WHERE verb_id = ?", verbId)
	if err != nil {
		return conjs, err
	}

	for rows.Next() {
		conj := Conjugation{}
		if err := rows.Scan(
			&conj.Conjugation,
			&conj.NormalisedConjugation,
			&conj.PronounID,
			&conj.TenseID); err != nil {
			return []Conjugation{}, err
		} else {
			conjs = append(conjs, conj)
		}
	}

	if len(conjs) == 0 {
		err = fmt.Errorf("no conjugations found for verb %d", verbId)
	}

	return conjs, err
}

func (s *RealLanguageService) getVerbsAndConjugations(langId int) ([]Verb, error) {
	verbs, err := s.getVerbs(langId)
	if err != nil {
		return []Verb{},err
	}

	for i := range verbs {
		verb := verbs[i]
		conjs, err := s.getConjugations(verb.Id)

		if err != nil {
			return []Verb{}, err
		} else {
			verbs[i].Conjugations = ConjugationContainer{Data: conjs}
		}
	}
	return verbs, nil
}

func (s *RealLanguageService) getVerbsAndConjugationsSince(langId int, since int) ([]Verb, error) {
	verbs, err := s.getVerbsSince(langId, since)
	if err != nil {
		return []Verb{},err
	}

	for i, verb := range verbs {
		conjs, err := s.getConjugations(verb.Id)
		if err != nil {
			return []Verb{}, err
		} else {
			verbs[i].Conjugations = ConjugationContainer{Data: conjs}
		}
	}
	return verbs, nil
}