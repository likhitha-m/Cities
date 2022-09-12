package route

import (
	"sample-golang/controllers"
	// "sample-golang/middleware"

	"github.com/labstack/echo/v4"
)

func InitializeRoutes(e *echo.Group) {
	e.GET("/health", controllers.HealthCheck)
	//Credits Group
	gCities := e.Group("/cities")
	CitiesGroup(gCities)
}
