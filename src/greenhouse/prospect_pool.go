package greenhouse

import (
  "context"
  "fmt"
)

func GetAllProspectPools(c *Client) (*[]ProspectPool, error) {
  var obj []ProspectPool
  err := MultiGet(c, context.TODO(), "v1/prospect_pools", "", &obj)
  if err != nil {
    return nil, err
  }
  return &obj, nil
}

func GetProspectPool(c *Client, id int) (*ProspectPool, error) {
  var obj ProspectPool
  endpoint := fmt.Sprintf("v1/prospect_pools/%d", id)
  err := SingleGet(c, context.TODO(), endpoint, &obj)
  if err != nil {
    return nil, err
  }
  return &obj, nil
}
