package repositories

import (
	"github.com/martinsd3v/planets/core/domains/planet/entities"
)

//IPlanetRepository ...
type IPlanetRepository interface {
	All(filters *map[string]interface{}) (entities.Planets, error)
	Create(entities.Planet) (entities.Planet, error)
	FindByUUID(string) (entities.Planet, error)
	Destroy(uuid string) error
	Save(entities.Planet) (entities.Planet, error)
}
