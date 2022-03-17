package greenhouse

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

func GetAllUsers(c *Client) (*[]User, error) {
	var obj []User
	err := MultiGet(c, context.TODO(), "v1/users", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetUser(c *Client, id int) (*User, error) {
	var obj User
  endpoint := fmt.Sprintf("v1/users/%d", id)
	err := SingleGet(c, context.TODO(), endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateUser(c *Client, obj *UserCreateInfo) (int, error) {
	return Create(c, context.TODO(), "v1/users", obj)
}

func EnableUser(c *Client, id int, ctx context.Context) error {
	lookupInfo := GetLookupInfo(id)
	_, err := Patch(c, ctx, "v2/users/enable", lookupInfo)
	if err != nil {
		return err
	}
	return nil
}

func DisableUser(c *Client, id int, ctx context.Context) error {
	lookupInfo := GetLookupInfo(id)
	_, err := Patch(c, ctx, "v2/users/disable", lookupInfo)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(c *Client, id int, obj *UserUpdateInfo) error {
	return Update(c, context.TODO(), fmt.Sprintf("v1/users/%d", id), obj)
}

func GetLookupInfo(id int) []byte {
	return []byte(fmt.Sprintf("{\"user\":{\"user_id\":%d}}", id))
}

func ChangeUserPermissionLevel(ctx context.Context, c *Client) error {
	return errors.New("ChangeUserPermissionLevel not implemented.")
}

func AddEmailAddressToUser(ctx context.Context, c *Client, userId int, obj *UserEmailUpdateInfo) error {
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, fmt.Sprintf("v1/users/%d/email_addresses", userId), jsonBody)
	if err != nil {
		return err
	}
	return nil
}
