package usecaseports

import "github.com/akshay0074700747/order-inventory_management/entities"

//UserUsecasePort is the abstration interface for achieving loosely coupling between dependencies
type UserUsecasePort interface{
	Signup(user entities.User) (entities.User, error)
	Login(user entities.User) (entities.User, error)
	GetMostOrderedUsers(pageNo, limit int) ([]entities.GetUser, error)
}