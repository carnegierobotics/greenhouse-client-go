package greenhouse

import (
  "github.com/carnegierobotics/greenhouse-client-go/internal/http"
  "github.com/carnegierobotics/greenhouse-client-go/internal/utils"
)

type CloseReason struct {
  Id int `json:"id"`
  Name string `json:"name"`
}

func GetCloseReason(c *http.Client, id int) (*CloseReason, error) {
  var obj CloseReason
  err := utils.GetById(c, "close_reasons", id, obj)
  if err != nil {
    return nil, err
  }
  return &obj, nil
}
