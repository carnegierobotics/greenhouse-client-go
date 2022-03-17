package greenhouse

import (
	"context"
)

func GetAllDegrees(c *Client, ctx context.Context) (*[]Degree, error) {
	var obj []Degree
	err := MultiGet(c, ctx, "v1/degrees", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetAllDisciplines(c *Client, ctx context.Context) (*[]Discipline, error) {
	var obj []Discipline
	err := MultiGet(c, ctx, "v1/disciplines", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetAllSchools(c *Client, ctx context.Context) (*[]School, error) {
	var obj []School
	err := MultiGet(c, ctx, "v1/schools", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
