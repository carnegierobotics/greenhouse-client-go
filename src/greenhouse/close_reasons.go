package greenhouse

import ()

type CloseReason struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetCloseReason(c *Client, id int) (*CloseReason, error) {
	var obj CloseReason
	err := GetById(c, "close_reasons", id, obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
