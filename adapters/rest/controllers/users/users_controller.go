package users

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/martinsd3v/planets/adapters/persistence/mongodb/repositories/users"
	"github.com/martinsd3v/planets/adapters/rest/util"
	"github.com/martinsd3v/planets/core/domains/user/services"
	"github.com/martinsd3v/planets/core/domains/user/services/authenticate"
	"github.com/martinsd3v/planets/core/domains/user/services/create"
	"github.com/martinsd3v/planets/core/domains/user/services/update"
	"github.com/martinsd3v/planets/core/tools/providers/hash"
	"github.com/martinsd3v/planets/core/tools/providers/jwt"
	"github.com/martinsd3v/planets/core/tools/providers/logger"
	"github.com/spf13/viper"
)

func service(ctx context.Context) *services.Services {
	mongoRepo := users.Setup(ctx)

	return services.New(services.Dependences{
		Repository: mongoRepo,
		Logger:     logger.New(),
		Hash:       hash.New(),
		Jwt:        jwt.New(viper.GetString("jwt.secretKey")),
	})
}

//Auth ...
func Auth() echo.HandlerFunc {
	return func(c echo.Context) error {
		dto := authenticate.Dto{}

		util.Parser(c.Request(), &dto)

		service := service(c.Request().Context()).Authenticate
		token, response := service.Execute(dto)
		if token != "" {
			response.Data = token
		}

		c.JSON(response.Status, response)
		return nil
	}
}

//Create ...
func Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		dto := create.Dto{}

		util.Parser(c.Request(), &dto)

		service := service(c.Request().Context()).Create
		created, response := service.Execute(dto)
		if created.UUID != "" {
			response.Data = created.PublicUser()
		}

		c.JSON(response.Status, response)
		return nil
	}
}

//Index ...
func Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		service := service(c.Request().Context()).Index
		result, response := service.Execute()
		response.Data = result.PublicUsers()

		c.JSON(response.Status, response)
		return nil
	}
}

//Show ...
func Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		uuid := c.Param("UUID")
		service := service(c.Request().Context()).Show
		result, response := service.Execute(uuid)

		if result.UUID != "" {
			response.Data = result.PublicUser()
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

		service := service(c.Request().Context()).Update
		result, response := service.Execute(dto)
		if result.UUID != "" {
			response.Data = result.PublicUser()
		}

		c.JSON(response.Status, response)
		return nil
	}
}

//Destroy ...
func Destroy() echo.HandlerFunc {
	return func(c echo.Context) error {
		uuid := c.Param("UUID")
		service := service(c.Request().Context()).Destroy
		response := service.Execute(uuid)

		c.JSON(response.Status, response)
		return nil
	}
}
