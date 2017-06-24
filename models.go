package main

type Language struct {
	Id int
	Lang string
	Code string
	Locale string
	Version int
	SchemaVersion int
	HasReflexives bool
	HasHelpers bool

	Tenses struct{
		Data []Tense
	}
	Pronouns struct{
		Data []Pronoun
	}
	Verbs struct {
		Data []Verb
	}
}

func (l *Language) Data() (struct{Data Language}) {
	return struct{Data Language}{Data: *l}
}

type Tense struct {
	Id int
	Identifier string
	DisplayName string
	Order int
}

type Pronoun struct {
	Id int
	Identifier string
	DisplayName string
	Reflexive string
	Order int
}

type Verb struct {
	Id int
	Infinitive string
	NormalisedInfinitive string
	English string
	HelperID int
	IsHelper bool
	IsReflexive bool
	Conjugations struct {
		Data []Conjugation
	}
}

type Conjugation struct {
	Conjugation string
	NormalisedConjugation string
	TenseID int
	PronounID int
}