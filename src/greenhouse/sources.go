package greenhouse

import (
  "context"
)

func GetAllSources(c *Client) (*[]Source, error) {
  var obj []Source
  err := PaginatedGet(c, context.TODO(), "v1/sources", "", &obj)
  if err != nil {
    return nil, err
  }
  return &obj, nil
}
