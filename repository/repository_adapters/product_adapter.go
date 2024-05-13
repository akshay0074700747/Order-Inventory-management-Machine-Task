package repositoryadapters

import (
	"errors"
	"time"

	customerrormessages "github.com/akshay0074700747/order-inventory_management/customError_messages"
	"github.com/akshay0074700747/order-inventory_management/entities"
	"gorm.io/gorm"
)

// this is the implementatioin of the userrepositoryport interface
type ProductRepositoryAdapter struct {
	DB *gorm.DB
}

func NewProductRepositoryAdapter(db *gorm.DB) *ProductRepositoryAdapter {
	return &ProductRepositoryAdapter{
		DB: db,
	}
}

func (productRepo *ProductRepositoryAdapter) AddProduct(product entities.Product) (entities.Product, error) {

	var result entities.Product
	//starting a transaction
	return result, productRepo.DB.Transaction(func(tx *gorm.DB) error {
		//adding new product in the table
		if err := tx.Create(&product).Scan(&result).Error; err != nil {
			return err
		}
		return nil
	})
}

func (productRepo *ProductRepositoryAdapter) UpdateProduct(product entities.Product) (entities.Product, error) {

	var updatedProduct entities.Product
	//starting a transaction
	return updatedProduct, productRepo.DB.Transaction(func(tx *gorm.DB) error {

		result := tx.Model(&entities.Product{}).Where("product_id = ?", product.ProductID).Updates(&product)

		if result.Error != nil {
			return result.Error
		}

		// checking for affected rows
		if result.RowsAffected == 0 {
			return errors.New(customerrormessages.ProductNotFoundError)
		}

		//if the stock of product also updates the updatedat column also gets updated
		if product.Stock != 0 {
			updatedRes := tx.Model(&entities.Product{}).Where("product_id = ?", product.ProductID).UpdateColumn("updated_at", time.Now())
			if updatedRes.Error != nil {
				return updatedRes.Error
			}
			if updatedRes.RowsAffected == 0 {
				return errors.New(customerrormessages.ProductNotFoundError)
			}
		}

		// fetching the updated product
		result = tx.Where("product_id = ?", product.ProductID).First(&updatedProduct)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})
}

func (productRepo *ProductRepositoryAdapter) DeleteProduct(productID string) error {

	//starting transaction
	err := productRepo.DB.Transaction(func(tx *gorm.DB) error {

		//deleting product by id
		result := productRepo.DB.Where("product_id = ?", productID).Delete(&entities.Product{})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New(customerrormessages.ProductNotFoundError)
		}
		return nil
	})

	return err
}

func (productRepo *ProductRepositoryAdapter) IncrementStock(product entities.Product) (entities.Product, error) {

	var updatedProduct entities.Product
	//strting transaction
	return updatedProduct, productRepo.DB.Transaction(func(tx *gorm.DB) error {

		//updating both the stock and updated_at column
		result := tx.Model(&entities.Product{}).
			Where("product_id = ?", product.ProductID).
			Updates(map[string]interface{}{"stock": gorm.Expr("stock + ?", product.Stock), "updated_at": time.Now()})

		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New(customerrormessages.ProductNotFoundError)
		}

		//getting the product
		result = tx.Where("product_id = ?", product.ProductID).First(&updatedProduct)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New(customerrormessages.ProductNotFoundError)
		}

		return nil
	})
}

func (productRepo *ProductRepositoryAdapter) DecrementStock(product entities.Product) (entities.Product, error) {

	var updatedProduct entities.Product
	//starting transaction
	return updatedProduct, productRepo.DB.Transaction(func(tx *gorm.DB) error {

		readRes := tx.Where("product_id = ?", product.ProductID).First(&updatedProduct)
		if readRes.Error != nil {
			return readRes.Error
		}
		//checking if the decrementing qty of stock is greater than the qty of stock
		if updatedProduct.Stock < product.Stock {
			return errors.New(customerrormessages.StockDecrementError)
		}

		result := tx.Model(&entities.Product{}).
			Where("product_id = ?", product.ProductID).
			Updates(map[string]interface{}{"stock": gorm.Expr("stock - ?", product.Stock), "updated_at": time.Now()})

		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return errors.New(customerrormessages.ProductNotFoundError)
		}

		result = tx.Where("product_id = ?", product.ProductID).First(&updatedProduct)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New(customerrormessages.ProductNotFoundError)
		}

		return nil
	})
}

func (productRepo *ProductRepositoryAdapter) GetProducts(offset, limit int) ([]entities.Product, error) {

	var products []entities.Product
	//getting the products based on the offset and limit
	result := productRepo.DB.Offset(offset).Limit(limit).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil

}

func (productRepo *ProductRepositoryAdapter) SearchProducts(namePrefix string, offset, limit int) ([]entities.Product, error) {

	var products []entities.Product

	//searching whether the namePrefix is a substring of the productName and setting offset and limit for the output
	result := productRepo.DB.Offset(offset).Limit(limit).
		Where("product_name LIKE ?", "%"+namePrefix+"%").
		Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (productRepo *ProductRepositoryAdapter) GetMostOrderedProducts(offset, limit int) ([]entities.GetProductwithOrderCount, error) {

	//uses raw query to get more control over the database
	query := `SELECT p.product_id, p.product_name, p.description, COUNT(o.order_id) AS order_count
	FROM products p INNER JOIN order_items oi ON oi.product_id = p.product_id
	INNER JOIN orders o ON o.order_id = oi.order_id GROUP BY p.product_id
	ORDER BY order_count DESC OFFSET $1 LIMIT $2;`
	var products []entities.GetProductwithOrderCount

	if err := productRepo.DB.Raw(query, offset, limit).Scan(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (productRepo *ProductRepositoryAdapter) GetProductswithLeastStocks(offset, limit int) ([]entities.Product, error) {

	var products []entities.Product
	result := productRepo.DB.
		Select("*").
		Model(&entities.Product{}).
		Order("stock ASC").
		Offset(offset).
		Limit(limit).
		Scan(&products)
	if result.Error != nil {
		return products, result.Error
	}

	return products, nil
}

func (productRepo *ProductRepositoryAdapter) GetTrendingProducts(offset, limit int) ([]entities.Product, error) {

	var products []entities.Product
	result := productRepo.DB.
		Select("p.*, COUNT(oi.order_id) AS order_count").
		Joins("JOIN order_items oi ON oi.product_id = p.product_id").
		Joins("JOIN orders o ON o.order_id = oi.order_id").
		Where("o.order_date >= ?", time.Now().AddDate(0, -1, -7)).
		Group("p.product_id").
		Order("order_count DESC").
		Offset(offset).
		Limit(limit).
		Scan(&products)
	if result.Error != nil {
		return products, result.Error
	}

	return products, nil
}
