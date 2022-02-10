package users

import (
	"github.com/labstack/echo/v4"
	"github.com/martinsd3v/planets/adapters/persistence/mongodb/repositories/users"
	"github.com/martinsd3v/planets/adapters/rest/util"
	"github.com/martinsd3v/planets/core/domains/user/services/authenticate"
	"github.com/martinsd3v/planets/core/domains/user/services/create"
	"github.com/martinsd3v/planets/core/domains/user/services/destroy"
	"github.com/martinsd3v/planets/core/domains/user/services/index"
	"github.com/martinsd3v/planets/core/domains/user/services/show"
	"github.com/martinsd3v/planets/core/domains/user/services/update"
	"github.com/martinsd3v/planets/core/tools/providers/hash"
	"github.com/martinsd3v/planets/core/tools/providers/jwt"
	"github.com/martinsd3v/planets/core/tools/providers/logger"
	"github.com/spf13/viper"
)

//Controller ...
type Controller struct{}

//Auth ...
func (ctrl *Controller) Auth() echo.HandlerFunc {
	return func(c echo.Context) error {
		mongoRepo := users.Setup(c.Request().Context())
		service := authenticate.Service{
			Repository: mongoRepo,
			Hash:       hash.New(),
			Logger:     logger.New(),
			Jwt:        jwt.New(viper.GetString("jwt.secretKey")),
		}
		dto := authenticate.Dto{}

		util.Parser(c.Request(), &dto)

		token, response := service.Execute(dto)
		if token != "" {
			response.Data = token
		}

		c.JSON(response.Status, response)
		return nil
	}
}

//Create ...
func (ctrl *Controller) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		mongoRepo := users.Setup(c.Request().Context())
		service := create.Service{
			Repository: mongoRepo,
			Hash:       hash.New(),
			Logger:     logger.New(),
		}
		dto := create.Dto{}

		util.Parser(c.Request(), &dto)

		created, response := service.Execute(dto)
		if created.UUID != "" {
			response.Data = created.PublicUser()
		}

		c.JSON(response.Status, response)
		return nil
	}
}

//Index ...
func (ctrl *Controller) Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		mongoRepo := users.Setup(c.Request().Context())
		service := index.Service{
			Repository: mongoRepo,
			Logger:     logger.New(),
		}

		result, response := service.Execute()
		response.Data = result.PublicUsers()

		c.JSON(response.Status, response)
		return nil
	}
}

//Show ...
func (ctrl *Controller) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		mongoRepo := users.Setup(c.Request().Context())
		service := show.Service{
			Repository: mongoRepo,
			Logger:     logger.New(),
		}

		uuid := c.Param("UUID")
		result, response := service.Execute(uuid)

		if result.UUID != "" {
			response.Data = result.PublicUser()
		}

		c.JSON(response.Status, response)
		return nil
	}
}

//Update ...
func (ctrl *Controller) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		mongoRepo := users.Setup(c.Request().Context())
		service := update.Service{
			Repository: mongoRepo,
			Hash:       hash.New(),
			Logger:     logger.New(),
		}
		dto := update.Dto{}
		dto.UUID = c.Param("UUID")

		util.Parser(c.Request(), &dto)
		result, response := service.Execute(dto)
		if result.UUID != "" {
			response.Data = result.PublicUser()
		}

		c.JSON(response.Status, response)
		return nil
	}
}

//Destroy ...
func (ctrl *Controller) Destroy() echo.HandlerFunc {
	return func(c echo.Context) error {
		mongoRepo := users.Setup(c.Request().Context())
		service := destroy.Service{
			Repository: mongoRepo,
			Logger:     logger.New(),
		}

		uuid := c.Param("UUID")
		response := service.Execute(uuid)

		c.JSON(response.Status, response)
		return nil
	}
}
