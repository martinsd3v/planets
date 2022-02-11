package planets

import (
	"github.com/labstack/echo/v4"
	"github.com/martinsd3v/planets/adapters/rest/controllers/planets"
	"github.com/martinsd3v/planets/adapters/rest/middlewares/auth"
)

//SetupRoutes ...
func SetupRoutes(Echo *echo.Echo) {
	//Public routes no authentication required
	routes := Echo.Group("/planets")

	routes.GET("/:UUID", planets.Show())
	routes.GET("", planets.Index())

	//Private routes authentication is required
	AuthMiddleware := auth.Auth{}
	routes.Use(AuthMiddleware.Authorize)

	routes.POST("", planets.Create())
	routes.PATCH("/:UUID", planets.Update())
	routes.DELETE("/:UUID", planets.Destroy())
}
