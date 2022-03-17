package greenhouse

import (
	"context"
	"encoding/json"
	"fmt"
)

func GetJobPermissions(c *Client, ctx context.Context, id int) (*[]UserPermission, error) {
	var obj []UserPermission
	endpoint := fmt.Sprintf("v1/users/%d/permissions/jobs", id)
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func DeleteJobPermission(c *Client, ctx context.Context, jobId int, permId int) error {
	endpoint := fmt.Sprintf("v1/users/%d/permissions/jobs", jobId)
	jsonBody, err := json.Marshal(map[string]int{"job_permission_id": permId})
	if err != nil {
		return err
	}
	return Delete(c, ctx, endpoint, jsonBody)
}

func CreateJobPermission(c *Client, ctx context.Context, id int, obj *UserPermission) (int, error) {
	endpoint := fmt.Sprintf("v1/users/%d/permissions/jobs", id)
	return Create(c, ctx, endpoint, obj)
}

func GetFutureJobPermission(c *Client, ctx context.Context, id int) (*[]FutureJobPermission, error) {
	var obj []FutureJobPermission
	endpoint := fmt.Sprintf("v1/users/%d/permissions/future_jobs", id)
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func DeleteFutureJobPermission(c *Client, ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("v1/users/%d/permissions/future_jobs", id)
	return Delete(c, ctx, endpoint, nil)
}

func CreateFutureJobPermission(c *Client, ctx context.Context, id int, obj *FutureJobPermission) (int, error) {
	endpoint := fmt.Sprintf("v1/users/%d/permissions/future_jobs", id)
	return Create(c, ctx, endpoint, obj)
}
