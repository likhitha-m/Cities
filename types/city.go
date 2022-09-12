package types

type AddCity struct {
	Message string `json:"message" example:"city has been successfully added"`
}

type CityPayload struct {
	City  string `json:"city" example:"mangalore"`
	State string `json:"state" example:"karnataka"`
}
