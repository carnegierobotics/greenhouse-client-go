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
	"encoding/json"
	"fmt"
	"strings"
)

// GetAllCandidates retrieves a list of all candidates.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-candidates
func GetAllCandidates(c *Client, ctx context.Context) (*[]Candidate, error) {
	//TODO: There are a whole lot of querystring params that need to be implemented here.
	var obj []Candidate
	err := MultiGet(c, ctx, "v1/candidates", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetCandidate retrieves a candidate by ID.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-candidate
func GetCandidate(c *Client, ctx context.Context, id int) (*Candidate, error) {
	var obj Candidate
	endpoint := fmt.Sprintf("v1/candidates/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetCandidateByName searches for a candidate by name.
func GetCandidateByName(c *Client, ctx context.Context, firstName string, lastName string) (*Candidate, error) {
	found := make([]Candidate, 0)
	list, err := GetAllCandidates(c, ctx)
	if err != nil {
		return nil, err
	}
	for _, candidate := range *list {
		if *candidate.FirstName == firstName && *candidate.LastName == lastName {
			found = append(found, candidate)
		}
	}
	if len(found) < 1 {
		return nil, nil
	} else if len(found) > 1 {
		return nil, fmt.Errorf("Found multiple matching candidates.")
	} else {
		return &found[0], nil
	}
}

// DeleteCandidate deletes a candidate by ID.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#delete-delete-candidate
func DeleteCandidate(c *Client, ctx context.Context, id int) error {
	return Delete(c, ctx, fmt.Sprintf("v1/candidates/%d", id), nil)
}

// UpdateCandidate updates a candidate's properties.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#patch-edit-candidate
func UpdateCandidate(c *Client, ctx context.Context, id int, obj *Candidate) error {
	return Update(c, ctx, fmt.Sprintf("v1/candidates/%d", id), obj)
}

// AddAttachmentToCandidate adds an attachment to a candidate.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-attachment
func AddAttachmentToCandidate(c *Client, ctx context.Context, id int, obj *Attachment) error {
	endpoint := fmt.Sprintf("v1/candidates/%d/attachments", id)
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, endpoint, jsonBody)
	return err
}

// CreateCandidate creates a candidate.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-candidate
func CreateCandidate(c *Client, ctx context.Context, obj *Candidate) (int, error) {
	return Create(c, ctx, "v1/candidates", obj)
}

// AddNoteToCandidate adds a note to a candidate.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-note
func AddNoteToCandidate(c *Client, ctx context.Context, id int, note *Note) error {
	endpoint := fmt.Sprintf("v1/candidates/%d/activity_feed/notes", id)
	jsonBody, err := json.Marshal(note)
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, endpoint, jsonBody)
	return err
}

// AddEmailNoteToCandidate adds an email note to a candidate.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-e-mail-note
func AddEmailNoteToCandidate(c *Client, ctx context.Context, id int, email *Email) error {
	endpoint := fmt.Sprintf("v1/candidates/%d/activity_feed/emails", id)
	jsonBody, err := json.Marshal(email)
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, endpoint, jsonBody)
	return err
}

// AddEducationToCandidate adds an education to a candidate.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-education
func AddEducationToCandidate(c *Client, ctx context.Context, id int, education *Education) error {
	endpoint := fmt.Sprintf("v1/candidates/%d/educations", id)
	jsonBody, err := json.Marshal(education)
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, endpoint, jsonBody)
	return err
}

// DeleteEducationFromCandidate deletes an education from a candidate.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#delete-remove-education-from-candidate
func DeleteEducationFromCandidate(c *Client, ctx context.Context, candidateId int, educationId int) error {
	return Delete(c, ctx, fmt.Sprintf("v1/candidates/%d/educations/%d", candidateId, educationId), nil)
}

// AddEmploymentToCandidate adds an employment to a candidate.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-employment
func AddEmploymentToCandidate(c *Client, ctx context.Context, id int, employment *Employment) error {
	endpoint := fmt.Sprintf("v1/candidates/%d/employments", id)
	jsonBody, err := json.Marshal(employment)
	if err != nil {
		return err
	}
	_, err = Post(c, ctx, endpoint, jsonBody)
	return err
}

// DeleteEmploymentFromCandidate deletes an employment from a candidate.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#delete-remove-employment-from-candidate
func DeleteEmploymentFromCandidate(c *Client, ctx context.Context, candidateId int, employmentId int) error {
	return Delete(c, ctx, fmt.Sprintf("v1/candidates/%d/employments/%d", candidateId, employmentId), nil)
}

// CreateProspect creates a prospect.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#post-add-prospect
func CreateProspect(c *Client, ctx context.Context, obj *Candidate) (int, error) {
	return Create(c, ctx, "v1/prospects", obj)
}

// AnonymizeCandidate anonymizes the data associated with a candidate.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#put-anonymize-candidate
func AnonymizeCandidate(c *Client, ctx context.Context, id int, fields []string) (int, error) {
	var obj Candidate
	endpoint := fmt.Sprintf("v1/candidates/%d/anonymize?fields=%s", id, strings.Join(fields, ","))
	resp, err := Put(c, ctx, endpoint, nil)
	if err != nil {
		return *obj.Id, err
	}
	err = json.Unmarshal(resp.Body(), &obj)
	if err != nil {
		return *obj.Id, err
	}
	return *obj.Id, nil
}

// MergeCandidates will merge two candidates.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#put-merge-candidates
func MergeCandidates(c *Client, ctx context.Context, targetId int, duplicateId int) error {
	endpoint := "v1/candidates/merge"
	jsonBody, err := json.Marshal(map[string]int{"primary_candidate_id": targetId, "duplicate_candidate_id": duplicateId})
	if err != nil {
		return err
	}
	_, err = Put(c, ctx, endpoint, jsonBody)
	return err
}
