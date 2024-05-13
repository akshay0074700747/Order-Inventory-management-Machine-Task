package usecaseports

import "github.com/akshay0074700747/order-inventory_management/entities"

// OrderUsecasePort is the abstration interface for achieving loosely coupling between dependencies
type OrderUsecasePort interface {
	PlaceOrder(order entities.Order, items []entities.OrderItems) error
	GetOrdersSortedbyTime(offset, limit int) ([]entities.UserOrder, error)
	GetUserOrders(userID string, offset, limit int) ([]entities.UserOrder, error)
	GetProductOrders(productID string, offset, limit int) ([]entities.UserOrder, error)
	GetOrderswithUsers(offset, limit int) ([]entities.UserOrder, error)
	GetOrdersSortedByPrice(offset, limit int) ([]entities.UserOrder, error)
}
