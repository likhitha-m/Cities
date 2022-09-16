package config

import "errors"

const (
	MsgCityAdded     = "City has been added"
	MsgCityDeleted   = "City has been successfully deleted"
	MsgCityUpdated   = "City has been updated"
	MsgFavAdded      = "City has been added to favourites"
	MsgFavRemoved    = "City has been removed from favourites"
	MsgUserAdded     = "User has been added"
	MsgEmailVerified = "Email has been verified"
)

var (
	ErrMissingBasicAuth            = errors.New("Authorization must be required in header")
	ErrWrongPayload                = errors.New("Wrong payload, please try again")
	ErrRecordNotFound              = errors.New("Record not found")
	ErrParameterMissing            = errors.New("Parameter missing")
	ErrTokenMissing                = errors.New("Error token missing")
	ErrInvalidHashKey              = errors.New("Invalid hash key")
	ErrInvalidHttpMethod           = errors.New("Invalid http method")
	ErrHttpCallBadRequest          = errors.New("Bad request")
	ErrHttpCallUnauthorized        = errors.New("Unauthorized")
	ErrHttpCallNotFound            = errors.New("Call not found")
	ErrHttpCallInternalServerError = errors.New("Internal server error")
	ErrWentWrong                   = errors.New("Something went wrong")
	ErrInvalidMobNum               = errors.New("Invalid mobile number")
	ErrInvalidPasswordFormat       = errors.New("Invalid password format")
	ErrDuplicateCustomer           = errors.New("User already exists with this email address")
	ErrVerKeyNotFound              = errors.New("verify key not found")
	ErrEmailAlreadyVerified        = errors.New("Email already verified")
	ErrUserDoesNotExist            = errors.New("User does not exist with this email address")
	ErrEmailNotVerified 		   = errors.New(" Email not verified")
	ErrInvalidToken 		   = errors.New("Invalid token")
)
