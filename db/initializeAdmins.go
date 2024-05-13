package db

import (
	"fmt"

	"github.com/akshay0074700747/order-inventory_management/entities"
	"github.com/akshay0074700747/order-inventory_management/helpers"
	"gorm.io/gorm"
)

// initializes the admins in the database if not already exists
func InitializeAdmin(email, password string, db *gorm.DB) error {

	var count int64
	res := db.Model(&entities.User{}).Count(&count)

	if res.Error != nil {
		return res.Error
	}

	if count > 0 {
		return nil
	}

	userID := helpers.GenUuid()
	hashedPass, err := helpers.HashPass(password)
	if err != nil {
		return err
	}

	if err := db.Create(&entities.User{
		UserID:   userID,
		Email:    email,
		Password: hashedPass,
		IsAdmin:  true,
	}).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
