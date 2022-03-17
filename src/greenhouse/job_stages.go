package greenhouse

import (
	"context"
	"fmt"
)

func GetAllJobStages(c *Client) (*[]JobStage, error) {
	var obj []JobStage
	err := MultiGet(c, context.TODO(), "v1/job_stages", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetJobStagesForJob(c *Client, id int) (*[]JobStage, error) {
	var obj []JobStage
	endpoint := fmt.Sprintf("v1/jobs/%d/stages", id)
	err := MultiGet(c, context.TODO(), endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetJobStage(c *Client, id int) (*JobStage, error) {
	var obj JobStage
	endpoint := fmt.Sprintf("v1/job_stages/%d", id)
	err := SingleGet(c, context.TODO(), endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
