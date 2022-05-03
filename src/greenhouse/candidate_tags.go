/*
Copyright 2021-2022
Carnegie Robotics, LLC
4501 Hatfield Street, Pittsburgh, PA 15201
https://www.carnegierobotics.com
All rights reserved.

This file is part of greenhouse-client-go.

greenhouse-client-go is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

greenhouse-client-go is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with greenhouse-client-go. If not, see <https://www.gnu.org/licenses/>.
*/
package greenhouse

import (
	"context"
	"fmt"
)

// GetAllCandidateTags lists all candidate tags.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-candidate-tags
func GetAllCandidateTags(c *Client, ctx context.Context) (*[]CandidateTag, error) {
	var obj []CandidateTag
	err := MultiGet(c, ctx, "v1/tags/candidate", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetCandidateTag retrieves a tag by ID.
func GetCandidateTag(c *Client, ctx context.Context, id int) (*CandidateTag, error) {
	list, err := GetAllCandidateTags(c, ctx)
	if err != nil {
		return nil, err
	}
	for _, item := range *list {
		if *item.Id == id {
			return &item, nil
		}
	}
	return nil, nil
}

// CreateCandidateTag creates a tag.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-new-candidate-tag
func CreateCandidateTag(c *Client, ctx context.Context, obj *CandidateTag) (int, error) {
	return Create(c, ctx, fmt.Sprintf("v1/tags/candidate"), obj)
}

// DeleteCandidateTag deletes a tag. Note that this is not synchronous; it simply kicks off a
// delete job which may take 24 hours to complete.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#delete-destroy-a-candidate-tag
func DeleteCandidateTag(c *Client, ctx context.Context, tagId int) error {
	return Delete(c, ctx, fmt.Sprintf("v1/tags/candidate/%d", tagId), nil)
}

// GetTagsForCandidate retrieves all tags associated with a candidate.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-tags-applied-to-candidate
func GetTagsForCandidate(c *Client, ctx context.Context, id int) (*[]CandidateTag, error) {
	var obj []CandidateTag
	endpoint := fmt.Sprintf("v1/candidates/%d/tags", id)
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// DeleteTagFromCandidate deletes a tag from a candidate.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#delete-remove-tag-from-candidate
func DeleteTagFromCandidate(c *Client, ctx context.Context, candidateId int, tagId int) error {
	endpoint := fmt.Sprintf("v1/candidates/%d/tags/%d", candidateId, tagId)
	return Delete(c, ctx, endpoint, nil)
}

// CreateTagForCandidate adds a tag to a candidate.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#put-add-a-candidate-tag
func CreateTagForCandidate(c *Client, ctx context.Context, candidateId int, tagId int) error {
	endpoint := fmt.Sprintf("v1/candidates/%d/tags/%d", candidateId, tagId)
	_, err := Put(c, ctx, endpoint, nil)
	return err
}
