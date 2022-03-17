package greenhouse

import (
	"context"
	"fmt"
	"strconv"
)

func GetAllRejectionReasons(c *Client, ctx context.Context, include_defaults bool, per_page int) (*[]RejectionReason, error) {
	var obj []RejectionReason
	querystring := fmt.Sprintf("per_page=%d&include_defaults=%s", per_page, strconv.FormatBool(include_defaults))
	endpoint := fmt.Sprintf("v1/rejection_reasons?%s", querystring)
	err := MultiGet(c, ctx, endpoint, querystring, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
