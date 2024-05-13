package repositoryports

import "github.com/akshay0074700747/order-inventory_management/entities"

// this orderrepositoryport interface is acting as an abstaction between orderusecaseAdapter and orderrepositoryadapter
type OrderRepositoryPort interface {
	PlaceOrder(order entities.Order, items []entities.OrderItems) error
	GetOrdersSortedbyTime(offset, limit int) ([]entities.UserOrder, error)
	GetUserOrders(userID string, offset, limit int) ([]entities.UserOrder, error)
	GetProductOrders(productID string, offset, limit int) ([]entities.UserOrder, error)
	GetOrderswithUsers(offset, limit int) ([]entities.UserOrder, error)
	GetOrdersSortedByPrice(offset, limit int) ([]entities.UserOrder, error)
}
