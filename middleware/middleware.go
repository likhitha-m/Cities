package middleware

import (
	"cities/config"
	// "strconv"
	// "cities/services"
	"cities/utils"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
)

func ValidateCustomerToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := ExtractCustomerTokenID(c)
		fmt.Println("err--------------------------------------------------------", )
		if err != nil {
			return utils.HttpErrorResponse(c, http.StatusUnauthorized, config.ErrHttpCallUnauthorized)
		}
		return next(c)
	}
}
func ExtractCustomerTokenID(c echo.Context) (string, error) {
	

	tokenString := c.Request().Header.Get("Authorization")
	fmt.Println("tokenString", tokenString)

	
	if tokenString == "" {
		return "", config.ErrTokenMissing
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.Error("func_ValidateCustomerToken: Error in jwt token method. Error: ")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("CUSTOMER_JWT_SECRET_KEY")), nil
	})

	fmt.Println("token:", token)
	if err != nil {
		logger.Error("func_ValidateCustomerToken: Error in jwt parse. Error: ", err)
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println("claims", claims)
	fmt.Println("token", token.Valid)
	if ok && token.Valid {
		uid := fmt.Sprintf("%v", claims["user_id"])

		
		c.Request().Header.Set("userId", uid)
		
	}
	fmt.Println("no error")
	return "", nil

}
