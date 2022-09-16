package models

import (
	// "time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const FavouritesCollection = "favourites"

type FavouriteCityModel struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	// City  string             `bson:"city,omitempty" json:"city,omitempty"`
	// State string             `bson:"state,omitempty" json:"state,omitempty"`
	CityId primitive.ObjectID `bson:"city_id,omitempty" json:"city_id,omitempty"`
	UserId primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`

	//CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
	//UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}
