package route

import (
	"cities/controllers"
	// "cities/middleware"

	"github.com/labstack/echo/v4"
)

func InitializeRoutes(e *echo.Group) {
	e.GET("/health", controllers.HealthCheck)
	//Cities Group
	gCities := e.Group("/cities")
	CitiesGroup(gCities)

	gFavourites := e.Group("/favourites")
	FavouritesGroup(gFavourites)
}
