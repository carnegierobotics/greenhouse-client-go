package greenhouse

import (
	"context"
	"fmt"
)

func GetAllEEOC(c *Client, ctx context.Context) (*[]EEOC, error) {
	var obj []EEOC
	endpoint := "v1/eeoc"
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetEEOCForApplication(c *Client, ctx context.Context, applicationId int) (*EEOC, error) {
	var obj EEOC
	endpoint := fmt.Sprintf("v1/applications/%d", applicationId)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
