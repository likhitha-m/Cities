package types

type AddCity struct {
	Message string `json:"message" example:"city has been successfully added"`
}

type CityPayload struct {
	City  string `json:"City" example:"mangalore"`
	State string `json:"State" example:"karnataka"`
}
