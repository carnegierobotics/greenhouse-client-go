package greenhouse

import (
	"context"
	"fmt"
)

func GetAllOffices(c *Client, ctx context.Context) (*[]Office, error) {
	var obj []Office
	err := MultiGet(c, ctx, "v1/offices", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetOffice(c *Client, ctx context.Context, id int) (*Office, error) {
	var obj Office
	endpoint := fmt.Sprintf("v1/offices/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateOffice(c *Client, ctx context.Context, obj *OfficeCreateInfo) (int, error) {
	return Create(c, ctx, "v1/offices", obj)
}

func UpdateOffice(c *Client, ctx context.Context, id int, obj *OfficeUpdateInfo) error {
	return Update(c, ctx, fmt.Sprintf("v1/offices/%d", id), obj)
}
