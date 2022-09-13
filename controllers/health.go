package controllers

import (
	"cities/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

/*
// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags Root
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get] */
func HealthCheck(c echo.Context) error {
	return utils.HttpSuccessResponse(c, http.StatusOK, map[string]interface{}{
		"data": "Server is up and running..",
	})
}
