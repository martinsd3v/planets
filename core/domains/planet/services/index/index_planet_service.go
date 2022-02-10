package index

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
func (service *Service) Execute(filter *map[string]interface{}) (planets entities.Planets, response communication.Response) {
	planets, err := service.Repository.All(filter)
	comm := communication.New()

	if err != nil {
		service.Logger.Error("domain.planet.service.index.index_planet_service.Repository.All", err)
		response = comm.Response(404, "error_list")
		return
	}

	response = comm.Response(200, "success")
	return
}
