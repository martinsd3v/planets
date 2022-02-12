package services

import (
	"github.com/martinsd3v/planets/core/domains/planet/repositories"
	"github.com/martinsd3v/planets/core/domains/planet/services/create"
	"github.com/martinsd3v/planets/core/domains/planet/services/destroy"
	"github.com/martinsd3v/planets/core/domains/planet/services/index"
	"github.com/martinsd3v/planets/core/domains/planet/services/show"
	"github.com/martinsd3v/planets/core/domains/planet/services/update"
	"github.com/martinsd3v/planets/core/tools/providers/cache"
	client "github.com/martinsd3v/planets/core/tools/providers/http_client"
	"github.com/martinsd3v/planets/core/tools/providers/logger"
)

type Dependences struct {
	Repository repositories.IPlanetRepository
	Logger     logger.ILoggerProvider
	HTTPClient client.IHTTPClientProvider
	Cache      cache.ICacheProvider
}

type Services struct {
	Create  create.Service
	Update  update.Service
	Index   index.Service
	Show    show.Service
	Destroy destroy.Service
	Cache   cache.ICacheProvider
}

func New(dep Dependences) *Services {
	return &Services{
		Create: create.Service{
			Repository: dep.Repository,
			Logger:     dep.Logger,
			HTTPClient: dep.HTTPClient,
			Cache:      dep.Cache,
		},
		Update: update.Service{
			Repository: dep.Repository,
			Logger:     dep.Logger,
			HTTPClient: dep.HTTPClient,
			Cache:      dep.Cache,
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
