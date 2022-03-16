package greenhouse

import (
	"errors"
)

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
