package usecaseadapters

import (
	"errors"

	"github.com/akshay0074700747/order-inventory_management/entities"
	"github.com/akshay0074700747/order-inventory_management/helpers"
	repositoryports "github.com/akshay0074700747/order-inventory_management/repository/repository_ports"
)

// OrderUsecaseAdapater implements the OrderUsecasePort interface
type OrderUsecaseAdpater struct {
	Repo repositoryports.OrderRepositoryPort
}

func NewOrderUsecaseAdpater(repo repositoryports.OrderRepositoryPort) *OrderUsecaseAdpater {

	return &OrderUsecaseAdpater{
		Repo: repo,
	}
}

func (orderUsecase *OrderUsecaseAdpater) PlaceOrder(order entities.Order, items []entities.OrderItems) error {

	//generating the orderID
	order.OrderID = helpers.GenUuid()

	return orderUsecase.Repo.PlaceOrder(order, items)
}

func (orderUsecase *OrderUsecaseAdpater) GetOrdersSortedbyTime(pageNo, limit int) ([]entities.UserOrder, error) {

	offset, limit := helpers.FindLimitandOffset(pageNo, limit)

	return orderUsecase.Repo.GetOrdersSortedbyTime(offset, limit)
}

func (orderUsecase *OrderUsecaseAdpater) GetUserOrders(userID string, pageNo, limit int) ([]entities.UserOrder, error) {

	//checking wheather the userID is valid
	if userID == "" {
		return nil, errors.New("the userID cannot be empty")
	}

	offset, limit := helpers.FindLimitandOffset(pageNo, limit)

	return orderUsecase.Repo.GetUserOrders(userID, offset, limit)
}

func (orderUsecase *OrderUsecaseAdpater) GetProductOrders(productID string, pageNo, limit int) ([]entities.UserOrder, error) {

	//checking wheather the productID is valid
	if productID == "" {
		return nil, errors.New("the productID cannot be empty")
	}

	offset, limit := helpers.FindLimitandOffset(pageNo, limit)

	return orderUsecase.Repo.GetProductOrders(productID, offset, limit)
}

func (orderUsecase *OrderUsecaseAdpater) GetOrderswithUsers(pageNo, limit int) ([]entities.UserOrder, error) {

	offset, limit := helpers.FindLimitandOffset(pageNo, limit)

	return orderUsecase.Repo.GetOrderswithUsers(offset, limit)
}

func (orderUsecase *OrderUsecaseAdpater) GetOrdersSortedByPrice(pageNo, limit int) ([]entities.UserOrder, error) {

	offset, limit := helpers.FindLimitandOffset(pageNo, limit)

	return orderUsecase.Repo.GetOrdersSortedByPrice(offset, limit)
}
