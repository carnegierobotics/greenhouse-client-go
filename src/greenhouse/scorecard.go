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

// GetAllScorecards retrieves a list of all scorecards for an organization.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-scorecards
func GetAllScorecards(c *Client, ctx context.Context) (*[]Scorecard, error) {
	var obj []Scorecard
	err := MultiGet(c, ctx, "v1/scorecards", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetScorecardsForApplication retrieves a list of all scorecards for an application.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-scorecards-for-application
func GetScorecardsForApplication(c *Client, ctx context.Context, id int) (*[]Scorecard, error) {
	var obj []Scorecard
	endpoint := fmt.Sprintf("v1/applications/%d/scorecards", id)
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetScorecard retrieves a scorecard by ID.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-scorecard
func GetScorecard(c *Client, ctx context.Context, id int) (*Scorecard, error) {
	var obj Scorecard
	endpoint := fmt.Sprintf("v1/scorecards/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
