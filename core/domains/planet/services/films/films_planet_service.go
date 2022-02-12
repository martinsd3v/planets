package films

import (
	"encoding/json"

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
	var quantityCahe int
	if service.Cache.Get(planetName, &quantityCahe) == nil {
		return quantityCahe
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
		quantity := 0
		for _, film := range films.Results {
			quantity += len(film.Films)
		}

		service.Cache.Set(planetName, quantity)

		return quantity
	}

	return 0
}
