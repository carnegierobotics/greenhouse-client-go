package greenhouse

import ()

type JobPost struct {
  Id int `json:"id"`
  OpeningId string `json:"opening_id"`
  Status string `json:"status"`
  OpenedAt string `json:"opened_at"`
  ClosedAt string `json:"closed_at"`
  ApplicationId int `json:"application_id"`
  CloseReason CloseReason `json:"close_reason"`
}
