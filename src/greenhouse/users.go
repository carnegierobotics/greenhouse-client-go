package greenhouse

import (
  "fmt"
)

type UserBasics struct {
  Id int `json:"id"`
  Name string `json:"name"`
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
  EmployeeId string `json:"employee_id,omitempty"`
}

type User struct {
  UserBasics
  PrimaryEmail string `json:"primary_email_address"`
  UpdatedAt string `json:"updated_at"`
  CreatedAt string `json:"created_at"`
  Disabled bool `json:"disabled"`
  SiteAdmin bool `json:"site_admin"`
  Emails []string `json:"emails"`
  LinkedCandidateIds []int `json:"linked_candidate_ids"`
}

func GetUser(c *Client, id int) (*User, error) {
  var obj User
  err := GetById(c, "users", id, obj)
  if err != nil {
    return nil, err
  }
  return &obj, nil
}

func CreateUser(c *Client, obj *User) error {
  err := Create(c, "users", obj)
  if err != nil {
    return err
  }
  return nil
}

func EnableUser(c *Client, obj *User) error {
  lookupInfo := GetLookupInfo(obj)
  err := Update(c, "users/enable", obj.Id, lookupInfo)
  if err != nil {
    return err
  }
  return nil
}

func DisableUser(c *Client, obj *User) error {
  lookupInfo := GetLookupInfo(obj)
  err := Update(c, "users/disable", obj.Id, lookupInfo)
  if err != nil {
    return err
  }
  return nil
}

func UpdateUser(c *Client, obj *UserBasics) error {
  err := Update(c, "users", obj.Id, obj)
  if err != nil {
    return err
  }
  return nil
}

func GetLookupInfo(obj *User) string {
  return fmt.Sprintf("{\"user\":{\"user_id\":%d}}", obj.Id)
}
