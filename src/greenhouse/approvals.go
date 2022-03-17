package greenhouse

import (
	"errors"
  "fmt"
)

func ListApprovalsForJob(c *Client, id int) error {
  var obj []Approval
  err := PaginatedGet(c, context.TODO(), fmt.Sprintf("v1/jobs/%d/approval_flows", id), "", &obj)
  if err != nil {
    return nil, err
  }
  return &obj, nil
}

func RetrieveApprovalFlow() error {
	return errors.New("RetrieveApprovalFlow not implemented.")
}

func RequestApprovals() error {
	return errors.New("RequestApprovals not implemented.")
}

func PendingApprovalsForUser() error {
	return errors.New("PendingApprovalsForUser not imeplemented.")
}

func ReplaceApproverInApproverGroup() error {
	return errors.New("ReplaceApproverInApproverGroup not implemented.")
}

func CreateReplaceApprovalFlow() error {
	return errors.New("CreateReplaceApprovalFlow not implemented.")
}
