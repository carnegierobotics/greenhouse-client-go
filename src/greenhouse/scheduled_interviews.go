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
	"strconv"
)

// GetAllScheduledInterviews retrieves a list of all scheduled interviews for an organization.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-scheduled-interviews
func GetAllScheduledInterviews(c *Client, ctx context.Context, actionable bool) (*[]ScheduledInterview, error) {
	var obj []ScheduledInterview
	querystring := fmt.Sprintf("&actionable=%s", strconv.FormatBool(actionable))
	endpoint := fmt.Sprintf("v1/scheduled_interviews?%s", querystring)
	err := MultiGet(c, ctx, endpoint, querystring, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetScheduledInterviewsForApplication retrieves a list of all scheduled interviews for an application.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-scheduled-interviews-for-application
func GetScheduledInterviewsForApplication(c *Client, ctx context.Context, id int, actionable bool) (*[]ScheduledInterview, error) {
	var obj []ScheduledInterview
	querystring := fmt.Sprintf("&actionable=%s", strconv.FormatBool(actionable))
	endpoint := fmt.Sprintf("v1/applications/%d/scheduled_interviews%s", id, querystring)
	err := MultiGet(c, ctx, endpoint, querystring, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetScheduledInterview retrieves a scheduled interview by ID.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-scheduled-interview
func GetScheduledInterview(c *Client, ctx context.Context, id int) (*ScheduledInterview, error) {
	var obj ScheduledInterview
	endpoint := fmt.Sprintf("v1/scheduled_interviews/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// CreateScheduledInterview creates a new scheduled interview.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-create-scheduled-interview
func CreateScheduledInterview(c *Client, ctx context.Context, obj *ScheduledInterviewCreateInfo) (int, error) {
	return Create(c, ctx, "v2/scheduled_interviews", obj)
}

// UpdateScheduledInterview updates a scheduled interview's properties.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#patch-update-scheduled-interview
func UpdateScheduledInterview(c *Client, ctx context.Context, id int, obj *ScheduledInterviewUpdateInfo) error {
	return Update(c, ctx, fmt.Sprintf("v2/scheduled_interviews/%d", id), obj)
}

// DeleteScheduledInterview deletes a scheduled interview.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#delete-remove-scheduled-interview
func DeleteScheduledInterview(c *Client, ctx context.Context, id int) error {
	return Delete(c, ctx, fmt.Sprintf("v1/scheduled_interviews/%d", id), nil)
}
