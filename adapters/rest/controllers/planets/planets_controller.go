package planets

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/martinsd3v/planets/adapters/persistence/mongodb/repositories/planets"
	"github.com/martinsd3v/planets/adapters/rest/util"
	"github.com/martinsd3v/planets/core/domains/planet/services"
	"github.com/martinsd3v/planets/core/domains/planet/services/create"
	"github.com/martinsd3v/planets/core/domains/planet/services/update"
	"github.com/martinsd3v/planets/core/tools/providers/cache"
	client "github.com/martinsd3v/planets/core/tools/providers/http_client"
	"github.com/martinsd3v/planets/core/tools/providers/logger"
	"github.com/spf13/viper"
)

func service(ctx context.Context) *services.Services {
	mongoRepo := planets.Setup(ctx)
	memcacheHost := viper.GetString("memcache.host") + ":" + viper.GetString("memcache.port")
	memcache, _ := cache.New(memcacheHost)

	return services.New(services.Dependences{
		Repository: mongoRepo,
		Logger:     logger.New(),
		HTTPClient: client.New(),
		Cache:      memcache,
	})
}

//Create ...
func Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		dto := create.Dto{}
		util.Parser(c.Request(), &dto)

		ctx := c.Request().Context()
		service := service(ctx).Create
		created, response := service.Execute(ctx, dto)
		if created.UUID != "" {
			response.Data = created.PublicPlanet()
		}

		c.JSON(response.Status, response)
		return nil
	}
}

//Index ...
func Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		filters := map[string]interface{}{}
		name := c.QueryParam("name")
		if name != "" {
			filters["name"] = name
		}

		ctx := c.Request().Context()
		service := service(ctx).Index
		result, response := service.Execute(ctx, &filters)
		response.Data = result.PublicPlanets()

		c.JSON(response.Status, response)
		return nil
	}
}

//Show ...
func Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		uuid := c.Param("UUID")

		ctx := c.Request().Context()
		service := service(ctx).Show
		result, response := service.Execute(ctx, uuid)

		if result.UUID != "" {
			response.Data = result.PublicPlanet()
		}

		c.JSON(response.Status, response)
		return nil
	}
}

//Update ...
func Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		dto := update.Dto{}
		dto.UUID = c.Param("UUID")
		util.Parser(c.Request(), &dto)

		ctx := c.Request().Context()
		service := service(ctx).Update
		result, response := service.Execute(ctx, dto)
		if result.UUID != "" {
			response.Data = result.PublicPlanet()
		}

		c.JSON(response.Status, response)
		return nil
	}
}

//Destroy ...
func Destroy() echo.HandlerFunc {
	return func(c echo.Context) error {
		uuid := c.Param("UUID")

		ctx := c.Request().Context()
		service := service(ctx).Destroy
		response := service.Execute(ctx, uuid)

		c.JSON(response.Status, response)
		return nil
	}
}
