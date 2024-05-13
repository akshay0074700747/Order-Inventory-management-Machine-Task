package db

import (
	"fmt"
	"log"

	"github.com/akshay0074700747/order-inventory_management/config"
	"github.com/akshay0074700747/order-inventory_management/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg config.Configurations) *gorm.DB {

	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBhost, cfg.DBuser, cfg.DBname, cfg.DBport, cfg.DBpassword)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatalf("cannot connect to the db: %s", err.Error())
	}
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Product{})
	db.AutoMigrate(&entities.Order{})
	db.AutoMigrate(&entities.OrderItems{})
	return db
}
