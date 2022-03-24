package greenhouse

import (
	"context"
	"fmt"
)

func GetAllDepartments(c *Client, ctx context.Context) (*[]Department, error) {
	var obj []Department
	err := MultiGet(c, ctx, "v1/departments", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetDepartment(c *Client, ctx context.Context, id int) (*Department, error) {
	var obj Department
	endpoint := fmt.Sprintf("v1/departments/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateDepartment(c *Client, ctx context.Context, obj *Department) (int, error) {
	return Create(c, ctx, "v1/departments", obj)
}

func UpdateDepartment(c *Client, ctx context.Context, id int, obj *Department) error {
	return Update(c, ctx, fmt.Sprintf("v1/departments/%d", id), obj)
}
