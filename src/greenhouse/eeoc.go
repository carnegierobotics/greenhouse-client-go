package greenhouse

import (
  "context"
)

func GetAllEEOC(c *Client) (*[]EEOC, error) {
  var obj []EEOC
  endpoint := "v1/eeoc"
  err := MultiGet(c, context.TODO(), endpoint, "", &obj)
  if err != nil {
    return nil, err
  }
  return &obj, nil
}

func GetEEOCForApplication(c *Client, applicationId int) (*EEOC, error) {
  var obj EEOC
  err := SingleGet(c, "applications", applicationId, &obj, context.TODO())
  if err != nil {
    return nil, err
  }
  return &obj, nil
}
