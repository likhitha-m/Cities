package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const UsersCollection = "users"

type UserModel struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string             `bson:"name,omitempty" json:"name,omitempty"`
	Email        string             `bson:"email,omitempty" json:"email,omitempty"`
	Password     string             `bson:"password,omitempty" json:"password,omitempty"`
	MobileNumber string             `bson:"mobile,omitempty" json:"mobile,omitempty"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
	IsVerified   bool               `bson:"is_verified" json:"is_verified"`
	
}
