package greenhouse

import (
	"context"
	"fmt"
	"strconv"
)

func GetAllScheduledInterviews(c *Client, ctx context.Context, actionable bool) (*[]ScheduledInterview, error) {
	var obj []ScheduledInterview
	querystring := fmt.Sprintf("&actionable=%s", strconv.FormatBool(actionable))
	endpoint := fmt.Sprintf("v1/scheduled_interviews?%s", querystring)
	err := MultiGet(c, ctx, endpoint, querystring, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetScheduledInterviewsForApplication(c *Client, ctx context.Context, id int, actionable bool) (*[]ScheduledInterview, error) {
	var obj []ScheduledInterview
	querystring := fmt.Sprintf("&actionable=%s", strconv.FormatBool(actionable))
	endpoint := fmt.Sprintf("v1/applications/%d/scheduled_interviews%s", id, querystring)
	err := MultiGet(c, ctx, endpoint, querystring, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetScheduledInterview(c *Client, ctx context.Context, id int) (*ScheduledInterview, error) {
	var obj ScheduledInterview
	endpoint := fmt.Sprintf("v1/scheduled_interviews/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateScheduledInterview(c *Client, ctx context.Context, obj *ScheduledInterviewCreateInfo) (int, error) {
	return Create(c, ctx, "v2/scheduled_interviews", obj)
}

func UpdateScheduledInterview(c *Client, ctx context.Context, id int, obj *ScheduledInterviewUpdateInfo) error {
	return Update(c, ctx, fmt.Sprintf("v2/scheduled_interviews/%d", id), obj)
}

func DeleteScheduledInterview(c *Client, ctx context.Context, id int) error {
	return Delete(c, ctx, fmt.Sprintf("v1/scheduled_interviews/%d", id), nil)
}
