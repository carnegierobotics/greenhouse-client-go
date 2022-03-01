package utils

import (
  "encoding/json"
  "errors"
  "fmt"
  "github.com/carnegierobotics/greenhouse-client-go/internal/http"
)

func Create(c *http.Client, itemType string, item interface{}) error {
  jsonBody, err := json.Marshal(item)
  if err != nil {
    return err
  }
  resp, err := c.Client.R().SetBody(jsonBody).Post(fmt.Sprintf("v1/%s", itemType))
  if err != nil {
    return err
  }
  if !resp.IsSuccess() {
    return errors.New(resp.Status())
  }
  return nil
}
