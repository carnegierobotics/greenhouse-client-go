package greenhouse

import ()

type Department struct {
  Id int `json:"id"`
  Name string `json:"name"`
  ParentId int `json:"parent_id"`
  ParentDepartmentExternalId int `json:"parent_department_external_id"`
  ChildIds []int `json:"child_ids"`
  ChildDepartmentExternalIds []string `json:"child_department_external_ids"`
  ExternalId string `json:"external_id"`
}
