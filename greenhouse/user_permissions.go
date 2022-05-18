/*
Copyright 2021-2022
Carnegie Robotics, LLC
4501 Hatfield Street, Pittsburgh, PA 15201
https://www.carnegierobotics.com
All rights reserved.

This file is part of greenhouse-client-go.

greenhouse-client-go is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

greenhouse-client-go is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with greenhouse-client-go. If not, see <https://www.gnu.org/licenses/>.
*/

package greenhouse

import (
	"context"
	"encoding/json"
	"fmt"
)

// GetJobPermissions retrieves a list of all permissions for a user.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-job-permissions
func GetJobPermissions(c *Client, ctx context.Context, id int) (*[]UserPermission, error) {
	var obj []UserPermission
	endpoint := fmt.Sprintf("v1/users/%d/permissions/jobs", id)
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetJobPermission retrieves a single permission by ID.
func GetJobPermission(c *Client, ctx context.Context, userId int, permId int) (*UserPermission, error) {
	list, err := GetJobPermissions(c, ctx, userId)
	if err != nil {
		return nil, err
	}
	for _, item := range *list {
		if *item.Id == permId {
			return &item, nil
		}
	}
	return nil, nil
}

// DeleteJobPermission deletes a user's job permission.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#delete-remove-a-job-permission
func DeleteJobPermission(c *Client, ctx context.Context, jobId int, permId int) error {
	endpoint := fmt.Sprintf("v1/users/%d/permissions/jobs", jobId)
	jsonBody, err := json.Marshal(map[string]int{"job_permission_id": permId})
	if err != nil {
		return err
	}
	return Delete(c, ctx, endpoint, jsonBody)
}

// CreateJobPermission adds a job permission to a user.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#put-add-a-job-permission
func CreateJobPermission(c *Client, ctx context.Context, id int, obj *UserPermission) (int, error) {
	endpoint := fmt.Sprintf("v1/users/%d/permissions/jobs", id)
	return Create(c, ctx, endpoint, obj)
}

// GetFutureJobPermissions retrieves a list of a user's future job permissions.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-future-job-permissions
func GetFutureJobPermissions(c *Client, ctx context.Context, id int) (*[]FutureJobPermission, error) {
	var obj []FutureJobPermission
	endpoint := fmt.Sprintf("v1/users/%d/permissions/future_jobs", id)
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetFutureJobPermission retrieves a single future job permission by ID.
func GetFutureJobPermission(c *Client, ctx context.Context, userId int, permId int) (*FutureJobPermission, error) {
	list, err := GetFutureJobPermissions(c, ctx, userId)
	if err != nil {
		return nil, err
	}
	for _, item := range *list {
		if *item.Id == permId {
			return &item, nil
		}
	}
	return nil, nil
}

// DeleteFutureJobPermission deletes a user's future job permission.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#delete-remove-a-future-job-permission
func DeleteFutureJobPermission(c *Client, ctx context.Context, id int, permId int) error {
	endpoint := fmt.Sprintf("v1/users/%d/permissions/future_jobs", id)
	jsonBody, err := json.Marshal(map[string]int{"future_job_permission_id": permId})
	if err != nil {
		return err
	}
	return Delete(c, ctx, endpoint, jsonBody)
}

// CreateFutureJobPermission adds a future job permission to a user.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#put-add-a-future-job-permission
func CreateFutureJobPermission(c *Client, ctx context.Context, id int, obj *FutureJobPermission) (int, error) {
	endpoint := fmt.Sprintf("v1/users/%d/permissions/future_jobs", id)
	return Create(c, ctx, endpoint, obj)
}
