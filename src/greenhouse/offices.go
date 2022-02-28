package greenhouse

import()

type Office struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Location map[string]string `json:"location"`
  PrimaryContactUserId int `json:"primary_contact_user_id"`
  ParentId int `json:"parent_id"`
  ParentOfficeExternalId string `json:"parent_office_external_id"`
  ChildIds []int `json:"child_ids"`
  ChildOfficeExternalIds []string `json:"child_office_external_ids"`
  ExternalId string `json:"external_id"`
}
