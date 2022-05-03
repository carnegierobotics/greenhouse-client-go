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

// GetAllOffers retrieves a list of all offers made by an organization.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-offers
func GetAllOffers(c *Client, ctx context.Context) (*[]Offer, error) {
	var obj []Offer
	err := MultiGet(c, ctx, "v1/offers", "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetAllOffersForApplication retrieves a list of all offers made for an application.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-list-offers-for-application
func GetAllOffersForApplication(c *Client, ctx context.Context, id int) (*[]Offer, error) {
	var obj []Offer
	err := MultiGet(c, ctx, fmt.Sprintf("v1/applications/%d/offers", id), "", &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetCurrentOfferForApplication retrieves the current offer for an application.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-current-offer-for-application
func GetCurrentOfferForApplication(c *Client, ctx context.Context, id int) (*Offer, error) {
	var obj Offer
	endpoint := fmt.Sprintf("v1/applications/%d/offers/current_offer", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetOffer retrieves an offer by ID.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#get-retrieve-offer
func GetOffer(c *Client, ctx context.Context, id int) (*Offer, error) {
	var obj Offer
	endpoint := fmt.Sprintf("v1/offers/%d", id)
	err := SingleGet(c, ctx, endpoint, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// UpdateCurrentOffer updates the current offer for an application.
// Greenhouse API docs: https://developers.greenhouse.io/harvest.html#patch-update-current-offer
func UpdateCurrentOffer(c *Client, ctx context.Context, id int, obj *Offer) error {
	endpoint := fmt.Sprintf("v1/applications/%d/offers/current_offer", id)
	return Update(c, ctx, endpoint, obj)
}
