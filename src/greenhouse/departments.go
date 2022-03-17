package greenhouse

import (
	"context"
  "fmt"
)

func GetAllDepartments(c *Client) (*[]Department, error) {
	var obj []Department
	err := MultiGet(c, context.TODO(), "v1/departments", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetDepartment(c *Client, id int) (*Department, error) {
	var obj Department
  endpoint := fmt.Sprintf("v1/departments/%d", id)
	err := SingleGet(c, context.TODO(), endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateDepartment(c *Client, obj *DepartmentCreateInfo) (int, error) {
	return Create(c, context.TODO(), "v1/departments", obj)
}

func UpdateDepartment(c *Client, id int, obj *DepartmentUpdateInfo) error {
	return Update(c, context.TODO(), fmt.Sprintf("v1/departments/%d", id), obj)
}
