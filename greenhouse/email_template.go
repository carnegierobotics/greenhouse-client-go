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

// GetAllEmailTemplates retrieves a list of all email templates.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-email-templates
func GetAllEmailTemplates(c *Client, ctx context.Context) (*[]EmailTemplate, error) {
	var obj []EmailTemplate
	err := MultiGet(c, ctx, "v1/email_templates", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetEmailTemplate retrieves an email template by ID.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-email-template
func GetEmailTemplate(c *Client, ctx context.Context, id int) (*EmailTemplate, error) {
	var obj EmailTemplate
	err := SingleGet(c, ctx, fmt.Sprintf("v1/email_templates/%d", id), &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
