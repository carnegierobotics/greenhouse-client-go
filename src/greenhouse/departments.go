package greenhouse

import (
  "github.com/carnegierobotics/greenhouse-client-go/internal/http"
  "github.com/carnegierobotics/greenhouse-client-go/internal/utils"
)

type Department struct {
  Id int `json:"id"`
  Name string `json:"name"`
  ParentId int `json:"parent_id,omitempty"`
  ChildIds []int `json:"child_ids,omitempty"`
  /* Not in our product tier.
  ParentDepartmentExternalId string `json:"parent_department_external_id"`
  ChildDepartmentExternalIds []string `json:"child_department_external_ids"`
  ExternalId string `json:"external_id"`
  */
}

func GetDepartment(c *http.Client, id int) (*Department, error) {
  var obj Department
  err := utils.GetById(c, "departments", id, obj)
  if err != nil {
    return nil, err
  }
  return &obj, nil
}

func CreateDepartment(c *http.Client, deptObj *Department) error {
  err := utils.Create(c, "departments", deptObj)
  if err != nil {
    return err
  }
  return nil
}

func UpdateDepartment(c *http.Client, deptObj *Department) error {
  err := utils.Update(c, "departments", deptObj.Id, deptObj)
  if err != nil {
    return err
  }
  return nil
}
