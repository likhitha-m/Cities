package route

import (
	"github.com/labstack/echo/v4"
	"sample-golang/controllers"
)

func FavouritesGroup(e *echo.Group) {

	e.POST("/:cityId", controllers.AddFavouriteCity)
	e.DELETE("/:cityId", controllers.RemoveFavouriteCity)

}
