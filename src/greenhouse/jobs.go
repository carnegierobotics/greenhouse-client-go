package greenhouse

import (
	"context"
	"fmt"
)

func GetAllJobs(c *Client, ctx context.Context) (*[]Job, error) {
	var obj []Job
	err := MultiGet(c, ctx, "v1/jobs", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetJob(c *Client, ctx context.Context, id int) (*Job, error) {
	var obj Job
	err := SingleGet(c, ctx, fmt.Sprintf("v1/jobs/%d", id), &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateJob(c *Client, ctx context.Context, obj *JobCreateInfo) (int, error) {
	return Create(c, ctx, "v1/jobs", obj)
}

func UpdateJob(c *Client, ctx context.Context, id int, obj *JobUpdateInfo) error {
	return Update(c, ctx, fmt.Sprintf("v1/jobs/%d", id), obj)
}
