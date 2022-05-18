/*
Copyright 2021-2022
Carnegie Robotics, LLC
4501 Hatfield Street, Pittsburgh, PA 15201
https://www.carnegierobotics.com
All rights reserved.

This file is part of greenhouse-client-go.

greenhouse-client-go is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

greenhouse-client-go is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with greenhouse-client-go. If not, see <https://www.gnu.org/licenses/>.
*/

package greenhouse

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/peterhellberg/link"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// RespObj represents a response object.
type RespObj struct {
	Id int `json:"id"`
}

// Post performs a POST against the specified endpoint.
func Post(c *Client, ctx context.Context, endpoint string, jsonBody []byte) (*resty.Response, error) {
	req := c.Client.R().SetContext(ctx)
	if jsonBody != nil {
		req = req.SetBody(jsonBody)
	}
	resp, err := req.Post(endpoint)
	if err != nil {
		return resp, err
	}
	if !resp.IsSuccess() {
		return resp, errors.New(resp.Status())
	}
	return resp, nil
}

// Create performs a POST and returns the ID of the created object.
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

// Get performs a GET against the specified endpoint.
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

// SingleGet performs a GET to retrieve a single object.
func SingleGet(c *Client, ctx context.Context, endpoint string, item interface{}) error {
	resp, err := Get(c, ctx, endpoint)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp.Body(), &item)
	if err != nil {
		return fmt.Errorf("%s. Response was: %+v", err.Error(), string(resp.Body()))
	}
	return nil
}

// MultiGet performs a GET to retrieve a list of objects, handling pagination in the process.
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

// Exists checks whether the specified object exists at the specified endpoint.
func Exists(c *Client, ctx context.Context, endpoint string) (bool, error) {
	resp, err := Get(c, ctx, endpoint)
	if err != nil && resp.IsSuccess() {
		return false, nil
	}
	return err == nil, err
}

// Patch performs a PATCH against the specified endpoint.
func Patch(c *Client, ctx context.Context, endpoint string, jsonBody []byte) (*resty.Response, error) {
	req := c.Client.R().SetContext(ctx)
	if jsonBody != nil {
		req = req.SetBody(jsonBody)
	}
	resp, err := req.Patch(endpoint)
	if err != nil {
		return resp, err
	}
	if !resp.IsSuccess() {
		return resp, errors.New(resp.Status())
	}
	return resp, nil
}

// Put performs a PUT against the specified endpoint.
func Put(c *Client, ctx context.Context, endpoint string, jsonBody []byte) (*resty.Response, error) {
	req := c.Client.R().SetContext(ctx)
	if jsonBody != nil {
		req = req.SetBody(jsonBody)
	}
	resp, err := req.Put(endpoint)
	if err != nil {
		return resp, err
	}
	if !resp.IsSuccess() {
		return resp, errors.New(resp.Status())
	}
	return resp, nil
}

// Update performs a Patch against the specified endpoint, handling object marshalling in the process.
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

// Delete performs a DELETE against the specified endpoint.
func Delete(c *Client, ctx context.Context, endpoint string, jsonBody []byte) error {
	req := c.Client.R().SetContext(ctx)
	if jsonBody != nil {
		req = req.SetBody(jsonBody)
	}
	resp, err := req.Delete(endpoint)
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return errors.New(resp.Status())
	}
	return nil
}

// GenerateQuerystring generates a querystring from a list of k-v pairs.
func GenerateQuerystring(params map[string]interface{}) (*string, error) {
	var querystring string
	for key, value := range params {
		val := Convert(value)
		if val != "" {
			querystring += fmt.Sprintf("%s=%s&", key, val)
		}
	}
	querystring = strings.TrimRight(querystring, "&")
	return &querystring, nil
}

// Convert (in theory) should be able to handle object type conversions. It was
// never really implemented and should be removed in future releases.
func Convert(item interface{}) string {
	//Needs to be expanded to support more types.
	switch valType := reflect.TypeOf(item); valType.Kind() {
	case reflect.Int:
		return strconv.Itoa(item.(int))
	case reflect.String:
		return item.(string)
	default:
		return ""
	}
}
