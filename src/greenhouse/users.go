package greenhouse

import (
	"context"
	"encoding/json"
	"fmt"
)

func GetAllUsers(c *Client, ctx context.Context) (*[]User, error) {
	var obj []User
	err := MultiGet(c, ctx, "v1/users", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetUser(c *Client, ctx context.Context, id int) (*User, error) {
	var obj User
	endpoint := fmt.Sprintf("v1/users/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateUser(c *Client, ctx context.Context, obj *UserCreateInfo) (int, error) {
	return Create(c, ctx, "v1/users", obj)
}

func EnableUser(c *Client, ctx context.Context, id int) error {
	lookupInfo := GetLookupInfo(id)
	_, err := Patch(c, ctx, "v2/users/enable", lookupInfo)
	return err
}

func DisableUser(c *Client, ctx context.Context, id int) error {
	lookupInfo := GetLookupInfo(id)
	_, err := Patch(c, ctx, "v2/users/disable", lookupInfo)
	return err
}

func UpdateUser(c *Client, ctx context.Context, id int, obj *UserUpdateInfo) error {
	return Update(c, ctx, fmt.Sprintf("v1/users/%d", id), obj)
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
	_, err = Patch(c, ctx, endpoint, jsonBody)
	return err
}

func AddEmailAddressToUser(c *Client, ctx context.Context, userId int, obj *UserEmailUpdateInfo) error {
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, fmt.Sprintf("v1/users/%d/email_addresses", userId), jsonBody)
	return err
}
