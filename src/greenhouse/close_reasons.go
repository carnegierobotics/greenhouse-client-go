package greenhouse

import (
	"context"
)

type CloseReason struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetAllCloseReasons(c *Client) (*[]CloseReason, error) {
	var obj []CloseReason
	err := GetAll(c, "close_reason", &obj, context.TODO())
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetCloseReason(c *Client, id int) (*CloseReason, error) {
	var obj CloseReason
	err := GetById(c, "close_reasons", id, &obj, context.TODO())
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
