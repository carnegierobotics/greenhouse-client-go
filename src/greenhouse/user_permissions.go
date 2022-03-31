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

func GetJobPermission(c *Client, ctx context.Context, userId int, permId int) (*UserPermission, error) {
  list, err := GetJobPermissions(c, ctx, userId)
  if err != nil {
    return nil, err
  }
  for _, item := range *list {
    if item.Id == permId {
      return &item, nil
    }
  }
  return nil, nil
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

func GetFutureJobPermissions(c *Client, ctx context.Context, id int) (*[]FutureJobPermission, error) {
	var obj []FutureJobPermission
	endpoint := fmt.Sprintf("v1/users/%d/permissions/future_jobs", id)
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetFutureJobPermission(c *Client, ctx context.Context, userId int, permId int) (*FutureJobPermission, error) {
  list, err := GetFutureJobPermissions(c, ctx, userId)
  if err != nil {
    return nil, err
  }
  for _, item := range *list {
    if item.Id == permId {
      return &item, nil
    }
  }
  return nil, nil
}

func DeleteFutureJobPermission(c *Client, ctx context.Context, id int, permId int) error {
	endpoint := fmt.Sprintf("v1/users/%d/permissions/future_jobs", id)
  jsonBody, err := json.Marshal(map[string]int{"future_job_permission_id": permId})
  if err != nil {
    return err
  }
	return Delete(c, ctx, endpoint, jsonBody)
}

func CreateFutureJobPermission(c *Client, ctx context.Context, id int, obj *FutureJobPermission) (int, error) {
	endpoint := fmt.Sprintf("v1/users/%d/permissions/future_jobs", id)
	return Create(c, ctx, endpoint, obj)
}
