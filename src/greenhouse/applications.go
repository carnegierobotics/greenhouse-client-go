package greenhouse

import (
	"context"
	"errors"
	"fmt"
)

func GetAllApplications(c *Client) (*[]Application, error) {
	var objList []Application
	err := GetAll(c, "applications", &objList, context.TODO())
	if err != nil {
		return nil, err
	}
	return &objList, nil
}

func GetApplication(c *Client, id int) (*Application, error) {
	var obj Application
	err := GetById(c, "applications", id, &obj, context.TODO())
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func DeleteApplication(c *Client, id int) error {
	err := Delete(context.TODO(), c, fmt.Sprintf("v1/applications/%d", id))
	if err != nil {
		return err
	}
	return nil
}

func AddApplicationToCandidate() error {
	return errors.New("AddApplicationToCandidate not implemented.")
}

func UpdateApplication() error {
	return errors.New("UpdateApplication not implemented.")
}

func AdvanceApplication() error {
	return errors.New("AdvanceApplication not implemented.")
}

func MoveApplicationDifferentJob() error {
	return errors.New("MoveApplicationDifferentJob not implemented.")
}

func MoveApplicationSameJob() error {
	return errors.New("MoveApplicationSameJob not implemented.")
}

func AddAttachmentToApplication() error {
	return errors.New("AddAttachmentToApplication not implemented.")
}

func HireApplication() error {
	return errors.New("HireApplication not implemented.")
}

func RejectApplication() error {
	return errors.New("RejectApplication not implemented.")
}

func UpdateRejectionReason() error {
	return errors.New("UpdateRejectionReason not implemented.")
}

func UnrejectApplication() error {
	return errors.New("UnrejectApplication not implemented.")
}
