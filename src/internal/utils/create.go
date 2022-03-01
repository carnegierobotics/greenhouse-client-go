package utils

import (
  //"encoding/json"
  "fmt"
  "github.com/carnegierobotics/greenhouse-client-go/internal/http"
)

func Create(c *http.Client, itemType string, item interface{}) error {
  _, err := c.Client.R().SetBody(item).Post(fmt.Sprintf("v1/%s", itemType))
  if err != nil {
    return err
  }
  return nil
}
