package create

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

//Execute Serviço responsável pela inserção de registros
func (service *Service) Execute(dto Dto) (created entities.Planet, response communication.Response) {
	response.Fields = validations.ValidateStruct(&dto, "")
	comm := communication.New()

	filter := &map[string]interface{}{"name": dto.Name}
	planets, err := service.Repository.All(filter)
	if err != nil {
		service.Logger.Info("domain.planet.service.create.create_planet_service.Repository.All", err)
	}

	//Check planet already exists
	if len(planets) > 0 && planets[0].UUID != "" {
		response.Fields = append(response.Fields, comm.Fields("name", "already_exists"))
	}

	if len(response.Fields) > 0 {
		service.Logger.Info("domain.planet.service.create.create_planet_service.ValidationError")
		resp := comm.Response(400, "validate_failed")
		resp.Fields = response.Fields
		response = resp
		return
	}

	planetEntity := entities.Planet{}
	planet := planetEntity.New()

	//Mergin entity and dto
	toMerge, _ := json.Marshal(dto)
	json.Unmarshal(toMerge, &planet)

	filmsService := films.Service{
		Logger:     service.Logger,
		HTTPClient: service.HTTPClient,
	}
	planet.Films = filmsService.Execute(planet.Name)

	created, err = service.Repository.Create(*planet)

	if err != nil {
		service.Logger.Error("domain.planet.service.create.create_planet_service.Repository.Create", err)
		response = comm.Response(500, "error_create")
		return
	}

	response = comm.Response(200, "success")
	return
}
