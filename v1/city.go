package route

import (
	"github.com/labstack/echo/v4"
	"sample-golang/controllers"
)

func CitiesGroup(e *echo.Group) {

	e.POST("", controllers.CreateCity)
	e.GET("", controllers.GetCities)
	e.GET("/:cityId", controllers.GetCityById)
	e.DELETE("/:cityId", controllers.DeleteCityById)
	e.PATCH("/:cityId", controllers.UpdateCity)

}
