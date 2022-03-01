package greenhouse

import (
  "github.com/carnegierobotics/greenhouse-client-go/internal/http"
  "github.com/carnegierobotics/greenhouse-client-go/internal/utils"
)

type UserBasics struct {
  Id int `json:"id"`
  Name string `json:"name"`
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
  EmployeeId string `json:"employee_id"`
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

func GetUser(c *http.Client, id int) (*User, error) {
  var obj User
  err := utils.GetById(c, "users", id, obj)
  if err != nil {
    return nil, err
  }
  return &obj, nil
}

func CreateUser(c *http.Client, obj *User) error {
  err := utils.Create(c, "users", obj)
  if err != nil {
    return err
  }
  return nil
}

func EnableUser(c *http.Client, obj *User) error {
  lookupInfo := GetLookupInfo(obj)
  err := utils.Update(c, "users/enable", lookupInfo)
  if err != nil {
    return err
  }
  return nil
}

func DisableUser(c *http.Client, obj *User) error {
  lookupInfo := GetLookupInfo(obj)
  err := utils.Update(c, "users/disable", lookupInfo)
  if err != nil {
    return err
  }
  return nil
}

func UpdateUser(c *http.Client, obj *UserBasics) error {
  err := utils.Update(c, "users", obj)
  if err != nil {
    return err
  }
  return nil
}

func GetLookupInfo(obj *User) string {
  return fmt.Sprintf("{\"user\":{\"user_id\":%d}}", obj.Id)
}
