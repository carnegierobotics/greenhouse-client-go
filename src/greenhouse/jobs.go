package greenhouse

import ()

type JobSpec struct {
	TemplateJobId  int      `json:"template_job_id"`
	NumberOpenings int      `json:"number_of_openings"`
	JobPostName    string   `json:"job_post_name,omitempty"`
	JobName        string   `json:"job_name,omitempty"`
	DepartmentId   int      `json:"department_id,omitempty"`
	OfficeIds      []int    `json:"office_ids,omitempty"`
	RequisitionId  string   `json:"requisition_id,omitempty"`
	OpeningIds     []string `json:"opening_ids"`
	/* Not in our product tier
	ExternalDepartmentId string `json:"external_department_id"`
	ExternalOfficeIds []string `json:"external_office_ids"`
	*/
}

type Job struct {
	Id                int                       `json:"id"`
	Name              string                    `json:"name"`
	RequisitionId     string                    `json:"requisition_id"`
	Notes             string                    `json:"notes"`
	Confidential      bool                      `json:"confidential"`
	Status            string                    `json:"status"`
	CreatedAt         string                    `json:"created_at"`
	OpenedAt          string                    `json:"opened_at"`
	ClosedAt          string                    `json:"closed_at"`
	UpdatedAt         string                    `json:"updated_at"`
	IsTemplate        bool                      `json:"is_template"`
	CopiedFromId      int                       `json:"copied_from_id"`
	Departments       []Department              `json:"departments"`
	Offices           []Office                  `json:"offices"`
	CustomFields      map[string]interface{}    `json:"custom_fields"`
	KeyedCustomFields map[string]interface{}    `json:"keyed_custom_fields"`
	HiringTeam        map[string][]HiringMember `json:"hiring_team"`
	Openings          []JobOpening              `json:"openings"`
}

type JobCreateInfo struct {
}

type JobUpdateInfo struct {
}

type HiringMember struct {
	UserBasics
	Responsible bool `json:"responsible"`
}
