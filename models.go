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
	Id					 int	`json:"id"`
	Infinitive			 string	`json:"i"`
	NormalisedInfinitive string	`json:"ni,omitempty"`
	English 		     string `json:"e"`
	HelperID 		     int    `json:"hid,omitempty"`
	IsHelper 		     bool   `json:"ih"`
	IsReflexive 		 bool   `json:"ir"`
	Conjugations 		 ConjugationContainer `json:"conjugations"`
}

type Conjugation struct {
	Conjugation 		    string	`json:"c"`
	NormalisedConjugation	string	`json:"nc"`
	TenseID 			    int		`json:"tid"`
	PronounID 			    int		`json:"pid"`
}

type VerbContainer struct {
	Data []Verb
}

type ConjugationContainer struct {
	Data []Conjugation
}