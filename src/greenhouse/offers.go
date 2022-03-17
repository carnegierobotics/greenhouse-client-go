package greenhouse

import (
	"context"
	"fmt"
)

func GetAllOffers(c *Client, ctx context.Context) (*[]Offer, error) {
	var obj []Offer
	err := MultiGet(c, ctx, "v1/offers", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetAllOffersForApplication(c *Client, ctx context.Context, id int) (*[]Offer, error) {
	var obj []Offer
	err := MultiGet(c, ctx, fmt.Sprintf("v1/applications/%d/offers", id), "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetCurrentOfferForApplication(c *Client, ctx context.Context, id int) (*Offer, error) {
	var obj Offer
	endpoint := fmt.Sprintf("v1/applications/%d/offers/current_offer", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetOffer(c *Client, ctx context.Context, id int) (*Offer, error) {
	var obj Offer
	endpoint := fmt.Sprintf("v1/offers/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func UpdateCurrentOffer(c *Client, ctx context.Context, id int, obj *Offer) error {
	endpoint := fmt.Sprintf("v1/applications/%d/offers/current_offer", id)
	return Update(c, ctx, endpoint, obj)
}
