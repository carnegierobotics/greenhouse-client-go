package utils

import (
  "encoding/json"
  "fmt"
  "github.com/carnegierobotics/greenhouse-client-go/internal/http"
)

func Update(c *http.Client, itemType string, id int, item interface{}) error {
  jsonBody, err := json.Marshal(item)
  if err != nil {
    return err
  }
  _, err = c.Client.R().SetBody(jsonBody).Patch(fmt.Sprintf("v1/%s/%d", itemType, id))
  return nil
}
