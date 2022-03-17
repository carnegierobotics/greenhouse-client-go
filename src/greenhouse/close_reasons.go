package greenhouse

import (
	"context"
)

func GetAllCloseReasons(c *Client) (*[]CloseReason, error) {
	var obj []CloseReason
	err := MultiGet(c, context.TODO(), "v1/close_reasons", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetCloseReason(c *Client, id int) (*CloseReason, error) {
	var obj CloseReason
	err := SingleGet(c, "close_reasons", id, &obj, context.TODO())
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
