package greenhouse

import (
	"context"
)

func GetAllDegrees(c *Client) (*[]Degree, error) {
	var obj []Degree
	err := MultiGet(c, context.TODO(), "v1/degrees", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetAllDisciplines(c *Client) (*[]Discipline, error) {
	var obj []Discipline
	err := MultiGet(c, context.TODO(), "v1/disciplines", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetAllSchools(c *Client) (*[]School, error) {
	var obj []School
	err := MultiGet(c, context.TODO(), "v1/schools", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
