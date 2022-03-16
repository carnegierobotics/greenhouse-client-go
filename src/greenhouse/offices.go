package greenhouse

import (
	"context"
)

func GetAllOffices(c *Client) (*[]Office, error) {
	var obj []Office
	err := PaginatedGet(c, context.TODO(), "v1/offices", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetOffice(c *Client, id int) (*Office, error) {
	var obj Office
	err := GetById(c, "offices", id, &obj, context.TODO())
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateOffice(c *Client, obj *OfficeCreateInfo) (int, error) {
	id, err := Create(c, "offices", obj, context.TODO())
	if err != nil {
		return id, err
	}
	return id, nil
}

func UpdateOffice(c *Client, id int, obj *OfficeUpdateInfo) error {
	err := Update(c, "offices", id, obj, context.TODO())
	if err != nil {
		return err
	}
	return nil
}
