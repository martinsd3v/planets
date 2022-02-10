package users

import (
	"context"
	"time"

	"github.com/martinsd3v/planets/adapters/persistence/mongodb"
	"github.com/martinsd3v/planets/adapters/persistence/mongodb/util"
	"github.com/martinsd3v/planets/core/domains/user/entities"
	"github.com/martinsd3v/planets/core/domains/user/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//Repository based on user repository from core
type Repository struct {
	Context    context.Context
	Collection *mongo.Collection
}

//Sing repository in interface
var _ repositories.IUserRepository = &Repository{}

//Setup the repository
func Setup(ctx context.Context) *Repository {
	conn := mongodb.New(ctx)
	return &Repository{
		Collection: conn.MongoDB.Collection("users"),
		Context:    ctx,
	}
}

//All return all users
func (repo *Repository) All() (entities.Users, error) {
	var data entities.Users
	cursor, err := repo.Collection.Find(repo.Context, bson.M{})
	if err != nil {
		return data, err
	}

	for cursor.Next(repo.Context) {
		row := entities.User{}
		cursor.Decode(&row)
		data = append(data, row)
	}

	return data, nil
}

//Create insert user in DB
func (repo *Repository) Create(data entities.User) (entities.User, error) {
	_, err := repo.Collection.InsertOne(repo.Context, data)
	if err != nil {
		return entities.User{}, err
	}
	return repo.FindByUUID(data.UUID)
}

//FindByEmail find user by email
func (repo *Repository) FindByEmail(email string) (entities.User, error) {
	var data entities.User
	err := repo.Collection.FindOne(repo.Context, bson.M{"email": email}).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

//FindByUUID find user by uuid
func (repo *Repository) FindByUUID(uuid string) (entities.User, error) {
	var data entities.User
	err := repo.Collection.FindOne(repo.Context, bson.M{"_id": uuid}).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

//Destroy delete user by uuid
func (repo *Repository) Destroy(uuid string) error {
	_, err := repo.Collection.DeleteOne(repo.Context, bson.M{"_id": uuid})
	if err != nil {
		return err
	}
	return nil
}

//Save update user
func (repo *Repository) Save(data entities.User) (entities.User, error) {
	now := time.Now()
	data.UpdatedAt = &now

	primitiveD, err := util.ToUpdate(data)
	if err != nil {
		return entities.User{}, err
	}

	_, err = repo.Collection.UpdateOne(
		repo.Context,
		bson.M{"_id": data.UUID},
		primitiveD,
	)
	if err != nil {
		return entities.User{}, err
	}
	return repo.FindByUUID(data.UUID)
}
