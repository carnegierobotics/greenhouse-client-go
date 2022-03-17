package greenhouse

import (
	"context"
	"fmt"
	"strconv"
)

func GetAllScheduledInterviews(c *Client, actionable bool) (*[]ScheduledInterview, error) {
	var obj []ScheduledInterview
	querystring := fmt.Sprintf("&actionable=%s", strconv.FormatBool(actionable))
	endpoint := fmt.Sprintf("v1/scheduled_interviews?%s", querystring)
	err := MultiGet(c, context.TODO(), endpoint, querystring, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetScheduledInterviewsForApplication(c *Client, id int, actionable bool) (*[]ScheduledInterview, error) {
	var obj []ScheduledInterview
	querystring := fmt.Sprintf("&actionable=%s", strconv.FormatBool(actionable))
	endpoint := fmt.Sprintf("v1/applications/%d/scheduled_interviews%s", id, querystring)
	err := MultiGet(c, context.TODO(), endpoint, querystring, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetScheduledInterview(c *Client, id int) (*ScheduledInterview, error) {
	var obj ScheduledInterview
	endpoint := fmt.Sprintf("v1/scheduled_interviews/%d", id)
	err := SingleGet(c, context.TODO(), endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateScheduledInterview(c *Client, obj *ScheduledInterview) (int, error) {
	return Create(c, context.TODO(), "v2/scheduled_interviews", obj)
}

func UpdateScheduledInterview(c *Client, id int, obj *ScheduledInterview) error {
	return Update(c, context.TODO(), fmt.Sprintf("v2/scheduled_interviews/%d", id), obj)
}

func DeleteScheduledInterview(c *Client, id int) error {
	return Delete(c, context.TODO(), fmt.Sprintf("v1/scheduled_interviews/%d", id), nil)
}
