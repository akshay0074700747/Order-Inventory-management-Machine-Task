package usecaseadapters

import (
	"errors"

	"github.com/akshay0074700747/order-inventory_management/entities"
	"github.com/akshay0074700747/order-inventory_management/helpers"
	repositoryports "github.com/akshay0074700747/order-inventory_management/repository/repository_ports"
)

// ProductUsecaseAdapter implements the ProductUsecasePort interface
type ProductUsecaseAdapter struct {
	Repo repositoryports.ProductRepositoryPort
}

func NewProductUsecaseAdapter(repo repositoryports.ProductRepositoryPort) *ProductUsecaseAdapter {

	return &ProductUsecaseAdapter{
		Repo: repo,
	}
}

func (productUsecase *ProductUsecaseAdapter) AddProduct(product entities.Product) (entities.Product, error) {

	//setting the productID
	product.ProductID = helpers.GenUuid()

	return productUsecase.Repo.AddProduct(product)
}

func (productUsecase *ProductUsecaseAdapter) UpdateProduct(product entities.Product) (entities.Product, error) {

	return productUsecase.Repo.UpdateProduct(product)
}

func (productUsecase *ProductUsecaseAdapter) DeleteProduct(productID string) error {

	//checking if productID is empty
	if productID == "" {
		return errors.New("the productID cannot be empty")
	}

	return productUsecase.Repo.DeleteProduct(productID)
}

func (productUsecase *ProductUsecaseAdapter) IncrementStock(product entities.Product) (entities.Product, error) {

	return productUsecase.Repo.IncrementStock(product)
}

func (productUsecase *ProductUsecaseAdapter) DecrementStock(product entities.Product) (entities.Product, error) {

	return productUsecase.Repo.DecrementStock(product)
}

func (productUsecase *ProductUsecaseAdapter) GetProducts(pageNo, limit int) ([]entities.Product, error) {

	//finding the offset and limit for pagination
	offset, limit := helpers.FindLimitandOffset(pageNo, limit)

	return productUsecase.Repo.GetProducts(offset, limit)
}

func (productUsecase *ProductUsecaseAdapter) SearchProducts(namePrefix string, pageNo, limit int) ([]entities.Product, error) {

	//checking if search value is empty or not
	if namePrefix == "" {
		return nil, errors.New("the search value cannot be empty")
	}

	offset, limit := helpers.FindLimitandOffset(pageNo, limit)

	return productUsecase.Repo.SearchProducts(namePrefix, offset, limit)
}

func (productUsecase *ProductUsecaseAdapter) GetMostOrderedProducts(pageNo, limit int) ([]entities.GetProductwithOrderCount, error) {

	offset, limit := helpers.FindLimitandOffset(pageNo, limit)

	return productUsecase.Repo.GetMostOrderedProducts(offset, limit)
}

func (productUsecase *ProductUsecaseAdapter) GetProductswithLeastStocks(pageNo, limit int) ([]entities.Product, error) {

	offset, limit := helpers.FindLimitandOffset(pageNo, limit)

	return productUsecase.Repo.GetProductswithLeastStocks(offset, limit)
}

func (productUsecase *ProductUsecaseAdapter) GetTrendingProducts(pageNo, limit int) ([]entities.Product, error) {

	offset, limit := helpers.FindLimitandOffset(pageNo, limit)

	return productUsecase.Repo.GetTrendingProducts(offset, limit)
}
