package greenhouse

import (
	"context"
	"fmt"
)

func GetAllScorecards(c *Client) (*[]Scorecard, error) {
	var obj []Scorecard
	err := MultiGet(c, context.TODO(), "v1/scorecards", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetScorecardsForApplication(c *Client, id int) (*[]Scorecard, error) {
	var obj []Scorecard
	endpoint := fmt.Sprintf("v1/applications/%d/scorecards", id)
	err := MultiGet(c, context.TODO(), endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetScorecard(c *Client, id int) (*Scorecard, error) {
	var obj Scorecard
	endpoint := fmt.Sprintf("v1/scorecards/%d", id)
	err := SingleGet(c, context.TODO(), endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
