package main

type LanguageService interface {
	GetLang(code string) Language
}

type FakeLanguageService struct{}

func (s *FakeLanguageService) GetLang(code string) Language {

	pronouns := []Pronoun{
		{
			Id: 1,
			Identifier: "je",
			DisplayName: "je",
			Reflexive: "je me",
			Order: 0,
		},
		{
			Id: 2,
			Identifier: "tu",
			DisplayName: "tu",
			Reflexive: "tu te",
			Order: 1,
		},
		{
			Id: 3,
			Identifier: "il",
			DisplayName: "il/elle/on",
			Reflexive: "il/elle/on se",
			Order: 2,
		},
		{
			Id: 4,
			Identifier: "nous",
			DisplayName: "nous",
			Reflexive: "nous nous",
			Order: 3,
		},
		{
			Id: 5,
			Identifier: "vous",
			DisplayName: "vous",
			Reflexive: "vous vous",
			Order: 4,
		},
		{
			Id: 6,
			Identifier: "ils",
			DisplayName: "ils/elles",
			Reflexive: "ils/elles se",
			Order: 5,
		},
	}

	tenses := []Tense{
		{
			Id: 1,
			Identifier: "present",
			DisplayName: "present",
			Order: 0,
		},
	}

	conjugations := []Conjugation{
		{
			TenseID: 1,
			PronounID: 1,
			Conjugation: "joue",
			NormalisedConjugation: "joue",
		},
	}

	verbs := []Verb{
		{
			Id: 1,
			Infinitive: "jouer",
			NormalisedInfinitive: "jouer",
			English: "to play",
			IsReflexive: false,
			IsHelper: false,
			HelperID: 5,
			Conjugations: ConjugationContainer{Data: conjugations},
		},
	}

	return Language{
		Id: 2,
		Lang: "Test Lang",
		Code: "tl",
		Locale: "tl-Tl",
		Pronouns: struct{ Data []Pronoun }{Data: pronouns},
		Tenses: struct{ Data []Tense }{Data: tenses},
		Verbs: struct{ Data []Verb }{Data: verbs},
	}
}