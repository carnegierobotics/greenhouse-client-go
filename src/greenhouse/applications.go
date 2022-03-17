package greenhouse

import (
	"context"
	"encoding/json"
	"fmt"
)

func GetAllApplications(c *Client) (*[]Application, error) {
	var objList []Application
	err := MultiGet(c, context.TODO(), "v1/applications", "", &objList)
	if err != nil {
		return nil, err
	}
	return &objList, nil
}

func GetApplication(c *Client, id int) (*Application, error) {
	var obj Application
	err := SingleGet(c, context.TODO(), fmt.Sprintf("v1/applications/%d", id), &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func DeleteApplication(c *Client, id int) error {
	return Delete(c, context.TODO(), fmt.Sprintf("v1/applications/%d", id), nil)
}

func AddApplicationToCandidate(c *Client, id int, obj *Application) error {
	endpoint := fmt.Sprintf("v1/candidates/%d/applications", id)
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = Post(c, context.TODO(), endpoint, jsonBody)
	return err
}

func UpdateApplication(c *Client, id int, obj *Application) error {
	endpoint := fmt.Sprintf("v1/applications/%d", id)
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = Patch(c, context.TODO(), endpoint, jsonBody)
	return err
}

func AdvanceApplication(c *Client, id int, from int) error {
	endpoint := fmt.Sprintf("v1/applications/%d/advance", id)
	jsonBody, err := json.Marshal(map[string]int{"from_stage_id": from})
	if err != nil {
		return err
	}
	_, err = Post(c, context.TODO(), endpoint, jsonBody)
	return err
}

func MoveApplicationDifferentJob(c *Client, id int, newJob int, newStage int) error {
	endpoint := fmt.Sprintf("v1/applications/%d/transfer_to_job", id)
	jsonBody, err := json.Marshal(map[string]int{"new_job_id": newJob, "new_stage_id": newStage})
	if err != nil {
		return err
	}
	_, err = Post(c, context.TODO(), endpoint, jsonBody)
	return err
}

func MoveApplicationSameJob(c *Client, id int, from int, to int) error {
	endpoint := fmt.Sprintf("v1/applications/%d/move", id)
	jsonBody, err := json.Marshal(map[string]int{"from_stage_id": from, "to_stage_id": to})
	if err != nil {
		return err
	}
	_, err = Post(c, context.TODO(), endpoint, jsonBody)
	return err
}

func AddAttachmentToApplication(c *Client, id int, item *Attachment) error {
	endpoint := fmt.Sprintf("v1/applications/%d/attachments", id)
	jsonBody, err := json.Marshal(item)
	if err != nil {
		return err
	}
	_, err = Post(c, context.TODO(), endpoint, jsonBody)
	return err
}

func HireApplication(c *Client, id int, obj *ApplicationHire) error {
	endpoint := fmt.Sprintf("v1/applications/%d/hire", id)
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = Post(c, context.TODO(), endpoint, jsonBody)
	return err
}

func RejectApplication(c *Client, id int, obj *ApplicationReject) error {
	endpoint := fmt.Sprintf("v1/applications/%d/reject", id)
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = Post(c, context.TODO(), endpoint, jsonBody)
	return err
}

func UpdateRejectionReason(c *Client, id int, rejectionReason int) error {
	endpoint := fmt.Sprintf("v1/applications/%d/reject", id)
	jsonBody, err := json.Marshal(map[string]int{"rejection_reason_id": rejectionReason})
	if err != nil {
		return err
	}
	_, err = Patch(c, context.TODO(), endpoint, jsonBody)
	return err
}

func UnrejectApplication(c *Client, id int) error {
	endpoint := fmt.Sprintf("v1/applications/%d/unreject", id)
	_, err := Post(c, context.TODO(), endpoint, nil)
	return err
}
