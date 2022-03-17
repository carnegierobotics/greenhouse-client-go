package greenhouse

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

func GetAllCustomFields(c *Client, fieldType string, includeInactive bool) (*[]CustomField, error) {
	var obj []CustomField
	endpoint := "v1/custom_fields"
	if fieldType != "" {
		endpoint = fmt.Sprintf("%s/%s", endpoint, fieldType)
	}
	querystring := fmt.Sprintf("include_inactive=%s", strconv.FormatBool(includeInactive))
	err := MultiGet(c, context.TODO(), endpoint, querystring, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetCustomField(c *Client, id int) (*CustomField, error) {
	var obj CustomField
	err := SingleGet(c, context.TODO(), fmt.Sprintf("v1/custom_field/%d", id), &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateCustomField(c *Client, obj *CustomField) (int, error) {
	return Create(c, context.TODO(), "v1/custom_fields", obj)
}

func UpdateCustomField(c *Client, id int, obj *CustomField) error {
	return Update(c, context.TODO(), fmt.Sprintf("v1/custom_fields/%d", id), obj)
}

func DeleteCustomField(c *Client, id int) error {
	return Delete(c, context.TODO(), fmt.Sprintf("v1/custom_fields/%d", id), nil)
}

func GetCustomFieldOptions(c *Client, id int, typeStr string) (*[]CustomFieldOption, error) {
	var obj []CustomFieldOption
	endpoint := fmt.Sprintf("v1/custom_field/%d/custom_field_options", id)
	querystring := fmt.Sprintf("type=%s", typeStr)
	err := MultiGet(c, context.TODO(), endpoint, querystring, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateCustomFieldOptions(c *Client, id int, obj *[]CustomFieldOption) (int, error) {
	endpoint := fmt.Sprintf("v1/custom_field/%d/custom_field_options", id)
	return Create(c, context.TODO(), endpoint, obj)
}

func UpdateCustomFieldOptions(c *Client, id int, obj *[]CustomFieldOption) error {
	endpoint := fmt.Sprintf("v1/custom_field/%d/custom_field_options", id)
	return Update(c, context.TODO(), endpoint, obj)
}

func DeleteCustomFieldOptions(c *Client, id int, ids []int) error {
	jsonBody, err := json.Marshal(map[string][]int{"option_ids": ids})
	if err != nil {
		return err
	}
	return Delete(c, context.TODO(), fmt.Sprintf("v1/custom_field/%d/custom_field_options", id), jsonBody)
}
