package greenhouse

import (
	"context"
	"fmt"
)

func GetAllEmailTemplates(c *Client) (*[]EmailTemplate, error) {
	var obj []EmailTemplate
	err := MultiGet(c, context.TODO(), "v1/email_templates", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetEmailTemplate(c *Client, id int) (*EmailTemplate, error) {
	var obj EmailTemplate
	err := SingleGet(c, context.TODO(), fmt.Sprintf("v1/email_templates/%d", id), &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
