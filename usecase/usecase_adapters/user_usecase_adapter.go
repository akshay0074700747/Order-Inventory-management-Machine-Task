package usecaseadapters

import (
	"github.com/akshay0074700747/order-inventory_management/entities"
	"github.com/akshay0074700747/order-inventory_management/helpers"
	repositoryports "github.com/akshay0074700747/order-inventory_management/repository/repository_ports"
)

// UserUsecaseAdapter implements the UserUsecasePort interface
type UserUsecaseAdapter struct {
	Repo repositoryports.UserRepositoryPort
}

func NewUserUsecaseAdapter(repo repositoryports.UserRepositoryPort) *UserUsecaseAdapter {

	return &UserUsecaseAdapter{
		Repo: repo,
	}
}

func (userUsecase *UserUsecaseAdapter) Signup(user entities.User) (entities.User, error) {

	//generating uuid
	user.UserID = helpers.GenUuid()
	//hashing password
	hashed, err := helpers.HashPass(user.Password)
	if err != nil {
		return entities.User{}, err
	}

	user.Password = hashed

	res, err := userUsecase.Repo.Signup(user)
	if err != nil {
		return entities.User{}, err
	}

	res.Password = ""

	return res, nil
}

func (userUsecase *UserUsecaseAdapter) Login(user entities.User) (entities.User, error) {

	res, err := userUsecase.Repo.Login(user)
	if err != nil {
		return entities.User{}, err
	}

	res.Password = ""

	return res, nil
}

func (userUsecase *UserUsecaseAdapter) GetMostOrderedUsers(pageNo, limit int) ([]entities.GetUser, error) {

	//finding offset and limit for pagination
	offset, limit := helpers.FindLimitandOffset(pageNo, limit)

	return userUsecase.Repo.GetMostOrderedUsers(offset, limit)
}
