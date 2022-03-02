package greenhouse

import ()

type JobPost struct {
	Id                       int                   `json:"id"`
	Active                   bool                  `json:"active"`
	Live                     bool                  `json:"live"`
	FirstPublishedAt         string                `json:"first_published_at"`
	Internal                 bool                  `json:"internal"`
	External                 bool                  `json:"external"`
	JobId                    int                   `json:"job_id"`
	Content                  string                `json:"content"`
	InternalContent          string                `json:"internal_content"`
	UpdatedAt                string                `json:"updated_at"`
	CreatedAt                string                `json:"created_at"`
	DemographicQuestionSetId int                   `json:"demographic_question_set_id"`
	Questions                []DemographicQuestion `json:"questions"`
}

type JobPostUpdateInfo struct {
}
