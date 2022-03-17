package greenhouse

import (
	"context"
	"encoding/json"
	"fmt"
)

func GetJobHiringTeam(c *Client, id int) (*HiringTeam, error) {
	resp, err := c.Client.R().SetContext(context.TODO()).Get(fmt.Sprintf("v1/jobs/%d/hiring_team", id))
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("Error getting hiring team: %s", err.Error())
	}
	var obj HiringTeam
	err = json.Unmarshal(resp.Body(), &obj)
	return &obj, nil
}

func UpdateJobHiringTeam(c *Client, id int, obj *map[string][]HiringMemberUpdateInfo, ctx context.Context) error {
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = Put(c, ctx, fmt.Sprintf("v1/jobs/%d/hiring_team", id), jsonBody)
	return err
}

//This is a more atomic operation; instead of replacing an entire team, you can work on one member.
func UpdateHiringTeamMembers(c *Client, ctx context.Context, id int, obj *map[string][]HiringMemberUpdateInfo) error {
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, fmt.Sprintf("v1/jobs/%d/hiring_team", id), jsonBody)
	return err
}

//This is a more atomic operation; instead of replacing an entire team, you can work on one member.
func DeleteHiringTeamMembers(c *Client, ctx context.Context, id int, obj *map[string][]int) error {
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	err = Delete(c, ctx, fmt.Sprintf("v1/jobs/%d/hiring_team", id), jsonBody)
	return err
}
