package greenhouse

import (
	"context"
)

func GetAllUserRoles(c *Client, ctx context.Context) (*[]UserRole, error) {
	var obj []UserRole
	err := MultiGet(c, ctx, "v1/user_roles", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
