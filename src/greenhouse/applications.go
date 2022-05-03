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

// GetAllApplications retrieves all applications.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-applications
func GetAllApplications(c *Client, ctx context.Context) (*[]Application, error) {
	var objList []Application
	err := MultiGet(c, ctx, "v1/applications", "", &objList)
	if err != nil {
		return nil, err
	}
	return &objList, nil
}

// GetApplication gets a single application by ID.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-application
func GetApplication(c *Client, ctx context.Context, id int) (*Application, error) {
	var obj Application
	err := SingleGet(c, ctx, fmt.Sprintf("v1/applications/%d", id), &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// DeleteApplication deletes a single application by ID.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#delete-delete-application
func DeleteApplication(c *Client, ctx context.Context, id int) error {
	return Delete(c, ctx, fmt.Sprintf("v1/applications/%d", id), nil)
}

// AddApplicationToCandidate adds an application to a candidate.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-application-to-candidate-prospect
func AddApplicationToCandidate(c *Client, ctx context.Context, id int, obj *Application) (int, error) {
	endpoint := fmt.Sprintf("v1/candidates/%d/applications", id)
	return Create(c, ctx, endpoint, obj)
}

// UpdateApplication updates an application's properties.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#patch-update-application
func UpdateApplication(c *Client, ctx context.Context, id int, obj *Application) error {
	endpoint := fmt.Sprintf("v1/applications/%d", id)
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = Patch(c, ctx, endpoint, jsonBody)
	return err
}

// AdvanceApplication moves an application to the next stage in an approval flow.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-advance-application
func AdvanceApplication(c *Client, ctx context.Context, id int, from int) error {
	endpoint := fmt.Sprintf("v1/applications/%d/advance", id)
	jsonBody, err := json.Marshal(map[string]int{"from_stage_id": from})
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, endpoint, jsonBody)
	return err
}

// MoveApplicationDifferentJob moves an application to a different job.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-move-application-different-job
func MoveApplicationDifferentJob(c *Client, ctx context.Context, id int, newJob int, newStage int) error {
	endpoint := fmt.Sprintf("v1/applications/%d/transfer_to_job", id)
	jsonBody, err := json.Marshal(map[string]int{"new_job_id": newJob, "new_stage_id": newStage})
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, endpoint, jsonBody)
	return err
}

// MoveApplicationSameJob moves an application to a different stage in the same job.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-move-application-same-job
func MoveApplicationSameJob(c *Client, ctx context.Context, id int, from int, to int) error {
	endpoint := fmt.Sprintf("v1/applications/%d/move", id)
	jsonBody, err := json.Marshal(map[string]int{"from_stage_id": from, "to_stage_id": to})
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, endpoint, jsonBody)
	return err
}

// AddAttachmentToApplication adds an attachment to an application.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-attachment-to-application
func AddAttachmentToApplication(c *Client, ctx context.Context, id int, item *Attachment) error {
	endpoint := fmt.Sprintf("v1/applications/%d/attachments", id)
	jsonBody, err := json.Marshal(item)
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, endpoint, jsonBody)
	return err
}

// HireApplication hires a candidate based on an application.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-hire-application
func HireApplication(c *Client, ctx context.Context, id int, obj *ApplicationHire) error {
	endpoint := fmt.Sprintf("v1/applications/%d/hire", id)
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, endpoint, jsonBody)
	return err
}

// RejectApplication rejects a candidate based on an application.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-reject-application
func RejectApplication(c *Client, ctx context.Context, id int, obj *ApplicationReject) error {
	endpoint := fmt.Sprintf("v1/applications/%d/reject", id)
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, endpoint, jsonBody)
	return err
}

// UpdateRejectionReason updates an application's rejection reason.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#patch-update-rejection-reason
func UpdateRejectionReason(c *Client, ctx context.Context, id int, rejectionReason int) error {
	endpoint := fmt.Sprintf("v1/applications/%d/reject", id)
	jsonBody, err := json.Marshal(map[string]int{"rejection_reason_id": rejectionReason})
	if err != nil {
		return err
	}
	_, err = Patch(c, ctx, endpoint, jsonBody)
	return err
}

// UnrejectApplication unrejects an application.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-unreject-application
func UnrejectApplication(c *Client, ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("v1/applications/%d/unreject", id)
	_, err := Post(c, ctx, endpoint, nil)
	return err
}
