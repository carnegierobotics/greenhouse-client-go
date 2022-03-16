package greenhouse

import (
  "context"
)

func GetAllSources(c *Client) (*[]Source, error) {
  var obj []Source
  err := GetAll(c, "sources", &obj, context.TODO())
  if err != nil {
    return &obj, err
  }
  return &obj, nil
}
