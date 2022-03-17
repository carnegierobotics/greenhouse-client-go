package greenhouse

import (
	"context"
	"fmt"
)

func GetAllCandidateTags(c *Client, ctx context.Context) (*[]CandidateTag, error) {
	var obj []CandidateTag
	err := MultiGet(c, ctx, "v1/tags/candidate", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func CreateCandidateTag(c *Client, ctx context.Context, obj *CandidateTag) (int, error) {
	return Create(c, ctx, fmt.Sprintf("v1/tags/candidate"), obj)
}

func DeleteCandidateTag(c *Client, ctx context.Context, tagId int) error {
	//Note that this is not synchronous; it simply kicks off a delete job which may take 24 hours
	//to complete.
	return Delete(c, ctx, fmt.Sprintf("v1/tags/candidate/%d", tagId), nil)
}

func GetTagsForCandidate(c *Client, ctx context.Context, id int) (*[]CandidateTag, error) {
	var obj []CandidateTag
	endpoint := fmt.Sprintf("v1/candidates/%d/tags", id)
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func DeleteTagFromCandidate(c *Client, ctx context.Context, candidateId int, tagId int) error {
	endpoint := fmt.Sprintf("v1/candidates/%d/tags/%d", candidateId, tagId)
	return Delete(c, ctx, endpoint, nil)
}

func CreateTagForCandidate(c *Client, ctx context.Context, candidateId int, tagId int) error {
	endpoint := fmt.Sprintf("v1/candidates/%d/tags/%d", candidateId, tagId)
	_, err := Put(c, ctx, endpoint, nil)
	return err
}
