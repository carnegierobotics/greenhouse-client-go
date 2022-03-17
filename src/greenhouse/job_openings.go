package greenhouse

import (
	"context"
	"encoding/json"
	"fmt"
)

func GetAllJobOpenings(c *Client, id int, status string) (*[]JobOpening, error) {
	var obj []JobOpening
	endpoint := fmt.Sprintf("v1/jobs/%d/openings", id)
	querystring := fmt.Sprintf("status=%s", status)
	err := MultiGet(c, context.TODO(), endpoint, querystring, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetJobOpening(c *Client, jobId int, openingId int) (*JobOpening, error) {
	var obj JobOpening
	endpoint := fmt.Sprintf("v1/jobs/%d/openings", jobId)
	err := SingleGet(c, context.TODO(), endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func DeleteJobOpenings(c *Client, jobId int, openingIds []int) error {
	jsonBody, err := json.Marshal(map[string][]int{"ids": openingIds})
	if err != nil {
		return err
	}
	return Delete(c, context.TODO(), fmt.Sprintf("v1/jobs/%d/openings", jobId), jsonBody)
}

func UpdateJobOpenings(c *Client, jobId int, openingId int, obj *JobOpeningUpdateInfo) error {
	endpoint := fmt.Sprintf("v1/jobs/%d/openings/%d", jobId, openingId)
	return Update(c, context.TODO(), endpoint, obj)
}

func CreateJobOpenings(c *Client, jobId int, obj JobOpeningCreateInfo) (int, error) {
	endpoint := fmt.Sprintf("v1/jobs/%d/openings", jobId)
	return Create(c, context.TODO(), endpoint, obj)
}
