package greenhouse

import (
	"context"
	"fmt"
)

func GetAllCloseReasons(c *Client, ctx context.Context) (*[]CloseReason, error) {
	var obj []CloseReason
	err := MultiGet(c, ctx, "v1/close_reasons", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetCloseReason(c *Client, ctx context.Context, id int) (*CloseReason, error) {
	var obj CloseReason
	endpoint := fmt.Sprintf("v1/close_reasons/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
