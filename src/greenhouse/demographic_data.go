package greenhouse

import ()

type DemographicQuestionSet struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

type DemographicQuestion struct {
	Id                       int           `json:"id"`
	DemographicQuestionSetId int           `json:"demographic_question_set_id"`
	Name                     string        `json:"name"`
	Translations             []Translation `json:"translations"`
	Active                   bool          `json:"active"`
	Required                 bool          `json:"required"`
	AnswerType               string        `json:"answer_type"`
}

type Translation struct {
	Language string `json:"language"`
	Name     string `json:"name"`
}

type DemographicAnswerOption struct {
	Id                    int           `json:"id"`
	Name                  string        `json:"name"`
	Translations          []Translation `json:"translations"`
	Active                bool          `json:"active"`
	FreeForm              bool          `json:"free_form"`
	DemographicQuestionId int           `json:"demographic_question_id"`
}
