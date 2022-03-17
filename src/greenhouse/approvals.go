package greenhouse

import (
	"context"
	"encoding/json"
	"fmt"
)

func ListApprovalsForJob(c *Client, ctx context.Context, id int) (*[]Approval, error) {
	var obj []Approval
	err := MultiGet(c, ctx, fmt.Sprintf("v1/jobs/%d/approval_flows", id), "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func RetrieveApprovalFlow(c *Client, ctx context.Context, id int) (*Approval, error) {
	var obj Approval
	err := SingleGet(c, ctx, fmt.Sprintf("v1/approval_flows/%d", id), &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func RequestApprovals(c *Client, ctx context.Context, id int) error {
	var obj []byte
	endpoint := fmt.Sprintf("v1/approval_flows/%d/request_approvals", id)
	_, err := Post(c, ctx, endpoint, obj)
	return err
}

func PendingApprovalsForUser(c *Client, ctx context.Context, id int) (*[]Approval, error) {
	var obj []Approval
	endpoint := fmt.Sprintf("v1/users/%d/pending_approvals", id)
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

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

func CreateReplaceApprovalFlow(c *Client, ctx context.Context, job int, obj *Approval) (int, error) {
	endpoint := fmt.Sprintf("v1/jobs/%d/approval_flows", job)
	return Create(c, ctx, endpoint, obj)
}
