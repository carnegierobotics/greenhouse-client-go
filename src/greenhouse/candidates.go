package greenhouse

import (
	"context"
	"errors"
	"fmt"
)

func GetAllCandidates() error {
	return errors.New("GetAllCandidates not implemented.")
}

func GetCandidate() error {
	return errors.New("GetCandidate not implemented.")
}

func DeleteCandidate(c *Client, id int) error {
	return Delete(c, context.TODO(), fmt.Sprintf("v1/candidates/%d", id), nil)
}

func UpdateCandidate() error {
	return errors.New("UpdateCandidate not implemented.")
}

func AddAttachmentToCandidate() error {
	return errors.New("AddAttachmentToCandidate not implemented.")
}

func CreateCandidate() error {
	return errors.New("CreateCandidate not implemented.")
}

func AddNoteToCandidate() error {
	return errors.New("AddNoteToCandidate not implemented.")
}

func AddEmailNoteToCandidate() error {
	return errors.New("AddEmailNoteToCandidate not implemented.")
}

func AddEducationToCandidate() error {
	return errors.New("AddEducationToCandidate not implemented.")
}

func DeleteEducationFromCandidate(c *Client, candidateId int, educationId int) error {
	return Delete(c, context.TODO(), fmt.Sprintf("v1/candidates/%d/educations/%d", candidateId, educationId), nil)
}

func AddEmploymentToCandidate() error {
	return errors.New("AddEmploymentToCandidate not implemented.")
}

func DeleteEmploymentFromCandidate(c *Client, candidateId int, employmentId int) error {
	return Delete(c, context.TODO(), fmt.Sprintf("v1/candidates/%d/employments/%d", candidateId, employmentId), nil)
}

func CreateProspect() error {
	return errors.New("CreateProspect not implemented.")
}

func AnonymizeCandidate() error {
	return errors.New("Anonymize candidate not implemented.")
}

func MergeCandidates() error {
	return errors.New("MergeCandidates not implemented.")
}
