package show

import (
	"github.com/martinsd3v/planets/core/domains/planet/entities"
	"github.com/martinsd3v/planets/core/domains/planet/repositories"
	"github.com/martinsd3v/planets/core/tools/communication"
	"github.com/martinsd3v/planets/core/tools/providers/logger"
)

//Service ...
type Service struct {
	Repository repositories.IPlanetRepository
	Logger     logger.ILoggerProvider
}

//Execute service responsible for find one register
func (service *Service) Execute(planetUUID string) (planet entities.Planet, response communication.Response) {
	planet, err := service.Repository.FindByUUID(planetUUID)
	comm := communication.New()

	if err != nil {
		service.Logger.Error("domain.planet.service.show.show_planet_service.Repository.FindByUUID", err)
		response = comm.Response(404, "error_list")
		return
	}

	response = comm.Response(200, "success")
	return
}
