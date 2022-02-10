package index

import (
	"github.com/martinsd3v/planets/core/domains/user/entities"
	"github.com/martinsd3v/planets/core/domains/user/repositories"
	"github.com/martinsd3v/planets/core/tools/communication"
	"github.com/martinsd3v/planets/core/tools/providers/logger"
)

//Service ...
type Service struct {
	Repository repositories.IUserRepository
	Logger     logger.ILoggerProvider
}

//Execute service responsible for find one register
func (service *Service) Execute() (users entities.Users, response communication.Response) {
	users, err := service.Repository.All()
	comm := communication.New()

	if err != nil {
		service.Logger.Error("domain.user.service.index.index_user_service.Repository.All", err)
		response = comm.Response(404, "error_list")
		return
	}

	response = comm.Response(200, "success")
	return
}
