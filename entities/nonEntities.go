package entities

import "time"

// this struct is to view order details
type UserOrder struct {
	OrderID   string       `json:"OrderID,omitempty"`
	UserID    string       `json:"UserID,omitempty"`
	Name      string       `json:"Name,omitempty"`
	Email     string       `json:"Email,omitempty"`
	Phone     string       `json:"Phone,omitempty"`
	OrderDate time.Time    `json:"OrderDate,omitempty"`
	Items     []OrderItems `json:"Items,omitempty"`
}

// this struct is to view user details
type GetUser struct {
	UserID     string `gorm:"primaryKey" json:"UserID,omitempty"`
	Name       string `json:"Name,omitempty"`
	Email      string `json:"Email,omitempty"`
	Phone      string `json:"Phone,omitempty"`
	OrderCount int    `json:"OrderCount,omitempty"`
}

// this struct is to view products with orderCounts
type GetProductwithOrderCount struct {
	ProductID   string `json:"ProductID,omitempty"`
	ProductName string `json:"ProductName,omitempty"`
	Description string `json:"Description,omitempty"`
	OrderCount  int    `json:"OrderCount,omitempty"`
}

//this struct is used to communicate between the user
type Response struct {
	StatusCode int         `json:"stastuscode,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Errors     interface{} `json:"error,omitempty"`
}