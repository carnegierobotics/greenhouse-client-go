package greenhouse

import (
	"encoding/base64"
  //"errors"
	"fmt"
	// "encoding/json"
	"github.com/go-resty/resty/v2"
  "net/http"
	"net/url"
	"strconv"
  "time"
)

type Client struct {
	BaseUrl    string
	Token      string
	OnBehalfOf int
	Client     *resty.Client
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
    SetRetryCount(5).
    SetRetryWaitTime(5 * time.Second).
    SetRetryMaxWaitTime(30 * time.Second)
	return nil
}
