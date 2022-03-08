package greenhouse

import (
	"context"
	"encoding/json"
	"fmt"
)

type Job struct {
	Id                int                               `json:"id"`
	Name              string                            `json:"name"`
	RequisitionId     string                            `json:"requisition_id"`
	Notes             string                            `json:"notes"`
	Confidential      bool                              `json:"confidential"`
	Status            string                            `json:"status"`
	CreatedAt         string                            `json:"created_at"`
	OpenedAt          string                            `json:"opened_at"`
	ClosedAt          string                            `json:"closed_at"`
	UpdatedAt         string                            `json:"updated_at"`
	IsTemplate        bool                              `json:"is_template"`
	CopiedFromId      int                               `json:"copied_from_id"`
	Departments       []Department                      `json:"departments"`
	Offices           []Office                          `json:"offices"`
	CustomFields      map[string]interface{}            `json:"custom_fields"`
	KeyedCustomFields map[string]map[string]interface{} `json:"keyed_custom_fields"`
	HiringTeam        map[string][]HiringMember         `json:"hiring_team"`
	Openings          []JobOpening                      `json:"openings"`
}

type JobCreateInfo struct {
	TemplateJobId  int      `json:"template_job_id"`
	NumberOpenings int      `json:"number_of_openings"`
	JobPostName    string   `json:"job_post_name,omitempty"`
	JobName        string   `json:"job_name,omitempty"`
	DepartmentId   int      `json:"department_id,omitempty"`
	OfficeIds      []int    `json:"office_ids,omitempty"`
	RequisitionId  string   `json:"requisition_id,omitempty"`
	OpeningIds     []string `json:"opening_ids,omitempty"`
	/* Not in our product tier
	ExternalDepartmentId string `json:"external_department_id"`
	ExternalOfficeIds []string `json:"external_office_ids"`
	*/
}

type JobUpdateInfo struct {
	Name                     string        `json:"name,omitempty"`
	Notes                    string        `json:"notes,omitempty"`
	Anywhere                 bool          `json:"anywhere,omitempty"`
	RequisitionId            string        `json:"requisition_id,omitempty"`
	TeamsandResponsibilities string        `json:"teams_and_responsibilities,omitempty"`
	HowToSellThisJob         string        `json:"how_to_sell_this_job,omitempty"`
	CustomFields             []CustomField `json:"custom_fields,omitempty"`
	OfficeIds                []int         `json:"office_ids,omitempty"`
	DepartmentId             int           `json:"department_id,omitempty"`
	/* Not in our product tier
	ExternalOfficeIds []string `json:"external_office_ids"`
	ExternalDepartmentId string `json:"external_department_id"`
	*/
}

type HiringMember struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	Active      bool   `json:"active"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Name        string `json:"name"`
	Responsible bool   `json:"responsible"`
	EmployeeId  string `json:"employee_id"`
}

type HiringMemberUpdateInfo struct {
	UserId                           int  `json:"user_id"`
	ResponsibleForFutureCandidates   bool `json:"responsible_for_future_candidates"`
	ResponsibleForActiveCandidates   bool `json:"responsible_for_active_candidates"`
	ResponsibleForInactiveCandidates bool `json:"responsible_for_inactive_candidates"`
}

func GetJob(c *Client, id int) (*Job, error) {
	var obj Job
	err := GetById(c, "jobs", id, &obj, context.TODO())
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateJob(c *Client, obj *JobCreateInfo) (int, error) {
	id, err := Create(c, "jobs", obj, context.TODO())
	if err != nil {
		return id, err
	}
	return id, nil
}

func UpdateJob(c *Client, id int, obj *JobUpdateInfo) error {
	err := Update(c, "jobs", id, obj, context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func UpdateJobHiringTeam(c *Client, id int, obj *map[string][]HiringMemberUpdateInfo, ctx context.Context) error {
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	resp, err := c.Client.R().SetContext(ctx).SetBody(jsonBody).Put(fmt.Sprintf("v1/jobs/%d", id))
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("Got %s for URL: %s", resp.Status(), resp.Request.URL)
	}
	return nil
}
