package controllers

import (
	// "fmt"
	"net/http"

	"cities/config"
	"cities/services"
	// "cities/types"
	"cities/utils"

	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "golang.org/x/text/message"
)

func AddFavouriteCity(c echo.Context) error {
	cId := c.Param("cityId")
	cr := services.FavouritesReceiver{}
	err := cr.AddFavouriteCity(cId)
	if err != nil {
		logger.Error("func_AddCreditsForGuestRecovery:  ", err.Error())
		return utils.HttpErrorResponse(c, http.StatusBadRequest, config.ErrRecordNotFound)
	}
	return utils.HttpSuccessResponse(c, http.StatusOK, nil)
}

func RemoveFavouriteCity(c echo.Context) error {
	cId := c.Param("cityId")
	cr := services.FavouritesReceiver{}

	err := cr.RemoveFavouriteCity(cId)
	if err != nil {
		logger.Error("func_DeleteCityById:  ", err.Error())
		return utils.HttpErrorResponse(c, http.StatusBadRequest, err)
	}

	return utils.HttpSuccessResponse(c, http.StatusOK, map[string]string{"message": config.MsgFavRemoved})
}
