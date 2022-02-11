package services

import (
	"github.com/martinsd3v/planets/core/domains/planet/repositories"
	"github.com/martinsd3v/planets/core/domains/planet/services/create"
	"github.com/martinsd3v/planets/core/domains/planet/services/destroy"
	"github.com/martinsd3v/planets/core/domains/planet/services/index"
	"github.com/martinsd3v/planets/core/domains/planet/services/show"
	"github.com/martinsd3v/planets/core/domains/planet/services/update"
	client "github.com/martinsd3v/planets/core/tools/providers/http_client"
	"github.com/martinsd3v/planets/core/tools/providers/logger"
)

type Dependences struct {
	Repository repositories.IPlanetRepository
	Logger     logger.ILoggerProvider
	HTTPClient client.IHTTPClientProvider
}

type Services struct {
	Create  create.Service
	Update  update.Service
	Index   index.Service
	Show    show.Service
	Destroy destroy.Service
}

func New(dep Dependences) *Services {
	return &Services{

		Create: create.Service{
			Repository: dep.Repository,
			Logger:     dep.Logger,
			HTTPClient: dep.HTTPClient,
		},
		Update: update.Service{
			Repository: dep.Repository,
			Logger:     dep.Logger,
			HTTPClient: dep.HTTPClient,
		},
		Index: index.Service{
			Repository: dep.Repository,
			Logger:     dep.Logger,
		},
		Show: show.Service{
			Repository: dep.Repository,
			Logger:     dep.Logger,
		},
		Destroy: destroy.Service{
			Repository: dep.Repository,
			Logger:     dep.Logger,
		},
	}
}
