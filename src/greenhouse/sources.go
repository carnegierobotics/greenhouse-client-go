package greenhouse

import (
	"context"
)

func GetAllSources(c *Client, ctx context.Context) (*[]Source, error) {
	var obj []Source
	err := MultiGet(c, ctx, "v1/sources", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
