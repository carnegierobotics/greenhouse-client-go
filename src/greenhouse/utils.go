package greenhouse

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

type RespObj struct {
	Id int `json:"id"`
}

func Create(c *Client, itemType string, item interface{}, ctx context.Context) (int, error) {
	var respObj RespObj
	jsonBody, err := json.Marshal(item)
	if err != nil {
		return respObj.Id, err
	}
	resp, err := c.Client.R().SetContext(ctx).SetBody(jsonBody).Post(fmt.Sprintf("v1/%s", itemType))
	err = json.Unmarshal(resp.Body(), &respObj)
	if err != nil {
		return respObj.Id, err
	}
	if !resp.IsSuccess() {
		return respObj.Id, errors.New(resp.Status())
	}
	return respObj.Id, nil
}

func Exists(c *Client, itemType string, id int, ctx context.Context) (bool, error) {
	resp, err := c.Client.R().SetContext(ctx).Get(fmt.Sprintf("v1/%s/%d", itemType, id))
	if err != nil && resp.IsSuccess() {
		return false, nil
	}
	return err == nil, err
}

func GetById(c *Client, itemType string, id int, item interface{}, ctx context.Context) error {
	resp, err := c.Client.R().SetContext(ctx).Get(fmt.Sprintf("v1/%s/%d", itemType, id))
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp.Body(), &item)
	if err != nil {
		return err
	}
	return nil
}

func GetAll(c *Client, itemType string, itemList interface{}, ctx context.Context) error {
	resp, err := c.Client.R().SetContext(ctx).Get(fmt.Sprintf("v1/%s", itemType))
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp.Body(), &itemList)
	if err != nil {
		return err
	}
	return nil
}

func Update(c *Client, itemType string, id int, item interface{}, ctx context.Context) error {
	jsonBody, err := json.Marshal(item)
	if err != nil {
		return err
	}
	resp, err := c.Client.R().SetContext(ctx).SetBody(jsonBody).Patch(fmt.Sprintf("v1/%s/%d", itemType, id))
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return errors.New(resp.Status())
	}
	return nil
}
