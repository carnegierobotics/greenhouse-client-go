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

// GetAllUsers retrieves a list of all users for an organization.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-users
func GetAllUsers(c *Client, ctx context.Context) (*[]User, error) {
	var obj []User
	err := MultiGet(c, ctx, "v1/users", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetUser retrieves a single user by ID.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-user
func GetUser(c *Client, ctx context.Context, id int) (*User, error) {
	var obj User
	endpoint := fmt.Sprintf("v1/users/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// CreateUser creates a new user.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-user
func CreateUser(c *Client, ctx context.Context, obj *UserCreateInfo) (int, error) {
	return Create(c, ctx, "v1/users", obj)
}

// EnableUser enables a user.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#patch-enable-user
func EnableUser(c *Client, ctx context.Context, id int) error {
	lookupInfo := GetLookupInfo(id)
	_, err := Patch(c, ctx, "v2/users/enable", lookupInfo)
	return err
}

// DisableUser disables a user.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#patch-disable-user
func DisableUser(c *Client, ctx context.Context, id int) error {
	lookupInfo := GetLookupInfo(id)
	_, err := Patch(c, ctx, "v2/users/disable", lookupInfo)
	return err
}

// UpdateUser updates a user's properties.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#patch-edit-user
func UpdateUser(c *Client, ctx context.Context, id int, obj *UserUpdateInfo) error {
	return Update(c, ctx, fmt.Sprintf("v1/users/%d", id), obj)
}

// GetLookupInfo formats a user lookup. Currently this can only be done by ID, although
// Greenhouse also supports other attributes, such as name.
func GetLookupInfo(id int) []byte {
	return []byte(fmt.Sprintf("{\"user\":{\"user_id\":%d}}", id))
}

// ChangeUserPermissionLevel changes a user's permission level. It currently only supports changing
// to Basic.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#patch-change-user-permission-level
func ChangeUserPermissionLevel(c *Client, ctx context.Context, user *User, level string) error {
	endpoint := fmt.Sprintf("v1/users/permission_level")
	jsonBody, err := json.Marshal(map[string]interface{}{"user": user, "level": level})
	if err != nil {
		return err
	}
	_, err = Patch(c, ctx, endpoint, jsonBody)
	return err
}

// AddEmailAddressToUser adds an email address to a user.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-e-mail-address-to-user
func AddEmailAddressToUser(c *Client, ctx context.Context, userId int, obj *UserEmailUpdateInfo) error {
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, fmt.Sprintf("v1/users/%d/email_addresses", userId), jsonBody)
	return err
}
