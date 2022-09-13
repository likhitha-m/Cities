package types

type UserPayload struct {
	Name         string             `json:"name"`
	Email        string             `json:"email"`
	Password     string             `json:"password"`
	MobileNumber string             `json:"mobile"`
	
}