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

func GetAllApplications(c *Client, ctx context.Context) (*[]Application, error) {
	var objList []Application
	err := MultiGet(c, ctx, "v1/applications", "", &objList)
	if err != nil {
		return nil, err
	}
	return &objList, nil
}

func GetApplication(c *Client, ctx context.Context, id int) (*Application, error) {
	var obj Application
	err := SingleGet(c, ctx, fmt.Sprintf("v1/applications/%d", id), &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func DeleteApplication(c *Client, ctx context.Context, id int) error {
	return Delete(c, ctx, fmt.Sprintf("v1/applications/%d", id), nil)
}

func AddApplicationToCandidate(c *Client, ctx context.Context, id int, obj *Application) (int, error) {
	endpoint := fmt.Sprintf("v1/candidates/%d/applications", id)
	return Create(c, ctx, endpoint, obj)
}

func UpdateApplication(c *Client, ctx context.Context, id int, obj *Application) error {
	endpoint := fmt.Sprintf("v1/applications/%d", id)
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = Patch(c, ctx, endpoint, jsonBody)
	return err
}

func AdvanceApplication(c *Client, ctx context.Context, id int, from int) error {
	endpoint := fmt.Sprintf("v1/applications/%d/advance", id)
	jsonBody, err := json.Marshal(map[string]int{"from_stage_id": from})
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, endpoint, jsonBody)
	return err
}

func MoveApplicationDifferentJob(c *Client, ctx context.Context, id int, newJob int, newStage int) error {
	endpoint := fmt.Sprintf("v1/applications/%d/transfer_to_job", id)
	jsonBody, err := json.Marshal(map[string]int{"new_job_id": newJob, "new_stage_id": newStage})
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, endpoint, jsonBody)
	return err
}

func MoveApplicationSameJob(c *Client, ctx context.Context, id int, from int, to int) error {
	endpoint := fmt.Sprintf("v1/applications/%d/move", id)
	jsonBody, err := json.Marshal(map[string]int{"from_stage_id": from, "to_stage_id": to})
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, endpoint, jsonBody)
	return err
}

func AddAttachmentToApplication(c *Client, ctx context.Context, id int, item *Attachment) error {
	endpoint := fmt.Sprintf("v1/applications/%d/attachments", id)
	jsonBody, err := json.Marshal(item)
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, endpoint, jsonBody)
	return err
}

func HireApplication(c *Client, ctx context.Context, id int, obj *ApplicationHire) error {
	endpoint := fmt.Sprintf("v1/applications/%d/hire", id)
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, endpoint, jsonBody)
	return err
}

func RejectApplication(c *Client, ctx context.Context, id int, obj *ApplicationReject) error {
	endpoint := fmt.Sprintf("v1/applications/%d/reject", id)
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, endpoint, jsonBody)
	return err
}

func UpdateRejectionReason(c *Client, ctx context.Context, id int, rejectionReason int) error {
	endpoint := fmt.Sprintf("v1/applications/%d/reject", id)
	jsonBody, err := json.Marshal(map[string]int{"rejection_reason_id": rejectionReason})
	if err != nil {
		return err
	}
	_, err = Patch(c, ctx, endpoint, jsonBody)
	return err
}

func UnrejectApplication(c *Client, ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("v1/applications/%d/unreject", id)
	_, err := Post(c, ctx, endpoint, nil)
	return err
}
