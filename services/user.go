package services

import (
	"context"
	"os"
	"time"

	// "fmt"
	// "time"

	"cities/models"
	"cities/types"
	"cities/utils"

	// "cities/types"

	"cities/storage"

	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

type UsersReceiver struct {
	// MDB        *mongo.Database
	// CustomerId int

	// UserPayload types.CityPayload
}

func GetUserByEmail(email string) (models.UserModel, error) {
	var user models.UserModel
	mdb := storage.MONGO_DB

	filter := bson.M{
		"email": email,
	}

	result := mdb.Collection(models.UsersCollection).FindOne(context.TODO(), filter)
	err := result.Decode(&user)
	if err != nil {
		logger.Error("func_S_GetGrant: Error in ", err)
		return user, err
	}
	return user, nil
}
func CreateUser(c *types.UserPayload) (interface{}, error) {
	um := models.UserModel{}
	um.Email = utils.ToLowerCase(c.Email)
	encPassword, err := utils.Encrypt(c.Password, os.Getenv("PASSWORD_ENC_KEY"))
	if err != nil {
		logger.Error("func_CreateCustomer: Error in encrypt password: ", err)
		return nil, err
	}
	um.Password = encPassword
	um.Name = c.Name
	um.MobileNumber = c.MobileNumber
	um.CreatedAt = time.Now().UTC()
	um.UpdatedAt = time.Now().UTC()
	mdb := storage.MONGO_DB
	_, err = mdb.Collection(models.UsersCollection).InsertOne(context.TODO(), um)
	if err != nil {
		logger.Error("func_AddCity: ", err)
		return nil, err
	}

	return nil, nil
}
