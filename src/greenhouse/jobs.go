package greenhouse

import (
	"context"
)

func GetAllJobs(c *Client) (*[]Job, error) {
	var obj []Job
	err := GetAll(c, "jobs", &obj, context.TODO())
	if err != nil {
		return nil, err
	}
	return &obj, nil
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
