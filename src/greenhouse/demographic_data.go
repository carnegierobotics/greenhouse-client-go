package greenhouse

import (
	"context"
	"fmt"
)

func GetAllDemographicQuestionSets(c *Client) (*[]DemographicQuestionSet, error) {
	var obj []DemographicQuestionSet
	err := MultiGet(c, context.TODO(), "v1/demographics/question_sets", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetDemographicQuestionSet(c *Client, id int) (*DemographicQuestionSet, error) {
	var obj DemographicQuestionSet
	endpoint := fmt.Sprintf("v1/demographics/question_sets/%d", id)
	err := SingleGet(c, context.TODO(), endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetAllDemographicQuestions(c *Client) (*[]DemographicQuestion, error) {
	var obj []DemographicQuestion
	err := MultiGet(c, context.TODO(), "v1/demographics/questions", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetAllDemographicQuestionsForQuestionSet(c *Client, id int) (*[]DemographicQuestion, error) {
	var obj []DemographicQuestion
	endpoint := fmt.Sprintf("v1/demographics/question_sets/%d/questions", id)
	err := MultiGet(c, context.TODO(), endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetDemographicQuestion(c *Client, id int) (*DemographicQuestion, error) {
	var obj DemographicQuestion
	endpoint := fmt.Sprintf("v1/demographics/questions/%d", id)
	err := SingleGet(c, context.TODO(), endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetAllDemographicAnswerOptions(c *Client) (*[]DemographicAnswerOption, error) {
	var obj []DemographicAnswerOption
	endpoint := fmt.Sprintf("v1/demographics/answer_options")
	err := MultiGet(c, context.TODO(), endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetAllDemographicAnswerOptionsForQuestion(c *Client, id int) (*[]DemographicAnswerOption, error) {
	var obj []DemographicAnswerOption
	endpoint := fmt.Sprintf("v1/demographics/questions/%d/answer_options", id)
	err := MultiGet(c, context.TODO(), endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetDemographicAnswerOption(c *Client, id int) (*DemographicAnswerOption, error) {
	var obj DemographicAnswerOption
	endpoint := fmt.Sprintf("v1/demographics/answer_options/%d", id)
	err := SingleGet(c, context.TODO(), endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetAllDemographicAnswers(c *Client) (*[]DemographicAnswer, error) {
	var obj []DemographicAnswer
	err := MultiGet(c, context.TODO(), "v1/demographics/answers", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetAllDemographicAnswersForApplication(c *Client) (*[]DemographicAnswer, error) {
	var obj []DemographicAnswer
	endpoint := "v1/applications/%d/demographics/answers"
	err := MultiGet(c, context.TODO(), endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetDemographicAnswer(c *Client, id int) (*DemographicAnswer, error) {
	var obj DemographicAnswer
	endpoint := fmt.Sprintf("v1/demographics/answers/%d", id)
	err := SingleGet(c, context.TODO(), endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
