package services

import (
	"context"
	"fmt"

	// "fmt"
	// "time"

	"cities/models"
	"cities/types"

	"cities/storage"

	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FavouritesReceiver struct {
	MDB *mongo.Database
	CId int

	CityPayload types.CityPayload
}

func (cr *FavouritesReceiver) AddFavouriteCity(cityID string, userId string) error {
	// var city models.CityModel
	mdb := storage.MONGO_DB
	cId, err := primitive.ObjectIDFromHex(cityID)
	if err != nil {
		logger.Error("func_GetGrantByOrder", err)

	}
	user, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		logger.Error("func_GetGrantByOrder", err)

	}
	favourite := models.FavouriteCityModel{
		CityId: cId,
		UserId: user,
	}
	// filter := bson.M{
	// 	"_id": cId,
	// }
	// favourite := mdb.Collection(models.CitiesCollection).FindOne(context.TODO(), filter)
	_, err = mdb.Collection(models.FavouritesCollection).InsertOne(context.TODO(), favourite)
	// err = result.Decode(&city)
	// if err != nil {
	// 	logger.Error("func_S_GetGrant: Error in ", err)
	// 	return city, err
	// }
	return nil
}

func (cr *FavouritesReceiver) RemoveFavouriteCity(cityId string, userId string) error {

	mdb := storage.MONGO_DB
	cId, err := primitive.ObjectIDFromHex(cityId)
	if err != nil {
		logger.Error("func_GetGrantByOrder", err)

	}
	user, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		logger.Error("func_GetGrantByOrder", err)

	}
	filter := bson.M{
		"city_id": cId,
		"user_id": user,
	}

	_, err = mdb.Collection(models.FavouritesCollection).DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (cr *FavouritesReceiver) ListFavourites(userId string) (interface{}, error) {

	user, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		logger.Error("func_GetGrantByOrder", err)

	}
	match := bson.D{primitive.E{Key: "$match",
		Value: bson.M{"user_id": user}}}

	fmt.Println("match", match)

	// lookup := bson.D{primitive.E{Key: "$lookup", Value: bson.M{"from": "cities",
	// 	"localField":   "city_id",
	// 	"foreignField": "_id",
	// 	"as":           "cities"}}}

	lookup := bson.D{primitive.E{Key: "$lookup",
		Value: bson.D{primitive.E{Key: "from", Value: "cities"},
			primitive.E{Key: "localField", Value: "city_id"},
			primitive.E{Key: "foreignField", Value: "_id"},
			primitive.E{Key: "as", Value: "city"}}}}

	unwind := bson.D{primitive.E{Key: "$unwind",
		Value: "$city"}}
	// group:=bson.D{primitive.E{Key: "$group",
	// Value: bson.M{"_id":"$user_id", "cities":bson.M{"$push":"city"}}}}

	var output []types.Favourites
	mdb := storage.MONGO_DB
	collection := mdb.Collection(models.FavouritesCollection)
	cursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{match, lookup, unwind})
	if err != nil {
		logger.Error("func_GetCredits: ", err)
		return 0, err
	}
	if err = cursor.All(context.TODO(), &output); err != nil {
		logger.Error("func_GetCredits: ", err)
		return 0, err
	}

	fmt.Println(output, cursor)

	// fmt.Println("credits****************", finalCredits)
	return output, nil
}
