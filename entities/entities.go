package entities

import "time"

type User struct {
	UserID   string `gorm:"primaryKey" json:"UserID,omitempty"`
	Name     string `json:"Name,omitempty"`
	Email    string `gorm:"unique;not null" json:"Email,omitempty"`
	Phone    string `gorm:"unique;not null" json:"Phone,omitempty"`
	Password string `json:"Password,omitempty"`
	IsAdmin  bool   `gorm:"default:false" json:"IsAdmin,omitempty"`
}

type Product struct {
	ProductID   string    `gorm:"primaryKey" json:"ProductID,omitempty"`
	ProductName string    `json:"ProductName,omitempty"`
	Description string    `json:"Description,omitempty"`
	Price       uint      `json:"Price,omitempty"`
	Stock       uint      `json:"Stock,omitempty"`
	UpdatedAt   time.Time `gorm:"default:NOW()" json:"UpdatedAt,omitempty"`
}

type Order struct {
	OrderID   string    `gorm:"primaryKey" json:"OrderID,omitempty"`
	UserID    string    `gorm:"foreignKey:UserID;references:users(user_id)" json:"UserID,omitempty"`
	OrderDate time.Time `gorm:"default:NOW()" json:"OrderDate,omitempty"`
	Price     int       `json:"Price,omitempty"`
}

type OrderItems struct {
	ID        uint   `gorm:"primaryKey" json:"ID,omitempty"`
	OrderID   string `gorm:"foreignKey:OrderID;references:orders(order_id)" json:"OrderID,omitempty"`
	ProductID string `gorm:"foreignKey:ProductID;references:products(product_id)" json:"ProductID,omitempty"`
	Quantity  int    `json:"Quantity,omitempty"`
}
