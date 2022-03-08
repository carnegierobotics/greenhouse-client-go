package greenhouse

import (
	"context"
	"encoding/json"
	"fmt"
)

type HiringTeam struct {
	Name    string         `json:"name"`
	Members []HiringMember `json:"members"`
}

type HiringMember struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	Active      bool   `json:"active"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Name        string `json:"name,omitempty"`
	Responsible bool   `json:"responsible"`
	EmployeeId  string `json:"employee_id,omitempty"`
}

type HiringMemberUpdateInfo struct {
	UserId                           int  `json:"user_id"`
	ResponsibleForFutureCandidates   bool `json:"responsible_for_future_candidates"`
	ResponsibleForActiveCandidates   bool `json:"responsible_for_active_candidates"`
	ResponsibleForInactiveCandidates bool `json:"responsible_for_inactive_candidates"`
}

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
	resp, err := c.Client.R().SetContext(ctx).SetBody(jsonBody).Put(fmt.Sprintf("v1/jobs/%d/hiring_team", id))
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("Got %s for URL: %s, body was %s", resp.Status(), resp.Request.URL, string(jsonBody))
	}
	return nil
}
