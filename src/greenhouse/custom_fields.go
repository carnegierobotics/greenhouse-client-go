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
	"strconv"
)

// GetAllCustomFields retrieves a list of all custom fields.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-custom-fields
func GetAllCustomFields(c *Client, ctx context.Context, fieldType string, includeInactive bool) (*[]CustomField, error) {
	var obj []CustomField
	endpoint := "v1/custom_fields"
	if fieldType != "" {
		endpoint = fmt.Sprintf("%s/%s", endpoint, fieldType)
	}
	querystring := fmt.Sprintf("include_inactive=%s", strconv.FormatBool(includeInactive))
	err := MultiGet(c, ctx, endpoint, querystring, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetCustomField retrieves a custom field by ID.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-custom-field
func GetCustomField(c *Client, ctx context.Context, id int) (*CustomField, error) {
	var obj CustomField
	err := SingleGet(c, ctx, fmt.Sprintf("v1/custom_field/%d", id), &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// CreateCustomField creates a custom field.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-create-custom-field
func CreateCustomField(c *Client, ctx context.Context, obj *CustomField) (int, error) {
	return Create(c, ctx, "v1/custom_fields", obj)
}

// UpdateCustomField updates a custom field.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#patch-update-custom-field
func UpdateCustomField(c *Client, ctx context.Context, id int, obj *CustomField) error {
	return Update(c, ctx, fmt.Sprintf("v1/custom_fields/%d", id), obj)
}

// DeleteCustomField deletes a custom field.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#delete-delete-custom-field
func DeleteCustomField(c *Client, ctx context.Context, id int) error {
	return Delete(c, ctx, fmt.Sprintf("v1/custom_fields/%d", id), nil)
}

// GetCustomFieldOptions retrieves a list of field options for a custom field.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-custom-field-options
func GetCustomFieldOptions(c *Client, ctx context.Context, id int, typeStr string) (*[]CustomFieldOption, error) {
	var obj []CustomFieldOption
	endpoint := fmt.Sprintf("v1/custom_field/%d/custom_field_options", id)
	querystring := fmt.Sprintf("type=%s", typeStr)
	err := MultiGet(c, ctx, endpoint, querystring, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// CreateCustomFieldOptions creates an option for a custom field.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-create-custom-field-options
func CreateCustomFieldOptions(c *Client, ctx context.Context, id int, obj *[]CustomFieldOption) (int, error) {
	endpoint := fmt.Sprintf("v1/custom_field/%d/custom_field_options", id)
	return Create(c, ctx, endpoint, obj)
}

// UpdateCustomFieldOptions updates the options for a custom field.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#patch-update-custom-field-options
func UpdateCustomFieldOptions(c *Client, ctx context.Context, id int, obj *[]CustomFieldOption) error {
	endpoint := fmt.Sprintf("v1/custom_field/%d/custom_field_options", id)
	return Update(c, ctx, endpoint, obj)
}

// DeleteCustomFieldOptions deletes options for a custom field.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#delete-remove-custom-field-options
func DeleteCustomFieldOptions(c *Client, ctx context.Context, id int, ids []int) error {
	jsonBody, err := json.Marshal(map[string][]int{"option_ids": ids})
	if err != nil {
		return err
	}
	return Delete(c, ctx, fmt.Sprintf("v1/custom_field/%d/custom_field_options", id), jsonBody)
}
