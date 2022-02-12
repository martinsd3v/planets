package update

import (
	"encoding/json"

	"github.com/martinsd3v/planets/core/domains/planet/entities"
	"github.com/martinsd3v/planets/core/domains/planet/repositories"
	"github.com/martinsd3v/planets/core/domains/planet/services/films"
	"github.com/martinsd3v/planets/core/tools/communication"
	client "github.com/martinsd3v/planets/core/tools/providers/http_client"
	"github.com/martinsd3v/planets/core/tools/providers/logger"
	"github.com/martinsd3v/planets/core/tools/validations"
)

//Dto object receiver
type Dto struct {
	UUID    string `json:"uuid" validate:"isRequired"`
	Name    string `json:"name" validate:"isRequired"`
	Terrain string `json:"terrain" validate:"isRequired"`
	Climate string `json:"climate" validate:"isRequired"`
}

//Service ...
type Service struct {
	Repository repositories.IPlanetRepository
	Logger     logger.ILoggerProvider
	HTTPClient client.IHTTPClientProvider
}

//Execute responsável por atualizar registros
func (service *Service) Execute(dto Dto) (updated entities.Planet, response communication.Response) {
	response.Fields = validations.ValidateStruct(&dto, "")
	comm := communication.New()

	//Check exists planet with this identifier
	planet, err := service.Repository.FindByUUID(dto.UUID)
	if err != nil {
		service.Logger.Error("domain.planet.service.update.update_planet_service.Repository.FindByUUID", err)
		response = comm.Response(500, "error_update")
		return
	}

	filter := &map[string]interface{}{"name": dto.Name}
	planets, err := service.Repository.All(filter)
	if err != nil {
		service.Logger.Info("domain.planet.service.update.update_planet_service.Repository.All", err)
	}

	//Check planet already exists
	if len(planets) > 0 && planets[0].UUID != planet.UUID {
		response.Fields = append(response.Fields, comm.Fields("name", "already_exists"))
	}

	if planet.UUID == "" {
		response.Fields = append(response.Fields, comm.Fields("uuid", "validate_invalid"))
	}

	if len(response.Fields) > 0 {
		service.Logger.Info("domain.planet.service.update.update_planet_service.ValidationError")
		resp := comm.Response(400, "validate_failed")
		resp.Fields = response.Fields
		response = resp
		return
	}

	//Mergin entity and DTO
	toMerge, _ := json.Marshal(dto)
	json.Unmarshal(toMerge, &planet)

	filmsService := films.Service{
		Logger:     service.Logger,
		HTTPClient: service.HTTPClient,
	}
	planet.Films = filmsService.Execute(planet.Name)

	updated, err = service.Repository.Save(planet)

	if err != nil {
		service.Logger.Error("domain.planet.service.update.update_planet_service.Repository.Save", err)
		response = comm.Response(500, "error_update")
		return
	}

	response = comm.Response(200, "success")
	return
}
