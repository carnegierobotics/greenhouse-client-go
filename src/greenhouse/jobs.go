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
	"fmt"
)

// GetAllJobs gets a list of all jobs for an organization.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-jobs
func GetAllJobs(c *Client, ctx context.Context) (*[]Job, error) {
	var obj []Job
	err := MultiGet(c, ctx, "v1/jobs", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetJob retrieves a job by ID.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-job
func GetJob(c *Client, ctx context.Context, id int) (*Job, error) {
	var obj Job
	err := SingleGet(c, ctx, fmt.Sprintf("v1/jobs/%d", id), &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// CreateJob creates a new job.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-create-job
func CreateJob(c *Client, ctx context.Context, obj *JobCreateInfo) (int, error) {
	return Create(c, ctx, "v1/jobs", obj)
}

// UpdateJob updates a job's properties.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#patch-update-job
func UpdateJob(c *Client, ctx context.Context, id int, obj *JobUpdateInfo) error {
	return Update(c, ctx, fmt.Sprintf("v1/jobs/%d", id), obj)
}
