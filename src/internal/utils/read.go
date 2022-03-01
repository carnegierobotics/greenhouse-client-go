package utils

import (
  "encoding/json"
  "fmt"
  "github.com/carnegierobotics/greenhouse-client-go/internal/http"
)

func GetById(c *http.Client, itemType string, id int, item interface{}) error {
  resp, err := c.Client.R().Get(fmt.Sprintf("v1/%s/%d", itemType, id))
  if err != nil {
    return err
  }
  err = json.Unmarshal(resp.Body(), &item)
  if err != nil {
    return err
  }
  return nil
}

func GetAll(c *http.Client, itemType string, itemList interface{}) error {
  resp, err := c.Client.R().Get(fmt.Sprintf("v1/%s", itemType))
  if err != nil {
    return err
  }
  err = json.Unmarshal(resp.Body(), &itemList)
  if err != nil {
    return err
  }
  return nil
}
