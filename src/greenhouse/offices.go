package greenhouse

import (
  "context"
)

type Office struct {
	Id                   int      `json:"id"`
	Name                 string   `json:"name"`
	Location             Location `json:"location"`
	PrimaryContactUserId int      `json:"primary_contact_user_id"`
	ParentId             int      `json:"parent_id"`
	ChildIds             []int    `json:"child_ids"`
	/* Not in our product tier.
	ParentOfficeExternalId string `json:"parent_office_external_id"`
	ChildOfficeExternalIds []string `json:"child_office_external_ids"`
	ExternalId string `json:"external_id"`
	*/
}

type OfficeCreateInfo struct {
	Name                 string `json:"name"`
	Location             string `json:"location,omitempty"`
	PrimaryContactUserId int    `json:"primary_contact_user_id,omitempty"`
	ParentId             int    `json:"parent_id,omitempty"`
	/* Not in out product tier.
	ExternalParentId     string `json:"external_parent_id,omitempty"`
	ExternalId           string `json:"external_id,omitempty"`
	*/
}

type OfficeUpdateInfo struct {
	Name     string `json:"name"`
	Location string `json:"location,omitempty"`
	/* Not in our product tier.
	ExternalId string `json:"external_id,omitempty"`
	*/
}

func GetOffice(c *Client, id int) (*Office, error) {
	var obj Office
	err := GetById(c, "offices", id, &obj, context.TODO())
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateOffice(c *Client, obj *OfficeCreateInfo) (int, error) {
	id, err := Create(c, "offices", obj, context.TODO())
	if err != nil {
		return id, err
	}
	return id, nil
}

func UpdateOffice(c *Client, id int, obj *OfficeUpdateInfo) error {
	err := Update(c, "offices", id, obj, context.TODO())
	if err != nil {
		return err
	}
	return nil
}
