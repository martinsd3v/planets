package users

import (
	"github.com/labstack/echo/v4"
	"github.com/martinsd3v/planets/adapters/rest/controllers/users"
	"github.com/martinsd3v/planets/adapters/rest/middlewares/auth"
)

//SetupRoutes ...
func SetupRoutes(Echo *echo.Echo) {
	controller := users.Controller{}

	//Public routes no authentication required
	routes := Echo.Group("/users")
	routes.POST("/auth", controller.Auth())

	//Private routes authentication is required
	AuthMiddleware := auth.Auth{}
	routes.Use(AuthMiddleware.Authorize)

	routes.POST("", controller.Create())
	routes.GET("", controller.Index())
	routes.GET("/:UUID", controller.Show())
	routes.PATCH("/:UUID", controller.Update())
	routes.DELETE("/:UUID", controller.Destroy())
}
