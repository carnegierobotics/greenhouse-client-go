package greenhouse

import (
	"context"
  "fmt"
)

func GetAllOffices(c *Client) (*[]Office, error) {
	var obj []Office
	err := MultiGet(c, context.TODO(), "v1/offices", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetOffice(c *Client, id int) (*Office, error) {
	var obj Office
  endpoint := fmt.Sprintf("v1/offices/%d", id)
	err := SingleGet(c, context.TODO(), endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateOffice(c *Client, obj *OfficeCreateInfo) (int, error) {
  return Create(c, context.TODO(), "v1/offices", obj)
}

func UpdateOffice(c *Client, id int, obj *OfficeUpdateInfo) error {
	return Update(c, context.TODO(), fmt.Sprintf("v1/offices/%d", id), obj)
}
