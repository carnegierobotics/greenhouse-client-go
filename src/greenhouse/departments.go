package greenhouse

import (
	"context"
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
	err := SingleGet(c, "departments", id, &obj, context.TODO())
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateDepartment(c *Client, obj *DepartmentCreateInfo) (int, error) {
	id, err := Create(c, "departments", obj, context.TODO())
	if err != nil {
		return id, err
	}
	return id, nil
}

func UpdateDepartment(c *Client, id int, obj *DepartmentUpdateInfo) error {
	err := Update(c, "departments", id, obj, context.TODO())
	if err != nil {
		return err
	}
	return nil
}
