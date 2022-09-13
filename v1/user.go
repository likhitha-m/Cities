package route

import (
	"cities/controllers"
	"github.com/labstack/echo/v4"
)

func UsersGroup(e *echo.Group) {

	e.POST("/signup", controllers.CreateUser)
	//verifyemail
	//verifyotp
	//login
	

}
