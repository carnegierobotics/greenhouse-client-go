package greenhouse

import (
	"context"
	"fmt"
)

func GetAllScorecards(c *Client, ctx context.Context) (*[]Scorecard, error) {
	var obj []Scorecard
	err := MultiGet(c, ctx, "v1/scorecards", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetScorecardsForApplication(c *Client, ctx context.Context, id int) (*[]Scorecard, error) {
	var obj []Scorecard
	endpoint := fmt.Sprintf("v1/applications/%d/scorecards", id)
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetScorecard(c *Client, ctx context.Context, id int) (*Scorecard, error) {
	var obj Scorecard
	endpoint := fmt.Sprintf("v1/scorecards/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
