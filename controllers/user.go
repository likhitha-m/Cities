package controllers

import (
	"cities/config"
	// "cities/models"
	"cities/services"
	"cities/types"
	"cities/utils"
	"net/http"
	// "strconv"

	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
)
func CreateUser(c echo.Context) error {
	// Initilize schema
	user := &types.UserPayload{}

	// Bind request body into userpayload
	if err := c.Bind(user); err != nil {
		logger.Error("func_CreateCustomer: Error in binding. Error: ", err)
		return utils.HttpErrorResponse(c,  http.StatusBadRequest, config.ErrWrongPayload)
	}
	// Validate request body
	if err := utils.ValidateStruct(user); err != nil {
		logger.Error("func_CreateCustomer: Error in validating request. Error: ", err)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, err)
	}

	validateMobNum := utils.CheckForNumbers(user.MobileNumber)
	if !validateMobNum {
		logger.Error("func_CreateCustomer: Error :", config.ErrInvalidMobNum)
		return utils.HttpErrorResponse(c,  http.StatusBadRequest, config.ErrInvalidMobNum)
	}
	

	


	// Check password format
	isValidPass, err := utils.IsPasswordValid(user.Password)
	if err != nil {
		logger.Error("func_CreateCustomer: is password valid. Error: ", err)
		return utils.HttpErrorResponse(c,  http.StatusBadRequest, err)
	}
	if !isValidPass {
		logger.Error("func_CreateCustomer: error in 'is password valid'")
		return utils.HttpErrorResponse(c, http.StatusBadRequest, config.ErrInvalidPasswordFormat)
	}
	

	_, err = services.GetUserByEmail(user.Email)
	if err == nil {
		logger.Error("func_CreateCustomer: Record found:", err)
		return utils.HttpErrorResponse(c,  utils.GetStatusCode(config.ErrDuplicateCustomer), config.ErrDuplicateCustomer)
	}
	// Create entry in BL - DB table
	_, err = services.CreateUser(user)
	if err != nil {
		logger.Error("func_CreateCustomer: Error in create user:", err)
		return utils.HttpErrorResponse(c, http.StatusBadRequest, err)
	}

	//sendemailverificationlink 
	//sendotp 


	return utils.HttpSuccessResponse(c, http.StatusOK, config.MsgUserAdded)
}
