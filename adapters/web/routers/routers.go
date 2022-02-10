package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/martinsd3v/planets/adapters/web/routers/planets"
	"github.com/martinsd3v/planets/adapters/web/routers/users"
)

//StartRouters ...
func StartRouters(Echo *echo.Echo) {
	users.SetupRoutes(Echo)
	planets.SetupRoutes(Echo)
}
