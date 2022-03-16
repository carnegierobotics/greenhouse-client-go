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

func PaginatedGet(c *Client, ctx context.Context, endpoint string, custom_query string) ([]byte, error) {
  // fmt.Printf(endpoint)
  allItems := make([]map[string]interface{}, 0)
  for {
    resp, err := Get(ctx, c, endpoint)
    // fmt.Printf("Request: %s\n", resp.Request.URL)
    if err != nil {
      return nil, err
    }
    // fmt.Printf("Body: %+v\n", wholeBody)
    var itemsOnPage []map[string]interface{}
    err = json.Unmarshal(resp.Body(), &itemsOnPage)
    if err != nil {
      return nil, err
    }
    for _, item := range itemsOnPage {
      allItems = append(allItems, item)
    }
    if resp.Header()["Link"] != nil {
      // fmt.Printf("%+v\n", resp.RawResponse)
      next := link.ParseResponse(resp.RawResponse)["next"]
      // fmt.Printf("Next: %+v\n", next)
      if next != nil {
        parsedUrl, err := url.ParseRequestURI(next.URI)
        if err != nil {
          return nil, err
        }
        endpoint = fmt.Sprintf("%s?%s", parsedUrl.Path, parsedUrl.RawQuery)
        if custom_query != "" {
          endpoint = fmt.Sprintf("%s%s", endpoint, custom_query)
        }
        // fmt.Printf("New endpoint: %+v\n", endpoint)
      } else {
        break
      }
    } else {
      break
    }
  }
  wholeBody, err := json.Marshal(allItems)
  if err != nil {
    return nil, err
  }
  return wholeBody, nil
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

func Delete(ctx context.Context, c *Client, endpoint string) error {
	resp, err := c.Client.R().SetContext(ctx).Delete(endpoint)
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return errors.New(resp.Status())
	}
	return nil
}
