package usecaseports

import "github.com/akshay0074700747/order-inventory_management/entities"

//ProductUsecasePort is the abstration interface for achieving loosely coupling between dependencies
type ProductUsecasePort interface {
	AddProduct(product entities.Product) (entities.Product, error)
	UpdateProduct(product entities.Product) (entities.Product, error)
	DeleteProduct(productID string) error
	IncrementStock(product entities.Product) (entities.Product, error)
	DecrementStock(product entities.Product) (entities.Product, error)
	GetProducts(offset, limit int) ([]entities.Product, error)
	SearchProducts(namePrefix string, offset, limit int) ([]entities.Product, error)
	GetMostOrderedProducts(offset, limit int) ([]entities.GetProductwithOrderCount, error)
	GetProductswithLeastStocks(offset, limit int) ([]entities.Product, error)
	GetTrendingProducts(offset, limit int) ([]entities.Product, error)
}