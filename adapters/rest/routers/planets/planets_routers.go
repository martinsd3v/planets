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
	routes.POST("", planets.Create(), auth.Authorize)
	routes.PATCH("/:UUID", planets.Update(), auth.Authorize)
	routes.DELETE("/:UUID", planets.Destroy(), auth.Authorize)
}
