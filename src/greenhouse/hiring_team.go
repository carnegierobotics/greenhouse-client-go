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

// GetJobHiringTeam retrieves the hiring team for a job.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-hiring-team
func GetJobHiringTeam(c *Client, ctx context.Context, id int) (*map[string][]HiringMember, error) {
	resp, err := c.Client.R().SetContext(ctx).Get(fmt.Sprintf("v1/jobs/%d/hiring_team", id))
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("Error getting hiring team: %d\n", id)
	}
	var obj map[string][]HiringMember
	err = json.Unmarshal(resp.Body(), &obj)
	return &obj, nil
}

// UpdateJobHiringTeam replaces a job's hiring team.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#put-replace-hiring-team
func UpdateJobHiringTeam(c *Client, ctx context.Context, id int, obj *map[string][]HiringMember) error {
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = Put(c, ctx, fmt.Sprintf("v1/jobs/%d/hiring_team", id), jsonBody)
	return err
}

// UpdateHiringTeamMembers adds members to a hiring team. This is a more atomic operation; instead of
// replacing an entire team, you can work on one member.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-hiring-team-members
func UpdateHiringTeamMembers(c *Client, ctx context.Context, id int, obj *map[string][]HiringMember) error {
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, fmt.Sprintf("v1/jobs/%d/hiring_team", id), jsonBody)
	return err
}

// DeleteHiringTeamMembers removes members from a hiring team. This is a more atomic operation; instead of
// replacing an entire team, you can work on one member.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#delete-remove-hiring-team-member
func DeleteHiringTeamMembers(c *Client, ctx context.Context, id int, obj *map[string][]int) error {
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	err = Delete(c, ctx, fmt.Sprintf("v1/jobs/%d/hiring_team", id), jsonBody)
	return err
}
