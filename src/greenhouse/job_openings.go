package greenhouse

import (
	"errors"
)

type JobOpening struct {
	Id            int               `json:"id"`
	OpeningId     string            `json:"opening_id"`
	Status        string            `json:"status"`
	OpenedAt      string            `json:"opened_at"`
	ClosedAt      string            `json:"closed_at"`
	ApplicationId int               `json:"application_id"`
	CloseReason   CloseReason       `json:"close_reason"`
	CustomFields  map[string]string `json:"custom_fields"`
}

type JobOpeningCreateInfo struct {
	Openings []Opening `json:"openings"`
}

type Opening struct {
	Id           int                 `json:"opening_id"`
	CustomFields []map[string]string `json:"custom_fields"`
}

type JobOpeningUpdateInfo struct {
	Id            int                 `json:"opening_id,omitempty"`
	Status        string              `json:"status,omitempty"`
	CloseReasonId int                 `json:"close_reason_id,omitempty"`
	CustomFields  []map[string]string `json:"custom_fields,omitempty"`
}

type JobOpeningDeleteInfo struct {
	Ids []int `json:"ids"`
}

func GetAllJobOpenings(c *Client) (*[]JobOpening, error) {
	return nil, errors.New("GetAllJobOpenings not implemented.")
}

func GetJobOpening(c *Client, jobId int, openingId int) (*JobOpening, error) {
	return nil, errors.New("GetJobOpening not implemented.")
}

func DeleteJobOpenings(c *Client, obj JobOpeningDeleteInfo) error {
	return errors.New("DeleteJobOpenings not implemented.")
}

func UpdateJobOpenings(c *Client, jobId int, obj JobOpeningUpdateInfo) error {
	return errors.New("UpdateJobOpenings not implemented.")
}

func CreateJobOpenings(c *Client, obj JobOpeningCreateInfo) error {
	return errors.New("CreateJobOpenings not implemented.")
}
