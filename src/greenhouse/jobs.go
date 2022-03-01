package greenhouse

import()

type Job struct {
  Id int `json:"id"`
  Name string `json:"name"`
  RequisitionId string `json:"requisition_id"`
  Notes string `json:"notes"`
  Confidential bool `json:"confidential"`
  Status string `json:"status"`
  CreatedAt string `json:"created_at"`
  OpenedAt string `json:"opened_at"`
  ClosedAt string `json:"closed_at"`
  UpdatedAt string `json:"updated_at"`
  IsTemplate bool `json:"is_template"`
  CopiedFromId int `json:"copied_from_id"`
  Departments []Department `json:"departments"`
  Offices []Office `json:"offices"`
  CustomFields map[string]interface{} `json:"custom_fields"`
  KeyedCustomFields map[string]interface{} `json:"keyed_custom_fields"`
  HiringTeam map[string][]HiringMember `json:"hiring_team"`
  Openings []JobOpening `json:"openings"`
}

type HiringMember struct {
  UserBasics
  responsible bool `json:"responsible"`
}
