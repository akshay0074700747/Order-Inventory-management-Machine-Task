package repositoryports

import "github.com/akshay0074700747/order-inventory_management/entities"

// this userrepositoryport interface is acting as an abstaction between userusecaseAdapter and userrepositoryadapter
type UserRepositoryPort interface {
	Signup(user entities.User) (entities.User, error)
	Login(user entities.User) (entities.User, error)
	GetMostOrderedUsers(offset, limit int) ([]entities.GetUser, error)
}
