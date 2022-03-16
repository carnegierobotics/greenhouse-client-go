package greenhouse

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

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
	_, err := Patch(ctx, c, "v2/users/enable", lookupInfo)
	if err != nil {
		return err
	}
	return nil
}

func DisableUser(c *Client, id int, ctx context.Context) error {
	lookupInfo := GetLookupInfo(id)
	_, err := Patch(ctx, c, "v2/users/disable", lookupInfo)
	if err != nil {
		return err
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
