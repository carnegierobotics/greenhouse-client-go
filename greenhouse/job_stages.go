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

// GetAllJobStages retrieves a list of all job stages for an organization.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-job-stages
func GetAllJobStages(c *Client, ctx context.Context) (*[]JobStage, error) {
	var obj []JobStage
	err := MultiGet(c, ctx, "v1/job_stages", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetJobStagesForJob retrieves a list of stages for a job.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-job-stages-for-job
func GetJobStagesForJob(c *Client, ctx context.Context, id int) (*[]JobStage, error) {
	var obj []JobStage
	endpoint := fmt.Sprintf("v1/jobs/%d/stages", id)
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetJobStage retrieves a job stage by ID.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-job-stage
func GetJobStage(c *Client, ctx context.Context, id int) (*JobStage, error) {
	var obj JobStage
	endpoint := fmt.Sprintf("v1/job_stages/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
