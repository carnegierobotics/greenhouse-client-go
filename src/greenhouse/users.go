package greenhouse

import (
	"fmt"
)

type User struct {
	Id                 int    `json:"id"`
	Name               string `json:"name"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	EmployeeId         string `json:"employee_id,omitempty"`
	PrimaryEmail       string   `json:"primary_email_address"`
	UpdatedAt          string   `json:"updated_at"`
	CreatedAt          string   `json:"created_at"`
	Disabled           bool     `json:"disabled"`
	SiteAdmin          bool     `json:"site_admin"`
	Emails             []string `json:"emails"`
	LinkedCandidateIds []int    `json:"linked_candidate_ids"`
}

type UserCreateInfo struct {
  FirstName string `json:"first_name"`
  LastName  string `json:"last_name"`
  Email     string `json:"email"`
  SendEmail bool   `json:"send_email_invite,omitempty"`
}

type UserUpdateInfo struct {
  FirstName string `json:"first_name,omitempty"`
  LastName  string `json:"last_name,omitempty"`
}

func GetUser(c *Client, id int) (*User, error) {
	var obj User
	err := GetById(c, "users", id, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateUser(c *Client, obj *UserCreateInfo) (int, error) {
	id, err := Create(c, "users", obj)
	if err != nil {
		return id, err
	}
	return id, nil
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

func UpdateUser(c *Client, id int, obj *UserUpdateInfo) error {
	err := Update(c, "users", id, obj)
	if err != nil {
		return err
	}
	return nil
}

func GetLookupInfo(obj *User) string {
	return fmt.Sprintf("{\"user\":{\"user_id\":%d}}", obj.Id)
}
