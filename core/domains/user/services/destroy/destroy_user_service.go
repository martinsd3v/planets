package destroy

import (
	"github.com/martinsd3v/planets/core/domains/user/repositories"
	"github.com/martinsd3v/planets/core/tools/communication"
	"github.com/martinsd3v/planets/core/tools/providers/logger"
)

//Service ...
type Service struct {
	Repository repositories.IUserRepository
	Logger     logger.ILoggerProvider
}

//Execute responsible for deleting register
func (service *Service) Execute(userUUID string) (response communication.Response) {
	comm := communication.New()

	if userUUID == "" {
		response = comm.Response(400, "validate_failed")
		response.Fields = append(response.Fields, comm.Fields("uuid", "validate_required"))
		service.Logger.Info("domain.user.service.destroy.destroy_user_service.Validation")
		return
	}

	user, err := service.Repository.FindByUUID(userUUID)
	if err != nil {
		service.Logger.Error("domain.user.service.destroy.destroy_user_service.Repository.FindByUUID", err)
		response = comm.Response(500, "error_delete")
		return
	}

	if user.UUID != "" {
		err = service.Repository.Destroy(userUUID)
		if err != nil {
			service.Logger.Error("domain.user.service.destroy.destroy_user_service.Repository.Destroy", err)
			response = comm.Response(500, "error_delete")
			return
		}

		response = comm.Response(200, "success")
		return
	}

	service.Logger.Error("domain.user.service.destroy.destroy_user_service.Repository", err)
	response = comm.Response(500, "error_delete")
	return
}
