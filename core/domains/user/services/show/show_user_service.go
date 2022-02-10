package show

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
func (service *Service) Execute(userUUID string) (user entities.User, response communication.Response) {
	user, err := service.Repository.FindByUUID(userUUID)
	comm := communication.New()

	if err != nil {
		service.Logger.Error("domain.user.service.show.show_user_service.Repository.FindByUUID", err)
		response = comm.Response(404, "error_list")
		return
	}

	response = comm.Response(200, "success")
	return
}
