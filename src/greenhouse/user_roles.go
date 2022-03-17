package greenhouse

import (
  "context"
)

func GetAllUserRoles(c *Client) (*[]UserRole, error) {
	var obj []UserRole
  err := MultiGet(c, context.TODO(), "v1/user_roles", "", &obj)
  if err != nil {
    return nil, err
  }
  return &obj, nil
}
