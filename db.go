package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

type RealLanguageService struct{}

func (s *RealLanguageService) GetVerbsOnly(code string) (int, VerbContainer, error) {
	db, err := sql.Open("mysql", "root:ufx366@tcp(localhost:3306)/reed.brad_iVerbs")
	if err != nil {
		return 0, VerbContainer{}, err
	}

	rows, err := db.Query("SELECT id FROM languages WHERE code = ?", code)
	if err != nil {
		return 0, VerbContainer{}, err
	}

	defer rows.Close()

	id := 0

	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, VerbContainer{}, err
		}
	}

	verbs, err := s.getVerbsAndConjugations(db, id)
	if err != nil {
		return id, VerbContainer{}, err
	}

	return id, VerbContainer{Data: verbs}, nil
}

func (s *RealLanguageService) GetLang(code string) (Language, error) {
	language := Language{}

	db, err := sql.Open("mysql", "root:ufx366@tcp(localhost:3306)/reed.brad_iVerbs")
	if err != nil {
		return language, err
	}

	rows, err := db.Query(`
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

	for rows.Next() {
		if err := rows.Scan(
			&language.Id,
			&language.Lang,
			&language.Code,
			&language.Locale,
			&language.Version,
			&language.SchemaVersion,
			&language.HasReflexives,
			&language.HasHelpers); err != nil {
			fmt.Print(err)
		}
	}

	if language.Id == 0 {
		// Language was not found
		return language, fmt.Errorf("Language not found with code '%s'", code)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Fetch schema
	tenses, err := s.getTenses(db, language.Id)
	if err != nil {
		return language, err
	}
	language.Tenses = struct{Data []Tense}{Data: tenses}

	pronouns, err := s.getPronouns(db, language.Id)
	if err != nil {
		return language, err
	}
	language.Pronouns = struct{Data []Pronoun}{Data: pronouns}

	verbs, err := s.getVerbsAndConjugations(db, language.Id)
	if err != nil {
		return language, err
	}
	language.Verbs = struct{Data []Verb}{Data: verbs}

	return language, nil
}

func (s *RealLanguageService) getTenses(db *sql.DB, langId int) ([]Tense, error) {
	tenses := []Tense{}

	rows, err := db.Query("SELECT id, identifier, displayName, `order` FROM tenses WHERE lang_id = ?", langId)
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

func (s *RealLanguageService) getPronouns(db *sql.DB, langId int) ([]Pronoun, error) {
	pronouns := []Pronoun{}

	rows, err := db.Query("SELECT id, identifier, displayName, `order`, reflexive FROM pronouns WHERE lang_id = ?", langId)
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

func (s *RealLanguageService) getVerbs(db *sql.DB, langId int) ([]Verb, error) {
	verbs := []Verb{}

	rows, err := db.Query("SELECT id, infinitive, normalisedInfinitive, english, helperID, isHelper, isReflexive FROM verbs WHERE lang_id = ?", langId)
	if err != nil {
		return verbs, err
	}

	defer rows.Close()

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

func (s *RealLanguageService) getConjugations(db *sql.DB, verbId int) ([]Conjugation, error) {
	conjs := []Conjugation{}

	rows, err := db.Query("SELECT conjugation, normalisedConjugation, pronoun_id, tense_id FROM conjugations WHERE verb_id = ?", verbId)
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
		// Conjugations were not found
		return conjs, fmt.Errorf("No conjugations found for verb %d", verbId)
	}

	return conjs, nil
}

func (s *RealLanguageService) getVerbsAndConjugations(db *sql.DB, langId int) ([]Verb, error) {
	verbs, err := s.getVerbs(db, langId)
	if err != nil {
		return []Verb{},err
	}

	for i := range verbs {
		verb := verbs[i]
		conjs, err := s.getConjugations(db, verb.Id)
		if err != nil {
			return []Verb{}, err
		} else {
			verbs[i].Conjugations = struct{Data []Conjugation}{Data: conjs}
		}
	}
	return verbs, nil
}