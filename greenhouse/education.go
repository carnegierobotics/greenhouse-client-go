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
)

// GetAllDegrees retrieves a list of all degrees.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-degrees
func GetAllDegrees(c *Client, ctx context.Context) (*[]Degree, error) {
	var obj []Degree
	err := MultiGet(c, ctx, "v1/degrees", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetAllDisciplines retrieves a list of all disciplines.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-disciplines
func GetAllDisciplines(c *Client, ctx context.Context) (*[]Discipline, error) {
	var obj []Discipline
	err := MultiGet(c, ctx, "v1/disciplines", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetAllSchoolds retrieves a list of all schools.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-schools
func GetAllSchools(c *Client, ctx context.Context) (*[]School, error) {
	var obj []School
	err := MultiGet(c, ctx, "v1/schools", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
