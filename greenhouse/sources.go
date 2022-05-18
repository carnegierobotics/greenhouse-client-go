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

// GetAllSources retrieves a list of all sources for an organization, grouped by strategy.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-sources
func GetAllSources(c *Client, ctx context.Context) (*[]Source, error) {
	var obj []Source
	err := MultiGet(c, ctx, "v1/sources", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
