package route

import (
	"sample-golang/controllers"
	"github.com/labstack/echo/v4"
)

func CitiesGroup(e *echo.Group) {

	e.POST("", controllers.CreateCity)
	
}
