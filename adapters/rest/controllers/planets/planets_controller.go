package planets

import (
	"github.com/labstack/echo/v4"
	"github.com/martinsd3v/planets/adapters/persistence/mongodb/repositories/planets"
	"github.com/martinsd3v/planets/adapters/rest/util"
	"github.com/martinsd3v/planets/core/domains/planet/services/create"
	"github.com/martinsd3v/planets/core/domains/planet/services/destroy"
	"github.com/martinsd3v/planets/core/domains/planet/services/index"
	"github.com/martinsd3v/planets/core/domains/planet/services/show"
	"github.com/martinsd3v/planets/core/domains/planet/services/update"
	client "github.com/martinsd3v/planets/core/tools/providers/http_client"
	"github.com/martinsd3v/planets/core/tools/providers/logger"
)

//Controller ...
type Controller struct{}

//Create ...
func (ctrl *Controller) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		mongoRepo := planets.Setup(c.Request().Context())
		service := create.Service{
			Repository: mongoRepo,
			Logger:     logger.New(),
			HTTPClient: client.New(),
		}
		dto := create.Dto{}

		util.Parser(c.Request(), &dto)

		created, response := service.Execute(dto)
		if created.UUID != "" {
			response.Data = created.PublicPlanet()
		}

		c.JSON(response.Status, response)
		return nil
	}
}

//Index ...
func (ctrl *Controller) Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		mongoRepo := planets.Setup(c.Request().Context())
		service := index.Service{
			Repository: mongoRepo,
			Logger:     logger.New(),
		}

		filters := map[string]interface{}{}
		name := c.QueryParam("name")
		if name != "" {
			filters["name"] = name
		}

		result, response := service.Execute(&filters)
		response.Data = result.PublicPlanets()

		c.JSON(response.Status, response)
		return nil
	}
}

//Show ...
func (ctrl *Controller) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		mongoRepo := planets.Setup(c.Request().Context())
		service := show.Service{
			Repository: mongoRepo,
			Logger:     logger.New(),
		}

		uuid := c.Param("UUID")
		result, response := service.Execute(uuid)

		if result.UUID != "" {
			response.Data = result.PublicPlanet()
		}

		c.JSON(response.Status, response)
		return nil
	}
}

//Update ...
func (ctrl *Controller) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		mongoRepo := planets.Setup(c.Request().Context())
		service := update.Service{
			Repository: mongoRepo,
			Logger:     logger.New(),
			HTTPClient: client.New(),
		}
		dto := update.Dto{}
		dto.UUID = c.Param("UUID")

		util.Parser(c.Request(), &dto)
		result, response := service.Execute(dto)
		if result.UUID != "" {
			response.Data = result.PublicPlanet()
		}

		c.JSON(response.Status, response)
		return nil
	}
}

//Destroy ...
func (ctrl *Controller) Destroy() echo.HandlerFunc {
	return func(c echo.Context) error {
		mongoRepo := planets.Setup(c.Request().Context())
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
