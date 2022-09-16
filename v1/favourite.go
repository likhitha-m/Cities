package route

import (
	"cities/controllers"

	"github.com/labstack/echo/v4"
	"cities/middleware"
)

func FavouritesGroup(e *echo.Group) {

	e.POST("/:cityId", controllers.AddFavouriteCity, middleware.ValidateCustomerToken)
	e.DELETE("/:cityId", controllers.RemoveFavouriteCity, middleware.ValidateCustomerToken)
	e.GET("",controllers.ListFavourites, middleware.ValidateCustomerToken)

}
