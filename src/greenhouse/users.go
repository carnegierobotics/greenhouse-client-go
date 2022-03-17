package greenhouse

import (
	"context"
	"encoding/json"
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
	return err
}

func DisableUser(c *Client, id int, ctx context.Context) error {
	lookupInfo := GetLookupInfo(id)
	_, err := Patch(c, ctx, "v2/users/disable", lookupInfo)
	return err
}

func UpdateUser(c *Client, id int, obj *UserUpdateInfo) error {
	return Update(c, context.TODO(), fmt.Sprintf("v1/users/%d", id), obj)
}

func GetLookupInfo(id int) []byte {
	return []byte(fmt.Sprintf("{\"user\":{\"user_id\":%d}}", id))
}

func ChangeUserPermissionLevel(c *Client, ctx context.Context, user *User, level string) error {
	endpoint := fmt.Sprintf("v1/users/permission_level")
	jsonBody, err := json.Marshal(map[string]interface{}{"user": user, "level": level})
	if err != nil {
		return err
	}
	_, err = Patch(c, context.TODO(), endpoint, jsonBody)
	return err
}

func AddEmailAddressToUser(ctx context.Context, c *Client, userId int, obj *UserEmailUpdateInfo) error {
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, fmt.Sprintf("v1/users/%d/email_addresses", userId), jsonBody)
	return err
}
