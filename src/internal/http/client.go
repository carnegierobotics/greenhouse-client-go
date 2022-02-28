package http

import (
  "encoding/base64"
  "fmt"
  // "encoding/json"
  "net/url"
  "github.com/go-resty/resty/v2"
)

type Client struct {
  BaseUrl   string
  Token string
  OnBehalfOf string
  Client *resty.Client
}

func (c *Client) BuildResty() (error) {
  u, err := url.ParseRequestURI(c.BaseUrl)
  if err != nil {
    return err
  }
  baseUrl := fmt.Sprintf("%s://%s", u.Scheme, u.Host)
  authString := base64.StdEncoding.EncodeToString([]byte(c.Token + ":"))
  c.Client = resty.New().SetHostURL(baseUrl).SetHeader("Authorization", fmt.Sprintf("Basic %s", authString)).SetHeader("On-Behalf-Of", c.OnBehalfOf).SetRetryCount(5)
  return nil
}
