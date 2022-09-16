package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type AddCity struct {
	Message string `json:"message" example:"city has been successfully added"`
}

type CityPayload struct {
	City  string `json:"City" example:"mangalore"`
	State string `json:"State" example:"karnataka"`
}

type Favourites struct {
	City struct {
		City string `json:"city" bson:"city"`
		State string `json:"state" bson:"state"`
	} `json:"cities"`
	CityId primitive.ObjectID `json:"city_id" bson:"city_id"`
	UserId primitive.ObjectID `json:"user_id" bson:"user_id"`
	Id primitive.ObjectID `json:"id" bson:"_id"`
	

}