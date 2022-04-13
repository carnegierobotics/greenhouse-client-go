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
	"encoding/base64"
	//"errors"
	"fmt"
	// "encoding/json"
	"github.com/go-resty/resty/v2"
  "math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	BaseUrl      string
	Token        string
	OnBehalfOf   int
	RetryCount   int
	RetryWait    int64
	RetryMaxWait int64
	Client       *resty.Client
	// At some point, also need to implement the job board API stuff.
}

func (c *Client) BuildResty() error {
	u, err := url.ParseRequestURI(c.BaseUrl)
	if err != nil {
		return err
	}
	baseUrl := fmt.Sprintf("%s://%s", u.Scheme, u.Host)
	authString := base64.StdEncoding.EncodeToString([]byte(c.Token + ":"))
	c.Client = resty.New().
		SetHostURL(baseUrl).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Basic %s", authString)).
		SetHeader("On-Behalf-Of", strconv.Itoa(c.OnBehalfOf)).
		AddRetryCondition(
			func(r *resty.Response, err error) bool { return r.StatusCode() == http.StatusTooManyRequests },
		).
		SetRetryCount(c.RetryCount).
		SetRetryWaitTime(time.Duration(rand.Int63n(c.RetryWait)) * time.Second).
		SetRetryMaxWaitTime(time.Duration(rand.Int63n(c.RetryMaxWait)) * time.Second)
	return nil
}
