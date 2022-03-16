package greenhouse

import (
  "context"
)

func GetAllProspectPools(c *Client) (*[]ProspectPool, error) {
  var obj []ProspectPool
  err := PaginatedGet(c, context.TODO(), "v1/prospect_pools", "", &obj)
  if err != nil {
    return nil, err
  }
  return &obj, nil
}

func GetProspectPool(c *Client, id int) (*ProspectPool, error) {
  var obj ProspectPool
  err := GetById(c, "prospect_pools", id, &obj, context.TODO())
  if err != nil {
    return nil, err
  }
  return &obj, nil
}
