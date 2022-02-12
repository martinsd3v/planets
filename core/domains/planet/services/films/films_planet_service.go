package films

import (
	"encoding/json"
	"time"

	"github.com/martinsd3v/planets/core/tools/providers/cache"
	client "github.com/martinsd3v/planets/core/tools/providers/http_client"
	"github.com/martinsd3v/planets/core/tools/providers/logger"
)

//Service ...
type Service struct {
	Logger     logger.ILoggerProvider
	HTTPClient client.IHTTPClientProvider
	Cache      cache.ICacheProvider
}

//ResponseAPI ...
type ResponseAPI struct {
	Results []Result `json:"results"`
}

//Result ...
type Result struct {
	Films []string `json:"films"`
}

//Execute service
func (service *Service) Execute(planetName string) int {
	var quantityFilms int
	if service.Cache.Get(planetName, &quantityFilms) == nil {
		return quantityFilms
	}

	apiBase := "https://swapi.dev/api/planets/?search="
	response, err := service.HTTPClient.Get(apiBase + planetName)
	if err != nil {
		service.Logger.Error("domain.movies.service.movies_planet_service.HTTPClient.Get", err)
		return 0
	}

	var films ResponseAPI
	err = json.NewDecoder(response.Body).Decode(&films)
	if err != nil {
		service.Logger.Error("domain.movies.service.movies_planet_service.json.NewDecoder.Decode", err)
		return 0
	}

	if len(films.Results) > 0 {
		for _, film := range films.Results {
			quantityFilms += len(film.Films)
		}

		cacheExpireTime := time.Minute * 20
		service.Cache.WithExpiration(cacheExpireTime).Set(planetName, quantityFilms)
	}

	return quantityFilms
}
