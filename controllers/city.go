package controllers

import (
	"fmt"
	"net/http"

	"sample-golang/config"
	"sample-golang/services"
	"sample-golang/types"
	"sample-golang/utils"

	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
	// "golang.org/x/text/message"
)

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

	cr := services.CreditsReceiver{}
	cr.CityPayload = *input
	fmt.Println("cr.CityPayload ", cr.CityPayload)
	err := cr.AddCity()
	if err != nil {
		logger.Error("func_AddCreditsForGuestRecovery:  ", err.Error())
		return utils.HttpErrorResponse(c, http.StatusBadRequest, err)
	}

	return utils.HttpSuccessResponse(c, http.StatusOK, map[string]string{"message": config.MsgCityAdded})
}
