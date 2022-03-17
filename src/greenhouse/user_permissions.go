package greenhouse

import (
	"context"
	"encoding/json"
	"fmt"
)

func GetJobPermissions(c *Client, id int) (*[]UserPermission, error) {
	var obj []UserPermission
	endpoint := fmt.Sprintf("v1/users/%d/permissions/jobs", id)
	err := MultiGet(c, context.TODO(), endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func DeleteJobPermission(c *Client, jobId int, permId int) error {
	endpoint := fmt.Sprintf("v1/users/%d/permissions/jobs", jobId)
	jsonBody, err := json.Marshal(map[string]int{"job_permission_id": permId})
	if err != nil {
		return err
	}
	return Delete(c, context.TODO(), endpoint, jsonBody)
}

func CreateJobPermission(c *Client, id int, obj *UserPermission) (int, error) {
	endpoint := fmt.Sprintf("v1/users/%d/permissions/jobs", id)
	return Create(c, context.TODO(), endpoint, obj)
}

func GetFutureJobPermission(c *Client, id int) (*[]FutureJobPermission, error) {
	var obj []FutureJobPermission
	endpoint := fmt.Sprintf("v1/users/%d/permissions/future_jobs", id)
	err := MultiGet(c, context.TODO(), endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func DeleteFutureJobPermission(c *Client, id int) error {
	endpoint := fmt.Sprintf("v1/users/%d/permissions/future_jobs", id)
	return Delete(c, context.TODO(), endpoint, nil)
}

func CreateFutureJobPermission(c *Client, id int, obj *FutureJobPermission) (int, error) {
	endpoint := fmt.Sprintf("v1/users/%d/permissions/future_jobs", id)
	return Create(c, context.TODO(), endpoint, obj)
}
