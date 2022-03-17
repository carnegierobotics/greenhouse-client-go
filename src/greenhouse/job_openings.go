package greenhouse

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

func GetAllJobOpenings(c *Client) (*[]JobOpening, error) {
	return nil, errors.New("GetAllJobOpenings not implemented.")
}

func GetJobOpening(c *Client, jobId int, openingId int) (*JobOpening, error) {
	return nil, errors.New("GetJobOpening not implemented.")
}

func DeleteJobOpenings(c *Client, jobId int, OpeningIds []int) error {
	jsonBody, err := json.Marshal(map[string][]int{"ids": OpeningIds})
	if err != nil {
		return err
	}
	return Delete(c, context.TODO(), fmt.Sprintf("v1/jobs/%d/openings", jobId), jsonBody)
}

func UpdateJobOpenings(c *Client, jobId int, obj JobOpeningUpdateInfo) error {
	return errors.New("UpdateJobOpenings not implemented.")
}

func CreateJobOpenings(c *Client, obj JobOpeningCreateInfo) error {
	return errors.New("CreateJobOpenings not implemented.")
}
