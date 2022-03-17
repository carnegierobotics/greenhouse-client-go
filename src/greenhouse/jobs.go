package greenhouse

import (
	"context"
	"fmt"
)

func GetAllJobs(c *Client) (*[]Job, error) {
	var obj []Job
	err := MultiGet(c, context.TODO(), "v1/jobs", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetJob(c *Client, id int) (*Job, error) {
	var obj Job
	err := SingleGet(c, context.TODO(), fmt.Sprintf("v1/jobs/%d", id), &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateJob(c *Client, obj *JobCreateInfo) (int, error) {
	return Create(c, context.TODO(), "v1/jobs", obj)
}

func UpdateJob(c *Client, id int, obj *JobUpdateInfo) error {
	return Update(c, context.TODO(), fmt.Sprintf("v1/jobs/%d", id), obj)
}
