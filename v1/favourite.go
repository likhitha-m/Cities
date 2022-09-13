package route

import (
	"cities/controllers"
	"github.com/labstack/echo/v4"
)

func FavouritesGroup(e *echo.Group) {

	e.POST("/:cityId", controllers.AddFavouriteCity, )
	e.DELETE("/:cityId", controllers.RemoveFavouriteCity)

}
