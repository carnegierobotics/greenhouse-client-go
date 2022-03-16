package greenhouse

import (
  "context"
  "encoding/json"
  "fmt"
)

func GetAllRejectionReasons(c *Client, include_default bool, per_page int) (*[]RejectionReason, error) {
  var obj []RejectionReason
  endpoint := "v1/rejection_reasons"
  var custom_query string
  if include_default {
    custom_query = "&include_defaults=true"
  }
  endpoint = fmt.Sprintf("%s?per_page=%d%s", endpoint, per_page, custom_query)
  body, err := PaginatedGet(c, context.TODO(), endpoint, custom_query)
  if err != nil {
    return &obj, err
  }
  err = json.Unmarshal(body, &obj)
  if err != nil {
    return &obj, err
  }
  return &obj, nil
}
