package handlers

import (
	"fmt"
	"net/http"

	"github.com/akshay0074700747/order-inventory_management/entities"
	"github.com/akshay0074700747/order-inventory_management/helpers"
	usecaseports "github.com/akshay0074700747/order-inventory_management/usecase/usecase_ports"
	"github.com/gofiber/fiber"
)

type ProductHandler struct {
	ProductUsecase usecaseports.ProductUsecasePort
}

func NewProductHandler(productUsecase usecaseports.ProductUsecasePort) *ProductHandler {

	return &ProductHandler{
		ProductUsecase: productUsecase,
	}
}

func (productHandler *ProductHandler) AddProduct(c *fiber.Ctx) {

	//bibding the body of the request with the product struct
	var product entities.Product
	if err := c.BodyParser(&product); err != nil {
		fmt.Printf("the body cannot be parsed : %s", err.Error())
		c.Status(http.StatusBadRequest).JSON(entities.Response{
			StatusCode: 400,
			Message:    "can't bind the body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	result, err := productHandler.ProductUsecase.AddProduct(product)
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		c.Status(http.StatusInternalServerError).JSON(entities.Response{
			StatusCode: 500,
			Message:    "unable to add product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(entities.Response{
		StatusCode: 200,
		Message:    "Added product successfully",
		Data:       result,
		Errors:     nil,
	})
}

func (productHandler *ProductHandler) UpdatedProduct(c *fiber.Ctx) {

	var product entities.Product
	if err := c.BodyParser(&product); err != nil {
		fmt.Printf("the body cannot be parsed : %s", err.Error())
		c.Status(http.StatusBadRequest).JSON(entities.Response{
			StatusCode: 400,
			Message:    "can't bind the body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	result, err := productHandler.ProductUsecase.UpdateProduct(product)
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		c.Status(http.StatusInternalServerError).JSON(entities.Response{
			StatusCode: 500,
			Message:    "unable to update product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(entities.Response{
		StatusCode: 200,
		Message:    "updated product successfully",
		Data:       result,
		Errors:     nil,
	})
}

func (productHandler *ProductHandler) DeleteProduct(c *fiber.Ctx) {

	//getting the id of the product to delete from the parameter
	err := productHandler.ProductUsecase.DeleteProduct(c.Params("productID"))
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		c.Status(http.StatusInternalServerError).JSON(entities.Response{
			StatusCode: 500,
			Message:    "unable to delete product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(entities.Response{
		StatusCode: 200,
		Message:    "deleted product successfully",
		Data:       nil,
		Errors:     nil,
	})

}

func (productHandler *ProductHandler) IncrementStock(c *fiber.Ctx) {

	//binding the request body with the product struct
	var product entities.Product
	if err := c.BodyParser(&product); err != nil {
		fmt.Printf("the body cannot be parsed : %s", err.Error())
		c.Status(http.StatusBadRequest).JSON(entities.Response{
			StatusCode: 400,
			Message:    "can't bind the body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	result, err := productHandler.ProductUsecase.IncrementStock(product)
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		c.Status(http.StatusInternalServerError).JSON(entities.Response{
			StatusCode: 500,
			Message:    "unable to increment stock",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(entities.Response{
		StatusCode: 200,
		Message:    "incremented product stock successfully",
		Data:       result,
		Errors:     nil,
	})
}

func (productHandler *ProductHandler) DecrementStock(c *fiber.Ctx) {

	//binding the request body with the product struct
	var product entities.Product
	if err := c.BodyParser(&product); err != nil {
		fmt.Printf("the body cannot be parsed : %s", err.Error())
		c.Status(http.StatusBadRequest).JSON(entities.Response{
			StatusCode: 400,
			Message:    "can't bind the body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	result, err := productHandler.ProductUsecase.DecrementStock(product)
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		c.Status(http.StatusInternalServerError).JSON(entities.Response{
			StatusCode: 500,
			Message:    "unable to decrement stock",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(entities.Response{
		StatusCode: 200,
		Message:    "decremented product stock successfully",
		Data:       result,
		Errors:     nil,
	})
}

func (productHandler *ProductHandler) GetProducts(c *fiber.Ctx) {

	//getting the pageNo and limit from query for pagination
	pageNostr := c.Query("pageNo")
	limitStr := c.Query("limit")
	var pageNo, limit int
	var err error

	if pageNostr != "" {
		pageNo, err = helpers.StrtoInt(pageNostr)
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(entities.Response{
				StatusCode: 400,
				Message:    "cannot parse query to int",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}

	if limitStr != "" {
		limit, err = helpers.StrtoInt(limitStr)
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(entities.Response{
				StatusCode: 400,
				Message:    "cannot parse query to int",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}

	result, err := productHandler.ProductUsecase.GetProducts(pageNo, limit)
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		c.Status(http.StatusInternalServerError).JSON(entities.Response{
			StatusCode: 500,
			Message:    "cannot get products",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(entities.Response{
		StatusCode: 200,
		Message:    "fetched products successfully",
		Data:       result,
		Errors:     nil,
	})

}

func (productHandler *ProductHandler) SearchProducts(c *fiber.Ctx) {

	//getting the pageNo and limit from query for pagination
	pageNostr := c.Query("pageNo")
	limitStr := c.Query("limit")
	var pageNo, limit int
	var err error

	if pageNostr != "" {
		pageNo, err = helpers.StrtoInt(pageNostr)
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(entities.Response{
				StatusCode: 400,
				Message:    "cannot parse query to int",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}

	if limitStr != "" {
		limit, err = helpers.StrtoInt(limitStr)
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(entities.Response{
				StatusCode: 400,
				Message:    "cannot parse query to int",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}

	//getting the substring to search for from the parameter
	result, err := productHandler.ProductUsecase.SearchProducts(c.Params("namePrefix"), pageNo, limit)
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		c.Status(http.StatusInternalServerError).JSON(entities.Response{
			StatusCode: 500,
			Message:    "cannot get products",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(entities.Response{
		StatusCode: 200,
		Message:    "fetched products successfully",
		Data:       result,
		Errors:     nil,
	})
}

func (productHandler *ProductHandler) GetMostOrderedProducts(c *fiber.Ctx) {

	pageNostr := c.Query("pageNo")
	limitStr := c.Query("limit")
	var pageNo, limit int
	var err error

	if pageNostr != "" {
		pageNo, err = helpers.StrtoInt(pageNostr)
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(entities.Response{
				StatusCode: 400,
				Message:    "cannot parse query to int",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}

	if limitStr != "" {
		limit, err = helpers.StrtoInt(limitStr)
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(entities.Response{
				StatusCode: 400,
				Message:    "cannot parse query to int",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}

	result, err := productHandler.ProductUsecase.GetMostOrderedProducts(pageNo, limit)
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		c.Status(http.StatusInternalServerError).JSON(entities.Response{
			StatusCode: 500,
			Message:    "cannot get most ordered products",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(entities.Response{
		StatusCode: 200,
		Message:    "fetched most ordered products successfully",
		Data:       result,
		Errors:     nil,
	})
}

func (productHandler *ProductHandler) GetProductswithLeastStocks(c *fiber.Ctx) {

	pageNostr := c.Query("pageNo")
	limitStr := c.Query("limit")
	var pageNo, limit int
	var err error

	if pageNostr != "" {
		pageNo, err = helpers.StrtoInt(pageNostr)
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(entities.Response{
				StatusCode: 400,
				Message:    "cannot parse query to int",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}

	if limitStr != "" {
		limit, err = helpers.StrtoInt(limitStr)
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(entities.Response{
				StatusCode: 400,
				Message:    "cannot parse query to int",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}

	//getting the substring to search for from the parameter
	result, err := productHandler.ProductUsecase.GetProductswithLeastStocks(pageNo, limit)
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		c.Status(http.StatusInternalServerError).JSON(entities.Response{
			StatusCode: 500,
			Message:    "cannot get products with least stocks",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(entities.Response{
		StatusCode: 200,
		Message:    "fetched products with least stocks successfully",
		Data:       result,
		Errors:     nil,
	})
}

func (productHandler *ProductHandler) GetTrendingProducts(c *fiber.Ctx) {

	pageNostr := c.Query("pageNo")
	limitStr := c.Query("limit")
	var pageNo, limit int
	var err error

	if pageNostr != "" {
		pageNo, err = helpers.StrtoInt(pageNostr)
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(entities.Response{
				StatusCode: 400,
				Message:    "cannot parse query to int",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}

	if limitStr != "" {
		limit, err = helpers.StrtoInt(limitStr)
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(entities.Response{
				StatusCode: 400,
				Message:    "cannot parse query to int",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}

	//getting the substring to search for from the parameter
	result, err := productHandler.ProductUsecase.GetTrendingProducts(pageNo, limit)
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		c.Status(http.StatusInternalServerError).JSON(entities.Response{
			StatusCode: 500,
			Message:    "cannot get trending products",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(entities.Response{
		StatusCode: 200,
		Message:    "fetched trending products successfully",
		Data:       result,
		Errors:     nil,
	})
}
