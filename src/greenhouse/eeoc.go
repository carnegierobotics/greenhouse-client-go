package greenhouse

import (
	"context"
	"fmt"
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
	endpoint := fmt.Sprintf("v1/applications/%d", applicationId)
	err := SingleGet(c, context.TODO(), endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
