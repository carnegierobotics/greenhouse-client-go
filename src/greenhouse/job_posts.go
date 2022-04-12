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
	"errors"
	"fmt"
	"strconv"
)

func GetAllJobPosts(c *Client, ctx context.Context, live bool, active bool) (*[]JobPost, error) {
	var obj []JobPost
	querystring := fmt.Sprintf("live=%s&active=%s", strconv.FormatBool(live), strconv.FormatBool(active))
	err := MultiGet(c, ctx, "v1/job_posts", querystring, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetJobPost(c *Client, ctx context.Context, id int) (*JobPost, error) {
	var obj JobPost
	err := SingleGet(c, ctx, fmt.Sprintf("v1/job_posts/%d", id), &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetAllJobPostsForJob(c *Client, ctx context.Context, id int) (*[]JobPost, error) {
	var obj []JobPost
	endpoint := fmt.Sprintf("v1/jobs/%d/job_posts", id)
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetJobPostForJob(c *Client, ctx context.Context, id int) (*JobPost, error) {
	//This should probably never be implemented since Greenhouse supports multiple posts per job as
	//of 2016.
	return nil, errors.New("GetJobPostForJob not implemented.")
}

func GetCustomLocationsForJobPost(c *Client, ctx context.Context, id int) (*[]Location, error) {
	var obj []Location
	endpoint := fmt.Sprintf("v1/job_posts/%d/custom_locations", id)
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func UpdateJobPost(c *Client, ctx context.Context, id int, obj *JobPost) error {
	return Update(c, ctx, fmt.Sprintf("v2/job_posts/%d", id), obj)
}

func UpdateJobPostStatus(c *Client, ctx context.Context, id int, status string) error {
	obj := map[string]string{"status": status}
	return Update(c, ctx, fmt.Sprintf("v2/job_posts/%d/status", id), &obj)
}
