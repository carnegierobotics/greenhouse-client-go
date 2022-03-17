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
