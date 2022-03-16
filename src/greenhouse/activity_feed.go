package greenhouse

import (
	"context"
	"encoding/json"
	"fmt"
)

func GetActivityFeed(c *Client, id int) (*ActivityFeed, error) {
	var obj ActivityFeed
	resp, err := Get(context.TODO(), c, fmt.Sprintf("v1/candidates/%d/activity_feed", id))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp.Body(), &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
