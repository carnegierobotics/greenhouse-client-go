package utils

import (
  "encoding/json"
  "fmt"
  "strconv"
  "github.com/carnegierobotics/greenhouse-client-go/internal/http"
)

func Update(c *http.Client, itemType string, id int, item BasicType) error {
  jsonBody, err := json.Marshal(item)
  if err != nil {
    return err
  }
  _, err = c.Client.R().SetBody(jsonBody).Patch(fmt.Sprintf("v1/%s/%d", itemType, id))
  return nil
}

type BasicType interface {
  Id int `json:"id"`
}
