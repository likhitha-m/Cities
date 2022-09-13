package services

import (
	"context"
	"fmt"
	"time"

	"sample-golang/models"
	"sample-golang/types"

	"sample-golang/storage"

	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CitiesReceiver struct {
	MDB        *mongo.Database
	CustomerId int

	CityPayload types.CityPayload
}

func (cr *CitiesReceiver) AddCity() error {

	cm := models.CityModel{}
	cm.City = cr.CityPayload.City
	cm.State = cr.CityPayload.State
	cm.CreatedAt = time.Now().UTC()
	cm.UpdatedAt = time.Now().UTC()
	mdb := storage.MONGO_DB

	_, err := mdb.Collection(models.CitiesCollection).InsertOne(context.TODO(), cm)
	if err != nil {
		logger.Error("func_AddCity: ", err)
		return err
	}

	return nil

}

func (cr *CitiesReceiver) GetCities() ([]models.CityModel, error) {
	mdb := storage.MONGO_DB
	filter := bson.M{}
	var cities []models.CityModel
	result, err := mdb.Collection(models.CitiesCollection).Find(context.TODO(), filter)

	fmt.Println("result***", result, err)

	if err := result.All(context.Background(), &cities); err != nil {
		logger.Error("func_GetCities: error cur.All() step ", err)
		return nil, err
	}

	fmt.Println("cities", cities, err)

	return cities, nil
}

func (cr *CitiesReceiver) GetCityById(cityId string) (models.CityModel, error) {
	var city models.CityModel
	mdb := storage.MONGO_DB
	cId, err := primitive.ObjectIDFromHex(cityId)
	if err != nil {
		logger.Error("func_GetGrantByOrder", err)

	}
	filter := bson.M{
		"_id": cId,
	}

	result := mdb.Collection(models.CitiesCollection).FindOne(context.TODO(), filter)
	err = result.Decode(&city)
	if err != nil {
		logger.Error("func_S_GetGrant: Error in ", err)
		return city, err
	}
	return city, nil
}

func (cr *CitiesReceiver) DeleteCityById(cityId string) error {

	mdb := storage.MONGO_DB
	cId, err := primitive.ObjectIDFromHex(cityId)
	if err != nil {
		logger.Error("func_GetGrantByOrder", err)

	}
	filter := bson.M{
		"_id": cId,
	}

	_, err = mdb.Collection(models.CitiesCollection).DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CitiesReceiver) UpdateCity(cityId string) error {

	mdb := storage.MONGO_DB
	cId, err := primitive.ObjectIDFromHex(cityId)
	if err != nil {
		logger.Error("func_GetGrantByOrder", err)

	}
	filter := bson.M{
		"_id": cId,
	}
	update := bson.M{"$set": bson.M{"city": cr.CityPayload.City, "state": cr.CityPayload.State}}

	_, err = mdb.Collection(models.CitiesCollection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
