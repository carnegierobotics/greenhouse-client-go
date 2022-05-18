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

// GetAllDemographicQuestionSets retrieves a list of all question sets.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-demographic-question-sets
func GetAllDemographicQuestionSets(c *Client, ctx context.Context) (*[]DemographicQuestionSet, error) {
	var obj []DemographicQuestionSet
	err := MultiGet(c, ctx, "v1/demographics/question_sets", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetDemographicQuestionSet retrieves a question set by ID.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-demographic-question-set
func GetDemographicQuestionSet(c *Client, ctx context.Context, id int) (*DemographicQuestionSet, error) {
	var obj DemographicQuestionSet
	endpoint := fmt.Sprintf("v1/demographics/question_sets/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetAllDemographicQuestions retrieves a list of questions.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-demographic-questions
func GetAllDemographicQuestions(c *Client, ctx context.Context) (*[]DemographicQuestion, error) {
	var obj []DemographicQuestion
	err := MultiGet(c, ctx, "v1/demographics/questions", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetAllDemographicQuestionsForQuestionSet retrieves a list of questions for a question set.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-demographic-questions-for-demographic-question-set
func GetAllDemographicQuestionsForQuestionSet(c *Client, ctx context.Context, id int) (*[]DemographicQuestion, error) {
	var obj []DemographicQuestion
	endpoint := fmt.Sprintf("v1/demographics/question_sets/%d/questions", id)
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetDemographicQuestion retrieves a question by ID.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-demographic-question
func GetDemographicQuestion(c *Client, ctx context.Context, id int) (*DemographicQuestion, error) {
	var obj DemographicQuestion
	endpoint := fmt.Sprintf("v1/demographics/questions/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetAllDemographicAnswerOptions retrieves a list of answer options.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-demographic-answer-options
func GetAllDemographicAnswerOptions(c *Client, ctx context.Context) (*[]DemographicAnswerOption, error) {
	var obj []DemographicAnswerOption
	endpoint := fmt.Sprintf("v1/demographics/answer_options")
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetAllDemographicAnswerOptionsForQuestion retrieves a list of answer options for a question.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-demographic-answer-options-for-demographic-question
func GetAllDemographicAnswerOptionsForQuestion(c *Client, ctx context.Context, id int) (*[]DemographicAnswerOption, error) {
	var obj []DemographicAnswerOption
	endpoint := fmt.Sprintf("v1/demographics/questions/%d/answer_options", id)
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetDemographicAnswerOption retrieves an answer option by ID.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-demographic-answer-option
func GetDemographicAnswerOption(c *Client, ctx context.Context, id int) (*DemographicAnswerOption, error) {
	var obj DemographicAnswerOption
	endpoint := fmt.Sprintf("v1/demographics/answer_options/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetAllDemographicAnswers retrieves a list of all answers.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-demographic-answers
func GetAllDemographicAnswers(c *Client, ctx context.Context) (*[]DemographicAnswer, error) {
	var obj []DemographicAnswer
	err := MultiGet(c, ctx, "v1/demographics/answers", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetAllDemographicAnswersForApplication retrieves a list of answers in an application.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-demographic-answers-for-application
func GetAllDemographicAnswersForApplication(c *Client, ctx context.Context) (*[]DemographicAnswer, error) {
	var obj []DemographicAnswer
	endpoint := "v1/applications/%d/demographics/answers"
	err := MultiGet(c, ctx, endpoint, "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetDemographicAnswer gets an answer by ID.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-demographic-answer
func GetDemographicAnswer(c *Client, ctx context.Context, id int) (*DemographicAnswer, error) {
	var obj DemographicAnswer
	endpoint := fmt.Sprintf("v1/demographics/answers/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
