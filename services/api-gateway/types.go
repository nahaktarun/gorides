package main

import "ride-sharing/shared/types"

type previewTripRequest struct {
	UserID      string
	Pickup      types.Coordinate
	Destination types.Coordinate
}
