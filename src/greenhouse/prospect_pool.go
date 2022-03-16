package greenhouse

import (
  "context"
)

func GetAllProspectPools(c *Client) (*[]ProspectPool, error) {
  var obj []ProspectPool
  err := GetAll(c, "prospect_pools", &obj, context.TODO())
  if err != nil {
    return &obj, err
  }
  return &obj, nil
}

func GetProspectPool(c *Client, id int) (*ProspectPool, error) {
  var obj ProspectPool
  err := GetById(c, "prospect_pools", id, &obj, context.TODO())
  if err != nil {
    return &obj, err
  }
  return &obj, nil
}
