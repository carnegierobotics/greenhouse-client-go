package http

import (
  "encoding/base64"
  "fmt"
  // "encoding/json"
  "net/url"
  "strconv"
  "github.com/go-resty/resty/v2"
)

type Client struct {
  BaseUrl   string
  Token string
  OnBehalfOf int
  Client *resty.Client
  // At some point, also need to implement the job board API stuff.
}

func (c *Client) BuildResty() (error) {
  u, err := url.ParseRequestURI(c.BaseUrl)
  if err != nil {
    return err
  }
  baseUrl := fmt.Sprintf("%s://%s", u.Scheme, u.Host)
  authString := base64.StdEncoding.EncodeToString([]byte(c.Token + ":"))
  c.Client = resty.New().SetHostURL(baseUrl).SetHeader("Content-Type", "application/json").SetHeader("Authorization", fmt.Sprintf("Basic %s", authString)).SetHeader("On-Behalf-Of", strconv.Itoa(c.OnBehalfOf)).SetRetryCount(5)
  return nil
}
