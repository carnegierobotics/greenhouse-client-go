package greenhouse

import ()

type CustomField struct {
	Id                  int                 `json:"id"`
	Name                string              `json:"name"`
	FieldType           string              `json:"field_type"`
	Active              bool                `json:"active"`
	Priority            int                 `json:"priority"`
	ValueType           string              `json:"value_type"`
	Private             bool                `json:"private"`
	Required            bool                `json:"required"`
	RequireApproval     bool                `json:"require_approval"`
	TriggerNewVersion   bool                `json:"trigger_new_version"`
	NameKey             string              `json:"name_key"`
	CustomFieldOptions  []CustomFieldOption `json:"custom_field_options"`
	Description         string              `json:"description"`
	ExposeInJobBoardAPI bool                `json:"expose_in_job_board_api"`
	ApiOnly             bool                `json:"api_only"`
	Offices             []Office            `json:"offices"`
	Departments         []Department        `json:"departments"`
	TemplateTokenString string              `json:"template_token_string"`
}

type CustomFieldOption struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Priority   int    `json:"priority"`
	ExternalId string `json:"external_id"`
}

func GetCustomField() error {
	return nil
}

func CreateCustomField() error {
	return nil
}

func UpdateCustomField() error {
	return nil
}

func DeleteCustomField() error {
	return nil
}

func CreateCustomFieldOption() error {
	return nil
}

func UpdateCustomFieldOption() error {
	return nil
}

func DeleteCustomFieldOption(c *Client, ids []int) error {
	return nil
}
