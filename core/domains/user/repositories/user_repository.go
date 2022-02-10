package repositories

import (
	"github.com/martinsd3v/planets/core/domains/user/entities"
)

//IUserRepository ...
type IUserRepository interface {
	All() (entities.Users, error)
	Create(entities.User) (entities.User, error)
	FindByEmail(email string) (entities.User, error)
	FindByUUID(string) (entities.User, error)
	Destroy(uuid string) error
	Save(entities.User) (entities.User, error)
}
