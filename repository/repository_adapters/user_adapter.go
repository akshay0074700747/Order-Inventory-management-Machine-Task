package repositoryadapters

import (
	"errors"

	customerrormessages "github.com/akshay0074700747/order-inventory_management/customError_messages"
	"github.com/akshay0074700747/order-inventory_management/entities"
	"github.com/akshay0074700747/order-inventory_management/helpers"
	"gorm.io/gorm"
)

// this is the implementatioin of the userrepositoryport interface
type UserRepositoryAdapter struct {
	DB *gorm.DB
}

// constructor for userrepositoryadapter
func NewUserRepositoryAdapter(db *gorm.DB) *UserRepositoryAdapter {
	return &UserRepositoryAdapter{
		DB: db,
	}
}

func (userRepo *UserRepositoryAdapter) Signup(user entities.User) (entities.User, error) {

	var result entities.User
	//starting a transaction
	return result, userRepo.DB.Transaction(func(tx *gorm.DB) error {
		//adding the user in the table
		if err := tx.Create(&user).Scan(&result).Error; err != nil {
			return err
		}
		return nil
	})
}

func (userRepo *UserRepositoryAdapter) Login(user entities.User) (entities.User, error) {

	var result entities.User
	record := userRepo.DB.Where("email = ?", user.Email).First(&result)

	// checking wheather the user is found or not
	if record.RowsAffected == 0 {
		return entities.User{}, errors.New(customerrormessages.UserNotFoundError)
	}

	// checking wheather the hashed and password match
	if !helpers.MatchPass(result.Password, user.Password) {
		return entities.User{}, errors.New(customerrormessages.EmailorPassNotMatch)
	}

	return result, nil
}

func (userRepo *UserRepositoryAdapter) GetMostOrderedUsers(offset, limit int) ([]entities.GetUser, error) {

	query := `SELECT u.user_id, u.name, u.email, u.phone, o.order_count FROM users u 
	INNER JOIN ( SELECT user_id, COUNT(*) AS order_count 
	FROM orders GROUP BY user_id ORDER BY order_count DESC ) o 
	ON u.user_id = o.user_id ORDER BY o.order_count DESC OFFSET $1 LIMIT $2`

	var users []entities.GetUser
	if err := userRepo.DB.Raw(query, offset, limit).Scan(&users).Error; err != nil {
		return nil, err
	}

	return users, nil

}
