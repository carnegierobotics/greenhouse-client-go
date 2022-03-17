package greenhouse

import (
	"context"
	"fmt"
)

func GetAllProspectPools(c *Client, ctx context.Context) (*[]ProspectPool, error) {
	var obj []ProspectPool
	err := MultiGet(c, ctx, "v1/prospect_pools", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetProspectPool(c *Client, ctx context.Context, id int) (*ProspectPool, error) {
	var obj ProspectPool
	endpoint := fmt.Sprintf("v1/prospect_pools/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
