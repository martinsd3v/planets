package planets

import (
	"context"
	"time"

	"github.com/martinsd3v/planets/adapters/persistence/mongodb"
	"github.com/martinsd3v/planets/adapters/persistence/mongodb/util"
	"github.com/martinsd3v/planets/core/domains/planet/entities"
	"github.com/martinsd3v/planets/core/domains/planet/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//Repository based on planet repository from core
type Repository struct {
	Context    context.Context
	Collection *mongo.Collection
}

//Sing repository in interface
var _ repositories.IPlanetRepository = &Repository{}

//Setup ...
func Setup(ctx context.Context) *Repository {
	conn := mongodb.New(ctx)
	return &Repository{
		Collection: conn.MongoDB.Collection("planets"),
		Context:    ctx,
	}
}

//All return all planets
func (repo *Repository) All(filters *map[string]interface{}) (entities.Planets, error) {
	var data entities.Planets

	cursor, err := repo.Collection.Find(repo.Context, filters)
	if err != nil {
		return data, err
	}

	for cursor.Next(repo.Context) {
		row := entities.Planet{}
		cursor.Decode(&row)
		data = append(data, row)
	}

	return data, nil
}

//Create insert planet in DB
func (repo *Repository) Create(data entities.Planet) (entities.Planet, error) {
	_, err := repo.Collection.InsertOne(repo.Context, data)
	if err != nil {
		return entities.Planet{}, err
	}
	return repo.FindByUUID(data.UUID)
}

//FindByUUID find planet by uuid
func (repo *Repository) FindByUUID(uuid string) (entities.Planet, error) {
	var data entities.Planet
	err := repo.Collection.FindOne(repo.Context, bson.M{"_id": uuid}).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

//Destroy delete planet by uuid
func (repo *Repository) Destroy(uuid string) error {
	_, err := repo.Collection.DeleteOne(repo.Context, bson.M{"_id": uuid})
	if err != nil {
		return err
	}
	return nil
}

//Save update planet
func (repo *Repository) Save(data entities.Planet) (entities.Planet, error) {
	now := time.Now()
	data.UpdatedAt = &now

	primitiveD, err := util.ToUpdate(data)
	if err != nil {
		return entities.Planet{}, err
	}

	_, err = repo.Collection.UpdateOne(
		repo.Context,
		bson.M{"_id": data.UUID},
		primitiveD,
	)
	if err != nil {
		return entities.Planet{}, err
	}
	return repo.FindByUUID(data.UUID)
}
