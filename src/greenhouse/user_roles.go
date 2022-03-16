package greenhouse

import (
  "context"
)

func GetAllUserRoles(c *Client) (*[]UserRole, error) {
	var obj []UserRole
  err := GetAll(c, "user_roles", &obj, context.TODO())
  if err != nil {
    return nil, err
  }
  return &obj, nil
}
