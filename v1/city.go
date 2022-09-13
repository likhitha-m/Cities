package route

import (
	"cities/controllers"
	"github.com/labstack/echo/v4"
)

func CitiesGroup(e *echo.Group) {

	e.POST("", controllers.CreateCity,)
	e.GET("", controllers.GetCities)
	e.GET("/:cityId", controllers.GetCityById)
	e.DELETE("/:cityId", controllers.DeleteCityById)
	e.PATCH("/:cityId", controllers.UpdateCity)

}
