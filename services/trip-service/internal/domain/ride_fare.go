package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RideFareModel struct {
	ID                primitive.ObjectID
	UserID            string
	PackageSlug       string // e.g. van, laxury, sedan
	TotalPriceInCents float64
	ExpiresAt         time.Time
}
