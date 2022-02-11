package users

import (
	"github.com/labstack/echo/v4"
	"github.com/martinsd3v/planets/adapters/rest/controllers/users"
	"github.com/martinsd3v/planets/adapters/rest/middlewares/auth"
)

//SetupRoutes ...
func SetupRoutes(Echo *echo.Echo) {
	//Public routes no authentication required
	routes := Echo.Group("/users")
	routes.POST("/auth", users.Auth())

	//Private routes authentication is required
	AuthMiddleware := auth.Auth{}
	routes.Use(AuthMiddleware.Authorize)

	routes.POST("", users.Create())
	routes.GET("", users.Index())
	routes.GET("/:UUID", users.Show())
	routes.PATCH("/:UUID", users.Update())
	routes.DELETE("/:UUID", users.Destroy())
}
