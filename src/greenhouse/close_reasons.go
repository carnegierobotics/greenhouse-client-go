package greenhouse

import (
	"context"
	"fmt"
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
	endpoint := fmt.Sprintf("v1/close_reasons/%d", id)
	err := SingleGet(c, context.TODO(), endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
