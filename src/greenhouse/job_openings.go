package greenhouse

import (
	"context"
	"encoding/json"
	"fmt"
)

func GetAllJobOpenings(c *Client, ctx context.Context, id int, status string) (*[]JobOpening, error) {
	var obj []JobOpening
	endpoint := fmt.Sprintf("v1/jobs/%d/openings", id)
	querystring := fmt.Sprintf("status=%s", status)
	err := MultiGet(c, ctx, endpoint, querystring, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetJobOpening(c *Client, ctx context.Context, jobId int, openingId int) (*JobOpening, error) {
	var obj JobOpening
	endpoint := fmt.Sprintf("v1/jobs/%d/openings", jobId)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func DeleteJobOpenings(c *Client, ctx context.Context, jobId int, openingIds []int) error {
	jsonBody, err := json.Marshal(map[string][]int{"ids": openingIds})
	if err != nil {
		return err
	}
	return Delete(c, ctx, fmt.Sprintf("v1/jobs/%d/openings", jobId), jsonBody)
}

func UpdateJobOpenings(c *Client, ctx context.Context, jobId int, openingId int, obj *JobOpeningUpdateInfo) error {
	endpoint := fmt.Sprintf("v1/jobs/%d/openings/%d", jobId, openingId)
	return Update(c, ctx, endpoint, obj)
}

func CreateJobOpenings(c *Client, ctx context.Context, jobId int, obj JobOpeningCreateInfo) (int, error) {
	endpoint := fmt.Sprintf("v1/jobs/%d/openings", jobId)
	return Create(c, ctx, endpoint, obj)
}
