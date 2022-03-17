package greenhouse

import (
  "context"
  "encoding/json"
  "fmt"
)

func ListApprovalsForJob(c *Client, id int) (*[]Approval, error) {
  var obj []Approval
  err := MultiGet(c, context.TODO(), fmt.Sprintf("v1/jobs/%d/approval_flows", id), "", &obj)
  if err != nil {
    return nil, err
  }
  return &obj, nil
}

func RetrieveApprovalFlow(c *Client, id int) (*Approval, error) {
  var obj Approval
  err := SingleGet(c, context.TODO(), fmt.Sprintf("v1/approval_flows/%d", id), &obj)
  if err != nil {
    return nil, err
  }
  return &obj, nil
}

func RequestApprovals(c *Client, id int) error {
  var obj []byte
  endpoint := fmt.Sprintf("v1/approval_flows/%d/request_approvals", id)
  _, err := Post(c, context.TODO(), endpoint, obj)
  if err != nil {
    return err
  }
  return nil
}

func PendingApprovalsForUser(c *Client, id int) (*[]Approval, error) {
  var obj []Approval
  endpoint := fmt.Sprintf("v1/users/%d/pending_approvals", id)
  err := MultiGet(c, context.TODO(), endpoint, "", &obj)
  if err != nil {
    return nil, err
  }
  return &obj, nil
}

func ReplaceApproverInApproverGroup(c *Client, group int, oldUser int, newUser int) error {
  endpoint := fmt.Sprintf("v1/approver_groups/%d/replace_approvers", group)
  replace := map[string]int{"remove_user_id": oldUser, "add_user_id": newUser}
  jsonBody, err := json.Marshal(replace)
  if err != nil {
    return err
  }
  _, err = Put(c, context.TODO(), endpoint, jsonBody)
  if err != nil {
    return err
  }
  return nil
}

func CreateReplaceApprovalFlow(c *Client, job int, obj *Approval) (int, error) {
  endpoint := fmt.Sprintf("v1/jobs/%d/approval_flows", job)
  return Create(c, context.TODO(), endpoint, obj)
}
