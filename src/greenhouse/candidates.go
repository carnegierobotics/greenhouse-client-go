package greenhouse

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

func GetAllCandidates(c *Client) (*[]Candidate, error) {
	//TODO: There are a whole lot of querystring params that need to be implemented here.
	var obj []Candidate
	err := MultiGet(c, context.TODO(), "v1/candidates", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetCandidate(c *Client, id int) (*Candidate, error) {
	var obj Candidate
	err := SingleGet(c, context.TODO(), "v1/candidates", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func DeleteCandidate(c *Client, id int) error {
	return Delete(c, context.TODO(), fmt.Sprintf("v1/candidates/%d", id), nil)
}

func UpdateCandidate(c *Client, id int, obj *Candidate) error {
	return Update(c, context.TODO(), fmt.Sprintf("v1/candidates/%d", id), obj)
}

func AddAttachmentToCandidate(c *Client, id int, obj *Attachment) error {
	endpoint := fmt.Sprintf("v1/candidates/%d/attachments", id)
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = Post(c, context.TODO(), endpoint, jsonBody)
	return err
}

func CreateCandidate(c *Client, obj *Candidate) (int, error) {
	return Create(c, context.TODO(), "v1/candidates", obj)
}

func AddNoteToCandidate(c *Client, id int, note *Note) error {
	endpoint := fmt.Sprintf("v1/candidates/%d/activity_feed/notes", id)
	jsonBody, err := json.Marshal(note)
	if err != nil {
		return err
	}
	_, err = Post(c, context.TODO(), endpoint, jsonBody)
	return err
}

func AddEmailNoteToCandidate(c *Client, id int, email *Email) error {
	endpoint := fmt.Sprintf("v1/candidates/%d/activity_feed/emails", id)
	jsonBody, err := json.Marshal(email)
	if err != nil {
		return err
	}
	_, err = Post(c, context.TODO(), endpoint, jsonBody)
	return err
}

func AddEducationToCandidate(c *Client, id int, education *Education) error {
	endpoint := fmt.Sprintf("v1/candidates/%d/educations", id)
	jsonBody, err := json.Marshal(education)
	if err != nil {
		return err
	}
	_, err = Post(c, context.TODO(), endpoint, jsonBody)
	return err
}

func DeleteEducationFromCandidate(c *Client, candidateId int, educationId int) error {
	return Delete(c, context.TODO(), fmt.Sprintf("v1/candidates/%d/educations/%d", candidateId, educationId), nil)
}

func AddEmploymentToCandidate(c *Client, id int, employment *Employment) error {
	endpoint := fmt.Sprintf("v1/candidates/%d/employments", id)
	jsonBody, err := json.Marshal(employment)
	if err != nil {
		return err
	}
	_, err = Post(c, context.TODO(), endpoint, jsonBody)
	return err
}

func DeleteEmploymentFromCandidate(c *Client, candidateId int, employmentId int) error {
	return Delete(c, context.TODO(), fmt.Sprintf("v1/candidates/%d/employments/%d", candidateId, employmentId), nil)
}

func CreateProspect(c *Client, obj *Candidate) (int, error) {
	return Create(c, context.TODO(), "v1/prospects", obj)
}

func AnonymizeCandidate(c *Client, id int, fields []string) (int, error) {
	var obj Candidate
	endpoint := fmt.Sprintf("v1/candidates/%d/anonymize?fields=%s", id, strings.Join(fields, ","))
	resp, err := Put(c, context.TODO(), endpoint, nil)
	if err != nil {
		return obj.Id, err
	}
	err = json.Unmarshal(resp.Body(), &obj)
	if err != nil {
		return obj.Id, err
	}
	return obj.Id, nil
}

func MergeCandidates(c *Client, targetId int, duplicateId int) error {
	endpoint := "v1/candidates/merge"
	jsonBody, err := json.Marshal(map[string]int{"primary_candidate_id": targetId, "duplicate_candidate_id": duplicateId})
	if err != nil {
		return err
	}
	_, err = Put(c, context.TODO(), endpoint, jsonBody)
	return err
}
