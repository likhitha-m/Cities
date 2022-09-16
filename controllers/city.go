package controllers

import (
	"fmt"
	"net/http"

	"cities/config"
	"cities/services"
	"cities/types"
	"cities/utils"

	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "golang.org/x/text/message"
)
// Create City godoc
// @Summary Add city
// @Description add a city
// @Tags City
// @Accept  json
// @Produce json
// @Param customer body types.CityPayload true "Add city"
// @Success 200 {object} utils.SuccessContent{data=types.MsgOut}
// @Failure 400 {object} utils.ErrorContent
// @Failure 404 {object} utils.ErrorContent
// @Failure 500 {object} utils.ErrorContent
// @Router /cards [post]
// @Security CustomerBasicAuth


func CreateCity(c echo.Context) error {
	input := &types.CityPayload{}
	if err := c.Bind(input); err != nil {
		logger.Error("func_CreatePassRefund: Error in binding. Error: ", err)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, config.ErrWrongPayload)
	}
	// Validate request body

	if err := utils.ValidateStruct(input); err != nil {
		logger.Error("func_AddGrant: Error in validating request. Error:", err)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, err)
	}

	cr := services.CitiesReceiver{}
	cr.CityPayload = *input
	fmt.Println("cr.CityPayload ", cr.CityPayload)
	err := cr.AddCity()
	if err != nil {
		logger.Error("func_AddCreditsForGuestRecovery:  ", err.Error())
		return utils.HttpErrorResponse(c, http.StatusBadRequest, config.ErrHttpCallInternalServerError)
	}

	return utils.HttpSuccessResponse(c, http.StatusOK, map[string]string{"message": config.MsgCityAdded})
}

func GetCities(c echo.Context) error {
	cr := services.CitiesReceiver{}
	cities, err := cr.GetCities()
	if err != nil {
		logger.Error("func_GetCities:  ", err.Error())
		return utils.HttpErrorResponse(c, http.StatusBadRequest, err)
	}

	return utils.HttpSuccessResponse(c, http.StatusOK, cities)
}

func GetCityById(c echo.Context) error {
	cId := c.Param("cityId")
	cr := services.CitiesReceiver{}

	city, err := cr.GetCityById(cId)
	if err != nil {
		logger.Error("func_AddCreditsForGuestRecovery:  ", err.Error())
		return utils.HttpErrorResponse(c, http.StatusBadRequest, config.ErrRecordNotFound)
	}
	// fmt.Println("city.Id", city.ID)

	return utils.HttpSuccessResponse(c, http.StatusOK, city)
}

func DeleteCityById(c echo.Context) error {
	cId := c.Param("cityId")
	cr := services.CitiesReceiver{}

	err := cr.DeleteCityById(cId)
	if err != nil {
		logger.Error("func_DeleteCityById:  ", err.Error())
		return utils.HttpErrorResponse(c, http.StatusBadRequest, err)
	}

	return utils.HttpSuccessResponse(c, http.StatusOK, map[string]string{"message": config.MsgCityDeleted})
}
func UpdateCity(c echo.Context) error {
	cId := c.Param("cityId")
	cr := services.CitiesReceiver{}
	input := &types.CityPayload{}
	if err := c.Bind(input); err != nil {
		logger.Error("func_UpdateCity: Error in binding. Error: ", err)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, config.ErrWrongPayload)
	}
	// Validate request body

	if err := utils.ValidateStruct(input); err != nil {
		logger.Error("func_UpdateCity: Error in validating request. Error:", err)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, err)
	}

	cr.CityPayload = *input
	err := cr.UpdateCity(cId)
	if err != nil {
		logger.Error("func_UpdateCity:  ", err.Error())
		return utils.HttpErrorResponse(c, http.StatusBadRequest, err)
	}

	return utils.HttpSuccessResponse(c, http.StatusOK, map[string]string{"message": config.MsgCityUpdated})
}
