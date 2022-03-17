package greenhouse

import (
	"context"
	"fmt"
)

func GetAllEmailTemplates(c *Client, ctx context.Context) (*[]EmailTemplate, error) {
	var obj []EmailTemplate
	err := MultiGet(c, ctx, "v1/email_templates", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetEmailTemplate(c *Client, ctx context.Context, id int) (*EmailTemplate, error) {
	var obj EmailTemplate
	err := SingleGet(c, ctx, fmt.Sprintf("v1/email_templates/%d", id), &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
