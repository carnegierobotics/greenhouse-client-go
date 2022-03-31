package greenhouse

import (
	"context"
	"encoding/json"
	"fmt"
)

func GetAllJobOpenings(c *Client, ctx context.Context, id int, status string) (*[]JobOpening, error) {
	var obj []JobOpening
	endpoint := fmt.Sprintf("v1/jobs/%d/openings", id)
	querystring := fmt.Sprintf("status=%s", status)
	err := MultiGet(c, ctx, endpoint, querystring, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetJobOpening(c *Client, ctx context.Context, jobId int, openingId int) (*JobOpening, error) {
	var obj JobOpening
	endpoint := fmt.Sprintf("v1/jobs/%d/openings/%d", jobId, openingId)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func DeleteJobOpenings(c *Client, ctx context.Context, jobId int, openingIds []int) error {
	jsonBody, err := json.Marshal(map[string][]int{"ids": openingIds})
	if err != nil {
		return err
	}
	return Delete(c, ctx, fmt.Sprintf("v1/jobs/%d/openings", jobId), jsonBody)
}

func UpdateJobOpenings(c *Client, ctx context.Context, jobId int, openingId int, obj *JobOpeningUpdateInfo) error {
	endpoint := fmt.Sprintf("v1/jobs/%d/openings/%d", jobId, openingId)
	return Update(c, ctx, endpoint, obj)
}

func CreateJobOpenings(c *Client, ctx context.Context, jobId int, obj JobOpeningCreateInfo) (*[]int, error) {
  idList := make([]int, len(obj.Openings), len(obj.Openings))
  type RespObj struct {
    Openings []Opening `json:"openings"`
  }
  var respObj RespObj
	endpoint := fmt.Sprintf("v1/jobs/%d/openings", jobId)
  jsonBody, err := json.Marshal(obj)
  if err != nil {
    return nil, err
  }
  resp, err := Post(c, ctx, endpoint, jsonBody)
  if err != nil {
    return nil, err
  }
  err = json.Unmarshal(resp.Body(), &respObj)
  if err != nil {
    return nil, err
  }
  for i, item := range respObj.Openings {
    idList[i] = item.Id
  }
  return &idList, nil
}
