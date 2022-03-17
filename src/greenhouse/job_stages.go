package greenhouse

import (
	"context"
	"fmt"
)

func GetAllJobStages(c *Client, ctx context.Context) (*[]JobStage, error) {
	var obj []JobStage
	err := MultiGet(c, ctx, "v1/job_stages", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetJobStagesForJob(c *Client, ctx context.Context, id int) (*[]JobStage, error) {
	var obj []JobStage
	endpoint := fmt.Sprintf("v1/jobs/%d/stages", id)
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetJobStage(c *Client, ctx context.Context, id int) (*JobStage, error) {
	var obj JobStage
	endpoint := fmt.Sprintf("v1/job_stages/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
