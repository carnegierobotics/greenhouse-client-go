package greenhouse

import (
	"context"
	"fmt"
)

func GetAllCandidateTags(c *Client) (*[]CandidateTag, error) {
	var obj []CandidateTag
	err := MultiGet(c, context.TODO(), "v1/tags/candidate", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateCandidateTag(c *Client, obj *CandidateTag) (int, error) {
	return Create(c, context.TODO(), fmt.Sprintf("v1/tags/candidate"), obj)
}

func DeleteCandidateTag(c *Client, tagId int) error {
	//Note that this is not synchronous; it simply kicks off a delete job which may take 24 hours
	//to complete.
	return Delete(c, context.TODO(), fmt.Sprintf("v1/tags/candidate/%d", tagId), nil)
}

func GetTagsForCandidate(c *Client, id int) (*[]CandidateTag, error) {
	var obj []CandidateTag
	endpoint := fmt.Sprintf("v1/candidates/%d/tags", id)
	err := MultiGet(c, context.TODO(), endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func DeleteTagFromCandidate(c *Client, candidateId int, tagId int) error {
	endpoint := fmt.Sprintf("v1/candidates/%d/tags/%d", candidateId, tagId)
	return Delete(c, context.TODO(), endpoint, nil)
}

func CreateTagForCandidate(c *Client, candidateId int, tagId int) error {
	endpoint := fmt.Sprintf("v1/candidates/%d/tags/%d", candidateId, tagId)
	_, err := Put(c, context.TODO(), endpoint, nil)
	return err
}
