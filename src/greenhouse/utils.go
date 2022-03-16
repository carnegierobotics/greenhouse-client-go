package greenhouse

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
  "github.com/go-resty/resty/v2"
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
	resp, err := Post(c, ctx, fmt.Sprintf("v1/%s", itemType), jsonBody)
	if err != nil {
		return respObj.Id, err
	}
  err = json.Unmarshal(resp.Body(), &respObj)
  if err != nil {
    return respObj.Id, err
  }
	return respObj.Id, nil
}

func Post(c *Client, ctx context.Context, endpoint string, jsonBody []byte) (*resty.Response, error) {
  resp, err := c.Client.R().SetContext(ctx).SetBody(jsonBody).Post(endpoint)
	if err != nil {
		return resp, err
	}
	if !resp.IsSuccess() {
		return resp, errors.New(resp.Status())
	}
  return resp, nil
}

func Exists(c *Client, itemType string, id int, ctx context.Context) (bool, error) {
	resp, err := Get(ctx, c, fmt.Sprintf("v1/%s/%d", itemType, id))
	if err != nil && resp.IsSuccess() {
		return false, nil
	}
	return err == nil, err
}

func GetById(c *Client, itemType string, id int, item interface{}, ctx context.Context) error {
	resp, err := Get(ctx, c, fmt.Sprintf("v1/%s/%d", itemType, id))
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
	resp, err := Get(ctx, c, fmt.Sprintf("v1/%s", itemType))
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp.Body(), &itemList)
	if err != nil {
		return err
	}
	return nil
}

func Get(ctx context.Context, c *Client, endpoint string) (*resty.Response, error) {
  resp, err := c.Client.R().SetContext(ctx).Get(endpoint)
  if err != nil {
    return resp, err
  }
  if !resp.IsSuccess() {
    return resp, errors.New(resp.Status())
  }
  return resp, nil
}

func Update(c *Client, itemType string, id int, item interface{}, ctx context.Context) error {
	jsonBody, err := json.Marshal(item)
	if err != nil {
		return err
	}
	_, err = Patch(ctx, c, fmt.Sprintf("v1/%s/%d", itemType, id), jsonBody)
	if err != nil {
		return err
	}
	return nil
}

func Patch(ctx context.Context, c *Client, endpoint string, jsonBody []byte) (*resty.Response, error) {
  resp, err := c.Client.R().SetContext(ctx).SetBody(jsonBody).Patch(endpoint)
	if err != nil {
		return resp, err
	}
	if !resp.IsSuccess() {
		return resp, errors.New(resp.Status())
	}
  return resp, nil
}

func Delete(ctx context.Context, c*Client, endpoint string) error {
  resp, err := c.Client.R().SetContext(ctx).Delete(endpoint)
  if err != nil {
    return err
  }
  if !resp.IsSuccess() {
    return errors.New(resp.Status())
  }
  return nil
}
