package greenhouse

import (
	"context"
	"fmt"
)

func GetActivityFeed(c *Client, id int) (*ActivityFeed, error) {
	var obj ActivityFeed
  endpoint := fmt.Sprintf("v1/candidates/%d/activity_feed", id)
	err := MultiGet(c, context.TODO(), endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
