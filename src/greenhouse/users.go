package greenhouse

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

type User struct {
	Id                 int      `json:"id"`
	Name               string   `json:"name"`
	FirstName          string   `json:"first_name"`
	LastName           string   `json:"last_name"`
	EmployeeId         string   `json:"employee_id,omitempty"`
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

type UserEmailUpdateInfo struct {
	Email            string `json:"email"`
	SendVerification bool   `json:"send_verification"`
}

func GetAllUsers(c *Client) (*[]User, error) {
	var obj []User
	err := GetAll(c, "users", &obj, context.TODO())
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetUser(c *Client, id int) (*User, error) {
	var obj User
	err := GetById(c, "users", id, &obj, context.TODO())
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateUser(c *Client, obj *UserCreateInfo) (int, error) {
	id, err := Create(c, "users", obj, context.TODO())
	if err != nil {
		return id, err
	}
	return id, nil
}

func EnableUser(c *Client, id int, ctx context.Context) error {
	lookupInfo := GetLookupInfo(id)
	resp, err := c.Client.R().SetContext(ctx).SetBody(lookupInfo).Patch("v2/users/enable")
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return errors.New(resp.Status())
	}
	return nil
}

func DisableUser(c *Client, id int, ctx context.Context) error {
	lookupInfo := GetLookupInfo(id)
	resp, err := c.Client.R().SetContext(ctx).SetBody(lookupInfo).Patch("v2/users/disable")
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return errors.New(resp.Status())
	}
	return nil
}

func UpdateUser(c *Client, id int, obj *UserUpdateInfo) error {
	err := Update(c, "users", id, obj, context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func GetLookupInfo(id int) string {
	return fmt.Sprintf("{\"user\":{\"user_id\":%d}}", id)
}

func AddEmailAddressToUser(ctx context.Context, c *Client, userId int, obj *UserEmailUpdateInfo) error {
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	resp, err := c.Client.R().SetContext(ctx).SetBody(jsonBody).Post(fmt.Sprintf("v1/users/%d/email_addresses", userId))
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("Error occurred adding email address to user: %s", resp.Status())
	}
	return nil
}
