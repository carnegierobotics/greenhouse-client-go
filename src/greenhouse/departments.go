package greenhouse

import (
  "encoding/json"
  "errors"
  "fmt"
  "github.com/carnegierobotics/greenhouse-client-go/internal/http"
  "github.com/carnegierobotics/greenhouse-client-go/internal/utils"
)

const deptEndpoint = "v1/departments"

type DepartmentBasics struct {
  Name string `json:"name"`
}

type Department struct {
  Id int `json:"id,omitempty"`
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

func CreateDepartment(c *http.Client, deptObj *DepartmentBasics) error {
  jsonBody, err := json.Marshal(*deptObj)
  if err != nil {
    return err
  }
  resp, err := c.Client.R().SetBody(jsonBody).Post(deptEndpoint)
  if err != nil {
    return err
  }
  if !resp.IsSuccess() {
    return errors.New(resp.Status())
  }
  return nil
}
