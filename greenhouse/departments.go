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

// GetAllDepartments retrieves a list of all departments.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-departments
func GetAllDepartments(c *Client, ctx context.Context) (*[]Department, error) {
	var obj []Department
	err := MultiGet(c, ctx, "v1/departments", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetDepartment retrieves a single department by ID.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-department
func GetDepartment(c *Client, ctx context.Context, id int) (*Department, error) {
	var obj Department
	endpoint := fmt.Sprintf("v1/departments/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// CreateDepartment creates a department.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-department
func CreateDepartment(c *Client, ctx context.Context, obj *Department) (int, error) {
	return Create(c, ctx, "v1/departments", obj)
}

// UpdateDepartment updates a department's properties.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#patch-edit-department
func UpdateDepartment(c *Client, ctx context.Context, id int, obj *Department) error {
	return Update(c, ctx, fmt.Sprintf("v1/departments/%d", id), obj)
}
