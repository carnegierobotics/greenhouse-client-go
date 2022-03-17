package greenhouse

import (
	"context"
	"encoding/json"
	"errors"
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
	resp, err := Put(c, ctx, fmt.Sprintf("v1/jobs/%d/hiring_team", id), jsonBody)
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("Got %s for URL: %s, body was %s", resp.Status(), resp.Request.URL, string(jsonBody))
	}
	return nil
}

//This is a more atomic operation; instead of replacing an entire team, you can work on one member.
func UpdateHiringTeamMembers() error {
	return errors.New("UpdateHiringTeamMembers not implemented.")
}

//This is a more atomic operation; instead of replacing an entire team, you can work on one member.
func DeleteHiringTeamMembers() error {
	return errors.New("DeleteHiringTeamMembers not implemented.")
}
