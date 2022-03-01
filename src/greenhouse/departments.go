package greenhouse

import ()

type Department struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	ParentId int    `json:"parent_id,omitempty"`
	ChildIds []int  `json:"child_ids,omitempty"`
	/* Not in our product tier.
	ParentDepartmentExternalId string `json:"parent_department_external_id"`
	ChildDepartmentExternalIds []string `json:"child_department_external_ids"`
	ExternalId string `json:"external_id"`
	*/
}

type DepartmentCreateInfo struct {
	Name     string `json:"name"`
	ParentId int    `json:"parent_id,omitempty"`
}

type DepartmentUpdateInfo struct {
	Name string `json:"name"`
}

func GetDepartment(c *Client, id int) (*Department, error) {
	var obj Department
	err := GetById(c, "departments", id, obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateDepartment(c *Client, obj *DepartmentCreateInfo) error {
	err := Create(c, "departments", obj)
	if err != nil {
		return err
	}
	return nil
}

func UpdateDepartment(c *Client, id int, obj *DepartmentUpdateInfo) error {
	err := Update(c, "departments", id, obj)
	if err != nil {
		return err
	}
	return nil
}
