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

// GetAllEEOC retrieves a list of all EEOC data.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-eeoc
func GetAllEEOC(c *Client, ctx context.Context) (*[]EEOC, error) {
	var obj []EEOC
	endpoint := "v1/eeoc"
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetEEOCForApplication retrieves EEOC data for an application.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-eeoc-data-for-application
func GetEEOCForApplication(c *Client, ctx context.Context, applicationId int) (*EEOC, error) {
	var obj EEOC
	endpoint := fmt.Sprintf("v1/applications/%d", applicationId)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
