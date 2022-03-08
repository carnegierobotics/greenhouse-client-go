package greenhouse

import (
	"errors"
)

type JobPost struct {
	Id                       int                   `json:"id"`
	Active                   bool                  `json:"active"`
	Live                     bool                  `json:"live"`
	FirstPublishedAt         string                `json:"first_published_at"`
	Internal                 bool                  `json:"internal"`
	External                 bool                  `json:"external"`
	JobId                    int                   `json:"job_id"`
	Content                  string                `json:"content"`
	InternalContent          string                `json:"internal_content"`
	UpdatedAt                string                `json:"updated_at"`
	CreatedAt                string                `json:"created_at"`
	DemographicQuestionSetId int                   `json:"demographic_question_set_id"`
	Questions                []DemographicQuestion `json:"questions"`
}

type JobPostUpdateInfo struct {
}

func GetAllJobPosts(c *Client) (*[]JobPost, error) {
	return nil, errors.New("GetAllJobPosts not implemented.")
}

func GetJobPost(c *Client, postId int) (*JobPost, error) {
	return nil, errors.New("GetJobPost not implemented.")
}

func GetAllJobPostsForJob(c *Client, jobId int) (*[]JobPost, error) {
	return nil, errors.New("GetAllJobPostsForJob not implemented.")
}

func GetJobPostForJob(c *Client, jobId int) (*JobPost, error) {
	//This should probably never be implemented since Greenhouse supports multiple posts per job as
	//of 2016.
	return nil, errors.New("GetJobPostForJob not implemented.")
}

func GetCustomLocationsForJobPost(c *Client, postId int) (*[]Location, error) {
	return nil, errors.New("GetCustomLocationsForJobPost not implemented.")
}

func UpdateJobPost(c *Client, obj *JobPost) error {
	return errors.New("UpdateJobPost not implemented.")
}

func UpdateJobPostStatus(c *Client, postId int, status string) error {
	return errors.New("UpdateJobPostStatus not implemented.")
}
