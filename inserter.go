package main

import (
	"database/sql"
	"fmt"
)

type Inserter struct {
	Db *sql.DB
}

func (s *Inserter) GetLangIds() (ids []int, err error) {
	rows, err := s.Db.Query("SELECT id FROM languages")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return
		}
		ids = append(ids)
	}
	return
}

func (s *Inserter) InsertLanguage(lang *Language) error {
	rows, err := s.Db.Query("SELECT COUNT(*) FROM languages WHERE id = ? OR code = ?", lang.Id, lang.Code)
	defer rows.Close()
	if err != nil {
		return err
	}

	rows.Next()

	var count int
	err = rows.Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("language already exists with ID %d or code '%s'", lang.Id, lang.Code)
	}

	_, err = s.Db.Exec("INSERT INTO languages (id, lang, code, locale, hasReflexives, hasHelpers, created_at) VALUES (?,?,?,?,?,?,NOW())",
		lang.Id,
		lang.Lang,
		lang.Code,
		lang.Locale,
		lang.HasReflexives,
		lang.HasHelpers,
	)
	if err != nil {
		return err
	}

	err = s.InsertTenses(lang.Id, &lang.Tenses.Data)
	if err != nil {
		return err
	}

	err = s.InsertPronouns(lang.Id, &lang.Pronouns.Data)
	if err != nil {
		return err
	}

	err = s.InsertVerbs(lang.Id, &lang.Verbs.Data)
	if err != nil {
		return err
	}

	return nil
}

func (s *Inserter) InsertTenses(langID int, tenses *[]Tense) error {
	tenseIds := ""

	for i, tense := range *tenses {
		if i == 0 {
			tenseIds = fmt.Sprintf("%d", tense.Id)
		} else {
			tenseIds = fmt.Sprintf("%d,%s", tense.Id, tenseIds)
		}
	}

	rows, err := s.Db.Query("SELECT id FROM tenses WHERE id IN (" + tenseIds + ")")
	defer rows.Close()
	if err != nil {
		return err
	}

	var existingTenseIds []int

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return err
		}

		existingTenseIds = append(existingTenseIds, id)
	}

	if len(existingTenseIds) > 0 {
		return fmt.Errorf("existing tenses found: %+v", existingTenseIds)
	}

	stmt, err := s.Db.Prepare("INSERT INTO tenses (id, lang_id, identifier, displayName, `order`, updated_at) VALUES (?,?,?,?,?,NOW())")
	if err != nil {
		return err
	}

	for _, tense := range *tenses {
		_, err = stmt.Exec(tense.Id, langID, tense.Identifier, tense.DisplayName, tense.Order)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Inserter) InsertPronouns(langID int, pronouns *[]Pronoun) error {
	pronounIds := ""

	for i, pronoun := range *pronouns {
		if i == 0 {
			pronounIds = fmt.Sprintf("%d", pronoun.Id)
		} else {
			pronounIds = fmt.Sprintf("%d,%s", pronoun.Id, pronounIds)
		}
	}

	rows, err := s.Db.Query("SELECT id FROM pronouns WHERE id IN (" + pronounIds + ")")
	defer rows.Close()
	if err != nil {
		return err
	}

	var existingPronounIds []int

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return err
		}

		existingPronounIds = append(existingPronounIds, id)
	}

	if len(existingPronounIds) > 0 {
		return fmt.Errorf("existing pronouns found: %+v", existingPronounIds)
	}

	stmt, err := s.Db.Prepare("INSERT INTO pronouns (id, lang_id, identifier, displayName, reflexive, `order`, updated_at) VALUES (?,?,?,?,?,?,NOW())")
	if err != nil {
		return err
	}

	for _, pronoun := range *pronouns {
		_, err = stmt.Exec(pronoun.Id, langID, pronoun.Identifier, pronoun.DisplayName, pronoun.Reflexive, pronoun.Order)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Inserter) InsertVerbs(langID int, verbs *[]Verb) error {
	verbsIds := ""

	for i, verb := range *verbs {
		if i == 0 {
			verbsIds = fmt.Sprintf("%d", verb.Id)
		} else {
			verbsIds = fmt.Sprintf("%d,%s", verb.Id, verbsIds)
		}
	}

	rows, err := s.Db.Query("SELECT id FROM verbs WHERE id IN (" + verbsIds + ")")
	defer rows.Close()
	if err != nil {
		return err
	}

	var existingVerbIds []int

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return err
		}

		existingVerbIds = append(existingVerbIds, id)
	}

	if len(existingVerbIds) > 0 {
		return fmt.Errorf("existing verbs found: %+v", existingVerbIds)
	}

	stmt, err := s.Db.Prepare("INSERT INTO verbs (id, lang_id, infinitive, normalisedInfinitive, english, helperID, isHelper, isReflexive, updated_at) VALUES (?,?,?,?,?,?,?,?,NOW())")
	if err != nil {
		return err
	}

	for _, verb := range *verbs {
		_, err = stmt.Exec(
			verb.Id,
			langID,
			verb.Infinitive,
			verb.NormalisedInfinitive,
			verb.English,
			verb.HelperID,
			verb.IsHelper,
			verb.IsReflexive,
		)
		if err != nil {
			return err
		}
		err = s.InsertConjugations(verb.Id, &verb.Conjugations.Data)
	}

	return nil
}

func (s *Inserter) InsertConjugations(verbID int, conjugations *[]Conjugation) error {
	rows, err := s.Db.Query("SELECT COUNT(*) FROM conjugations WHERE verb_id = ?", verbID)
	defer rows.Close()
	if err != nil {
		return err
	}

	var count int

	rows.Next()

	err = rows.Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("%d existing conjugations found for verb ID %d", count, verbID)
	}

	stmt, err := s.Db.Prepare("INSERT INTO conjugations (verb_id, tense_id, pronoun_id, conjugation, normalisedConjugation, updated_at) VALUES (?,?,?,?,?,NOW())")
	if err != nil {
		return err
	}

	for _, conjugation := range *conjugations {
		_, err = stmt.Exec(
			verbID,
			conjugation.TenseID,
			conjugation.PronounID,
			conjugation.Conjugation,
			conjugation.NormalisedConjugation,
		)
		if err != nil {
			return err
		}
	}

	return nil
}