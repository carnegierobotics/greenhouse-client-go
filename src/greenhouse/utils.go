package greenhouse

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
  "net/url"
	"github.com/go-resty/resty/v2"
  "github.com/peterhellberg/link"
)

type RespObj struct {
	Id int `json:"id"`
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

func Create(c *Client, ctx context.Context, endpoint string, item interface{}) (int, error) {
	var respObj RespObj
	jsonBody, err := json.Marshal(item)
	if err != nil {
		return respObj.Id, err
	}
	resp, err := Post(c, ctx, endpoint, jsonBody)
	if err != nil {
		return respObj.Id, err
	}
	err = json.Unmarshal(resp.Body(), &respObj)
	if err != nil {
		return respObj.Id, err
	}
	return respObj.Id, nil
}

func Get(c *Client, ctx context.Context, endpoint string) (*resty.Response, error) {
	resp, err := c.Client.R().SetContext(ctx).Get(endpoint)
	if err != nil {
		return resp, err
	}
	if !resp.IsSuccess() {
		return resp, errors.New(resp.Status())
	}
	return resp, nil
}

func SingleGet(c *Client, ctx context.Context, endpoint string, item interface{}) error {
	resp, err := Get(c, ctx, endpoint)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp.Body(), &item)
	if err != nil {
		return err
	}
	return nil
}

func MultiGet(c *Client, ctx context.Context, endpoint string, querystring string, obj interface{}) error {
  allItems := make([]map[string]interface{}, 0)
  for {
    resp, err := Get(c, ctx, endpoint)
    if err != nil {
      return err
    }
    var itemsOnPage []map[string]interface{}
    err = json.Unmarshal(resp.Body(), &itemsOnPage)
    if err != nil {
      return err
    }
    for _, item := range itemsOnPage {
      allItems = append(allItems, item)
    }
    if resp.Header()["Link"] != nil {
      next := link.ParseResponse(resp.RawResponse)["next"]
      if next != nil {
        parsedUrl, err := url.ParseRequestURI(next.URI)
        if err != nil {
          return err
        }
        endpoint = fmt.Sprintf("%s?%s%s", parsedUrl.Path, parsedUrl.RawQuery, querystring)
      } else {
        break
      }
    } else {
      break
    }
  }
  wholeBody, err := json.Marshal(allItems)
  if err != nil {
    return err
  }
  err = json.Unmarshal(wholeBody, &obj)
  if err != nil {
    return err
  }
  return nil
}

func Exists(c *Client, ctx context.Context, endpoint string) (bool, error) {
	resp, err := Get(c, ctx, endpoint)
	if err != nil && resp.IsSuccess() {
		return false, nil
	}
	return err == nil, err
}

func Patch(c *Client, ctx context.Context, endpoint string, jsonBody []byte) (*resty.Response, error) {
	resp, err := c.Client.R().SetContext(ctx).SetBody(jsonBody).Patch(endpoint)
	if err != nil {
		return resp, err
	}
	if !resp.IsSuccess() {
		return resp, errors.New(resp.Status())
	}
	return resp, nil
}

func Put(c *Client, ctx context.Context, endpoint string, jsonBody []byte) (*resty.Response, error) {
  resp, err := c.Client.R().SetContext(ctx).SetBody(jsonBody).Put(endpoint)
  if err != nil {
    return resp, err
  }
  if !resp.IsSuccess() {
    return resp, errors.New(resp.Status())
  }
  return resp, nil
}

func Update(c *Client, ctx context.Context, endpoint string, item interface{}) error {
	jsonBody, err := json.Marshal(item)
	if err != nil {
		return err
	}
	_, err = Patch(c, ctx, endpoint, jsonBody)
	if err != nil {
		return err
	}
	return nil
}

func Delete(c *Client, ctx context.Context, endpoint string) error {
	resp, err := c.Client.R().SetContext(ctx).Delete(endpoint)
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return errors.New(resp.Status())
	}
	return nil
}
