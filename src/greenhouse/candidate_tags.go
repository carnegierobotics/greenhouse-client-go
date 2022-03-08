package greenhouse

import (
	"errors"
)

type CandidateTag struct {
	Id   int    `json:"id"`
	Name string `json:"string"`
}

func GetAllCandidateTags(c *Client) (*[]CandidateTag, error) {
	return nil, errors.New("GetAllCandidateTags not implemented.")
}

func CreateCandidateTag(c *Client, tagName string) (int, error) {
	return 0, errors.New("CreateCandidateTag not implemented.")
}

func DeleteCandidateTag(c *Client, tagId int) error {
	//Note that this is not synchronous; it simply kicks off a delete job which may take 24 hours
	//to complete.
	return errors.New("DeleteCandidateTag not implemented.")
}
