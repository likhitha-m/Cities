package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserPayload struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	MobileNumber string `json:"mobile"`
}
type VerifySignupEmailTemplate struct {
	VerifySignupLink string `validate:"required"`
}
type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutput struct {
	Email string `json:"email"`
	Token string `json:"token"`
	Id    primitive.ObjectID `json:"id"`
	Name string `json:"name"`
}
