package greenhouse

import (
	"context"
	"fmt"
)

func GetActivityFeed(c *Client, ctx context.Context, id int) (*ActivityFeed, error) {
	var obj ActivityFeed
	endpoint := fmt.Sprintf("v1/candidates/%d/activity_feed", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
