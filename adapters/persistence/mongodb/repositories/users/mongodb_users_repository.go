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
	Collection *mongo.Collection
}

//Sing repository in interface
var _ repositories.IUserRepository = &Repository{}

//Setup the repository
func Setup(ctx context.Context) *Repository {
	conn := mongodb.New(ctx)
	return &Repository{
		Collection: conn.MongoDB.Collection("users"),
	}
}

//All return all users
func (repo *Repository) All(ctx context.Context) (entities.Users, error) {
	var data entities.Users
	cursor, err := repo.Collection.Find(ctx, bson.M{})
	if err != nil {
		return data, err
	}

	for cursor.Next(ctx) {
		row := entities.User{}
		cursor.Decode(&row)
		data = append(data, row)
	}

	return data, nil
}

//Create insert user in DB
func (repo *Repository) Create(ctx context.Context, data entities.User) (entities.User, error) {
	_, err := repo.Collection.InsertOne(ctx, data)
	if err != nil {
		return entities.User{}, err
	}
	return repo.FindByUUID(ctx, data.UUID)
}

//FindByEmail find user by email
func (repo *Repository) FindByEmail(ctx context.Context, email string) (entities.User, error) {
	var data entities.User
	err := repo.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

//FindByUUID find user by uuid
func (repo *Repository) FindByUUID(ctx context.Context, uuid string) (entities.User, error) {
	var data entities.User
	err := repo.Collection.FindOne(ctx, bson.M{"_id": uuid}).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

//Destroy delete user by uuid
func (repo *Repository) Destroy(ctx context.Context, uuid string) error {
	_, err := repo.Collection.DeleteOne(ctx, bson.M{"_id": uuid})
	if err != nil {
		return err
	}
	return nil
}

//Save update user
func (repo *Repository) Save(ctx context.Context, data entities.User) (entities.User, error) {
	now := time.Now()
	data.UpdatedAt = &now

	primitiveD, err := util.ToUpdate(data)
	if err != nil {
		return entities.User{}, err
	}

	_, err = repo.Collection.UpdateOne(
		ctx,
		bson.M{"_id": data.UUID},
		primitiveD,
	)
	if err != nil {
		return entities.User{}, err
	}
	return repo.FindByUUID(ctx, data.UUID)
}
