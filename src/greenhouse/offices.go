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

// GetAllOffices retrieves a list of all offices for an organization.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-offices
func GetAllOffices(c *Client, ctx context.Context) (*[]Office, error) {
	var obj []Office
	err := MultiGet(c, ctx, "v1/offices", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetOffice retrieves an office by ID.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-office
func GetOffice(c *Client, ctx context.Context, id int) (*Office, error) {
	var obj Office
	endpoint := fmt.Sprintf("v1/offices/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// CreateOffice creates a new office.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-office
func CreateOffice(c *Client, ctx context.Context, obj *OfficeCreateInfo) (int, error) {
	return Create(c, ctx, "v1/offices", obj)
}

// UpdateOffice updates an office's properties.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#patch-edit-office
func UpdateOffice(c *Client, ctx context.Context, id int, obj *OfficeUpdateInfo) error {
	return Update(c, ctx, fmt.Sprintf("v1/offices/%d", id), obj)
}
