package planets

import (
	"github.com/labstack/echo/v4"
	"github.com/martinsd3v/planets/adapters/web/controllers/planets"
	"github.com/martinsd3v/planets/adapters/web/middlewares/auth"
)

//SetupRoutes ...
func SetupRoutes(Echo *echo.Echo) {
	controller := planets.Controller{}

	//Public routes no authentication required
	routes := Echo.Group("/planets")

	routes.GET("", controller.Index())
	routes.GET("/:UUID", controller.Show())

	//Private routes authentication is required
	AuthMiddleware := auth.Auth{}
	routes.Use(AuthMiddleware.Authorize)

	routes.POST("", controller.Create())
	routes.PATCH("/:UUID", controller.Update())
	routes.DELETE("/:UUID", controller.Destroy())
}
