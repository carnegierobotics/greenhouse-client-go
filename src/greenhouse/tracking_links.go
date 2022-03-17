package greenhouse

import (
  "context"
  "encoding/json"
  "fmt"
)

func GetTrackingLinkData(c *Client, ctx context.Context, token string) (*TrackingLink, error) {
  var obj TrackingLink
  endpoint := fmt.Sprintf("v1/%s/%s", "tracking_links", token)
  resp, err := Get(c, ctx, endpoint)
  if err != nil {
    return nil, err
  }
  err = json.Unmarshal(resp.Body(), &obj)
  if err != nil {
    return nil, err
  }
  return &obj, nil
}
