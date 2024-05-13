package repositoryadapters

import (
	"errors"

	customerrormessages "github.com/akshay0074700747/order-inventory_management/customError_messages"
	"github.com/akshay0074700747/order-inventory_management/entities"
	"gorm.io/gorm"
)

// this is the implementatioin of the orderrepositoryport interface
type OrderRepositoryAdapter struct {
	DB *gorm.DB
}

func NewOrderRepositoryAdapter(db *gorm.DB) *OrderRepositoryAdapter {
	return &OrderRepositoryAdapter{
		DB: db,
	}
}

func (orderRepo *OrderRepositoryAdapter) PlaceOrder(order entities.Order, items []entities.OrderItems) error {

	return orderRepo.DB.Transaction(func(tx *gorm.DB) error {
		// checking and updating the product stock for each item
		for i, item := range items {
			// checking product stock
			var product entities.Product
			result := tx.Where("product_id = ?", item.ProductID).First(&product)
			if result.RowsAffected == 0 {
				return errors.New(customerrormessages.ProductNotFoundError + item.ProductID)
			}

			if product.Stock < uint(item.Quantity) {
				return errors.New(customerrormessages.InsufficientStock + item.ProductID)
			}

			// newStock := product.Stock - uint(item.Quantity)
			// if err := tx.Model(&product).Update("stock", newStock).Error; err != nil {
			// 	return errors.New("failed to update product stock: " + item.ProductID)
			// }

			items[i].OrderID = order.OrderID
			order.Price += int(product.Price) * item.Quantity

			if err := tx.Create(&item).Error; err != nil {
				return err
			}

		}

		//creating the order
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		return nil
	})
}

func (orderRepo *OrderRepositoryAdapter) GetUserOrders(userID string, offset, limit int) ([]entities.UserOrder, error) {

	//uses raw query option to execute raw query for more control over the database
	query := `SELECT o.order_id, o.user_id, o.order_date, oi.id AS item_id,
	 oi.product_id, oi.quantity FROM orders o INNER JOIN order_items oi ON oi.order_id = o.order_id 
	 WHERE o.user_id = $1 ORDER BY o.order_date DESC OFFSET $2 LIMIT $3`

	var userOrder []entities.UserOrder
	if err := orderRepo.DB.Raw(query, userID, offset, limit).Scan(&userOrder).Error; err != nil {
		return nil, err
	}

	return userOrder, nil
}

func (orderRepo *OrderRepositoryAdapter) GetProductOrders(productID string, offset, limit int) ([]entities.UserOrder, error) {

	query := `SELECT o.order_id, o.user_id, o.order_date, oi.id AS item_id,
	 oi.product_id, oi.quantity FROM orders o 
	 INNER JOIN order_items oi ON oi.order_id = o.order_id AND oi.product_id = $1 
	 ORDER BY o.order_date DESC OFFSET $2 LIMIT $3`

	var userOrder []entities.UserOrder
	if err := orderRepo.DB.Raw(query, productID, offset, limit).Scan(&userOrder).Error; err != nil {
		return nil, err
	}

	return userOrder, nil
}

func (orderRepo *OrderRepositoryAdapter) GetOrders(offset, limit int) ([]entities.UserOrder, error) {

	query := `SELECT o.order_id, o.user_id, o.order_date, oi.id AS item_id,
	 oi.product_id, oi.quantity FROM orders o 
	 INNER JOIN order_items oi ON oi.order_id = o.order_id ORDER BY o.order_date DESC OFFSET $1 LIMIT $2`

	var userOrder []entities.UserOrder
	if err := orderRepo.DB.Raw(query, offset, limit).Scan(&userOrder).Error; err != nil {
		return nil, err
	}

	return userOrder, nil
}

func (orderRepo *OrderRepositoryAdapter) GetOrderswithUsers(offset, limit int) ([]entities.UserOrder, error) {

	query := `SELECT o.order_id, o.user_id, u.name, u.email, u.phone, o.order_date, oi.id AS item_id,
	 oi.product_id, oi.quantity FROM orders o JOIN order_items oi ON oi.order_id = o.order_id 
	 JOIN users u ON o.user_id = u.user_id ORDER BY o.order_date DESC OFFSET $1 LIMIT $2`

	var userOrder []entities.UserOrder
	if err := orderRepo.DB.Raw(query, offset, limit).Scan(&userOrder).Error; err != nil {
		return nil, err
	}

	return userOrder, nil
}

func (orderRepo *OrderRepositoryAdapter) GetOrdersSortedByPrice(offset, limit int) ([]entities.UserOrder, error) {

	query := `SELECT o.order_id, o.user_id, u.name, u.email, u.phone, o.order_date, oi.id AS item_id,
	 oi.product_id, oi.quantity FROM orders o JOIN order_items oi ON oi.order_id = o.order_id 
	 JOIN users u ON o.user_id = u.user_id ORDER BY o.price DESC OFFSET $1 LIMIT $2`

	var userOrder []entities.UserOrder
	if err := orderRepo.DB.Raw(query, offset, limit).Scan(&userOrder).Error; err != nil {
		return nil, err
	}

	return userOrder, nil
}

func (orderRepo *OrderRepositoryAdapter) GetOrdersSortedbyTime(offset, limit int) ([]entities.UserOrder, error) {

	query := `SELECT o.order_id, o.user_id, u.name, u.email, u.phone, o.order_date, oi.id AS item_id,
	 oi.product_id, oi.quantity FROM orders o JOIN order_items oi ON oi.order_id = o.order_id 
	 JOIN users u ON o.user_id = u.user_id ORDER BY o.order_date DESC OFFSET $1 LIMIT $2`

	var userOrder []entities.UserOrder
	if err := orderRepo.DB.Raw(query, offset, limit).Scan(&userOrder).Error; err != nil {
		return nil, err
	}

	return userOrder, nil
}
