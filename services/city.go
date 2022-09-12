package services

import (
	"context"
	"time"

	"sample-golang/models"
	"sample-golang/types"

	"sample-golang/storage"

	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreditsReceiver struct {
	MDB            *mongo.Database
	CustomerId     int
	OrderDetailsId primitive.ObjectID
	TransactionId  primitive.ObjectID
	CRPoints       float64
	TxnType        string
	CRType         string
	CityPayload    types.CityPayload
	Credits        float64
}

func (cr *CreditsReceiver) AddCity() error {

	cm := models.CityModel{}
	cm.City = cr.CityPayload.City
	cm.State = cr.CityPayload.State
	cm.CreatedAt = time.Now().UTC()
	cm.UpdatedAt = time.Now().UTC()
	mdb := storage.MONGO_DB

	_, err := mdb.Collection(models.CitiesCollection).InsertOne(context.TODO(), cm)
	if err != nil {
		logger.Error("func_AddCredits: ", err)
		return err
	}

	return nil

}
