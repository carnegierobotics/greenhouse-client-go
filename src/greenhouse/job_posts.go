package greenhouse

import (
	"context"
	"errors"
	"fmt"
	"strconv"
)

func GetAllJobPosts(c *Client, live bool, active bool) (*[]JobPost, error) {
	var obj []JobPost
	querystring := fmt.Sprintf("live=%s&active=%s", strconv.FormatBool(live), strconv.FormatBool(active))
	err := MultiGet(c, context.TODO(), "v1/job_posts", querystring, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetJobPost(c *Client, id int) (*JobPost, error) {
	var obj JobPost
	err := SingleGet(c, context.TODO(), fmt.Sprintf("v1/job_posts/%d", id), &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetAllJobPostsForJob(c *Client, id int) (*[]JobPost, error) {
	var obj []JobPost
	endpoint := fmt.Sprintf("v1/jobs/%d/job_posts", id)
	err := MultiGet(c, context.TODO(), endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetJobPostForJob(c *Client, id int) (*JobPost, error) {
	//This should probably never be implemented since Greenhouse supports multiple posts per job as
	//of 2016.
	return nil, errors.New("GetJobPostForJob not implemented.")
}

func GetCustomLocationsForJobPost(c *Client, id int) (*[]Location, error) {
	var obj []Location
	endpoint := fmt.Sprintf("v1/job_posts/%d/custom_locations", id)
	err := MultiGet(c, context.TODO(), endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func UpdateJobPost(c *Client, id int, obj *JobPost) error {
	return Update(c, context.TODO(), fmt.Sprintf("v2/job_posts/%d", id), obj)
}

func UpdateJobPostStatus(c *Client, id int, status string) error {
	obj := map[string]string{"status": status}
	return Update(c, context.TODO(), fmt.Sprintf("v2/job_posts/%d/status", id), &obj)
}
