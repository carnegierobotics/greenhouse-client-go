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

// ListApprovalsForJob lists all approvals for a job.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-approvals-for-job
func ListApprovalsForJob(c *Client, ctx context.Context, id int) (*[]Approval, error) {
	var obj []Approval
	err := MultiGet(c, ctx, fmt.Sprintf("v1/jobs/%d/approval_flows", id), "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// RetrieveApprovalFlow retrieves an approval flow by ID.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-approval-flow
func RetrieveApprovalFlow(c *Client, ctx context.Context, id int) (*Approval, error) {
	var obj Approval
	err := SingleGet(c, ctx, fmt.Sprintf("v1/approval_flows/%d", id), &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// RequestApprovals requests approvals as part of an approval flow.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-request-approvals
func RequestApprovals(c *Client, ctx context.Context, id int) error {
	var obj []byte
	endpoint := fmt.Sprintf("v1/approval_flows/%d/request_approvals", id)
	_, err := Post(c, ctx, endpoint, obj)
	return err
}

// PendingApprovalsForUser retrieves all pending approvals for a user.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-pending-approvals-for-user
func PendingApprovalsForUser(c *Client, ctx context.Context, id int) (*[]Approval, error) {
	var obj []Approval
	endpoint := fmt.Sprintf("v1/users/%d/pending_approvals", id)
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// ReplaceApproverInApproverGroup replaces a single approver in an approver group.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#put-replace-an-approver-in-an-approver-group
func ReplaceApproverInApproverGroup(c *Client, ctx context.Context, group int, oldUser int, newUser int) error {
	endpoint := fmt.Sprintf("v1/approver_groups/%d/replace_approvers", group)
	replace := map[string]int{"remove_user_id": oldUser, "add_user_id": newUser}
	jsonBody, err := json.Marshal(replace)
	if err != nil {
		return err
	}
	_, err = Put(c, ctx, endpoint, jsonBody)
	return err
}

// CreateReplaceApprovalFlow creates or replaces an approval workflow.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#put-create-or-replace-an-approval-flow
func CreateReplaceApprovalFlow(c *Client, ctx context.Context, job int, obj *Approval) (int, error) {
	endpoint := fmt.Sprintf("v1/jobs/%d/approval_flows", job)
	return Create(c, ctx, endpoint, obj)
}
