package greenhouse

import (
	"context"
	"fmt"
)

func GetAllOffers(c *Client) (*[]Offer, error) {
	var obj []Offer
	err := MultiGet(c, context.TODO(), "v1/offers", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetAllOffersForApplication(c *Client, id int) (*[]Offer, error) {
	var obj []Offer
	err := MultiGet(c, context.TODO(), fmt.Sprintf("v1/applications/%d/offers", id), "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetCurrentOfferForApplication(c *Client, id int) (*Offer, error) {
	var obj Offer
	endpoint := fmt.Sprintf("v1/applications/%d/offers/current_offer", id)
	err := SingleGet(c, context.TODO(), endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetOffer(c *Client, id int) (*Offer, error) {
	var obj Offer
	endpoint := fmt.Sprintf("v1/offers/%d", id)
	err := SingleGet(c, context.TODO(), endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func UpdateCurrentOffer(c *Client, id int, obj *Offer) error {
	endpoint := fmt.Sprintf("v1/applications/%d/offers/current_offer", id)
	return Update(c, context.TODO(), endpoint, obj)
}
