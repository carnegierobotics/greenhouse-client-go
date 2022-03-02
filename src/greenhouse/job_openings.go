package greenhouse

import ()

type JobOpening struct {
	Id            int               `json:"id"`
	OpeningId     string            `json:"opening_id"`
	Status        string            `json:"status"`
	OpenedAt      string            `json:"opened_at"`
	ClosedAt      string            `json:"closed_at"`
	ApplicationId int               `json:"application_id"`
	CloseReason   CloseReason       `json:"close_reason"`
	CustomFields  map[string]string `json:"custom_fields"`
}

type JobOpeningCreateInfo struct {
	Openings []Opening `json:"openings"`
}

type Opening struct {
	Id           int                 `json:"opening_id"`
	CustomFields []map[string]string `json:"custom_fields,omitempty"`
}

type JobOpeningUpdateInfo struct {
	Id            int                 `json:"opening_id,omitempty"`
	Status        string              `json:"status,omitempty"`
	CloseReasonId int                 `json:"close_reason_id,omitempty"`
	CustomFields  []map[string]string `json:"custom_fields,omitempty"`
}
