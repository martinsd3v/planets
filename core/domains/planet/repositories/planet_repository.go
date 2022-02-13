package repositories

import (
	"context"

	"github.com/martinsd3v/planets/core/domains/planet/entities"
)

//IPlanetRepository ...
type IPlanetRepository interface {
	All(ctx context.Context, filters *map[string]interface{}) (entities.Planets, error)
	Create(ctx context.Context, data entities.Planet) (entities.Planet, error)
	FindByUUID(ctx context.Context, uuid string) (entities.Planet, error)
	Destroy(ctx context.Context, uuid string) error
	Save(ctx context.Context, data entities.Planet) (entities.Planet, error)
}
