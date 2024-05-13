package handlers

import (
	"fmt"
	"net/http"

	"github.com/akshay0074700747/order-inventory_management/entities"
	"github.com/akshay0074700747/order-inventory_management/helpers"
	usecaseports "github.com/akshay0074700747/order-inventory_management/usecase/usecase_ports"
	"github.com/gofiber/fiber"
)

type OrderHandler struct {
	OrderUsecase usecaseports.OrderUsecasePort
}

func NewOrderHandler(orderUsecase usecaseports.OrderUsecasePort) *OrderHandler {

	return &OrderHandler{
		OrderUsecase: orderUsecase,
	}
}

func (orderHandler *OrderHandler) PlaceOrder(c *fiber.Ctx) {

	//getting the user credentials from the request context ,in it which was already set by the global middleware
	value := c.Locals("values")
	if value == nil {
		c.Status(http.StatusInternalServerError).Write("user credentials are not available")
		return
	}

	valueMap, _ := value.(map[string]interface{})

	userID := valueMap["userID"].(string)

	//checking whether the userID is empty or not
	if userID == "" {
		c.Status(http.StatusBadRequest).Write("userID is empty")
		return
	}

	var order entities.UserOrder
	if err := c.BodyParser(&order); err != nil {
		fmt.Printf("the body cannot be parsed : %s", err.Error())
		c.Status(http.StatusBadRequest).JSON(entities.Response{
			StatusCode: 400,
			Message:    "can't bind the body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if err := orderHandler.OrderUsecase.PlaceOrder(entities.Order{UserID: userID}, order.Items); err != nil {
		fmt.Printf("Error : %s", err.Error())
		c.Status(http.StatusInternalServerError).JSON(entities.Response{
			StatusCode: 500,
			Message:    "unable to place order right now",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(entities.Response{
		StatusCode: 200,
		Message:    "Placed Order Successfully",
		Data:       nil,
		Errors:     nil,
	})
}

func (orderHandler *OrderHandler) GetOrdersByuser(c *fiber.Ctx) {

	//getting the user credentials from the request context ,in it which was already set by the global middleware
	value := c.Locals("values")
	if value == nil {
		c.Status(http.StatusInternalServerError).Write("user credentials are not available")
		return
	}

	valueMap, _ := value.(map[string]interface{})

	userID := valueMap["userID"].(string)

	//checking whether the userID is empty or not
	if userID == "" {
		c.Status(http.StatusBadRequest).Write("userID is empty")
		return
	}

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

	result, err := orderHandler.OrderUsecase.GetUserOrders(userID, pageNo, limit)
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		c.Status(http.StatusInternalServerError).JSON(entities.Response{
			StatusCode: 500,
			Message:    "cannot get orders by user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(entities.Response{
		StatusCode: 200,
		Message:    "fetched orders of the user successfully",
		Data:       result,
		Errors:     nil,
	})
}

func (orderHandler *OrderHandler) GetOrdersSorted(c *fiber.Ctx) {

	// admins can sort the order based on time and price
	sorted := c.Query("sorted")

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

	var result []entities.UserOrder

	//getting the sorted results based on the user input
	if sorted == "" || sorted == "time" {
		result, err = orderHandler.OrderUsecase.GetOrdersSortedbyTime(pageNo, limit)
	} else if sorted == "price" {
		result, err = orderHandler.OrderUsecase.GetOrdersSortedByPrice(pageNo, limit)
	}

	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		c.Status(http.StatusInternalServerError).JSON(entities.Response{
			StatusCode: 500,
			Message:    "cannot get the sorted result right now",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if len(result) == 0 {
		c.Status(http.StatusBadRequest).JSON(entities.Response{
			StatusCode: 400,
			Message:    "entered filter is not valid",
			Data:       nil,
			Errors:     "enter a valid filter",
		})
		return
	}

	c.Status(http.StatusOK).JSON(entities.Response{
		StatusCode: 200,
		Message:    "fetched the sorted orders successfully",
		Data:       result,
		Errors:     nil,
	})
}

func (orderHandler *OrderHandler) GetOrdersFiltered(c *fiber.Ctx) {

	// admins can filter the order based on products and users
	filtered := c.Query("filter")
	id := c.Query("id")

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

	var result []entities.UserOrder

	//getting the filtered results based on the user input
	if filtered == "product" {
		result, err = orderHandler.OrderUsecase.GetProductOrders(id, pageNo, limit)
	} else if filtered == "user" {
		result, err = orderHandler.OrderUsecase.GetUserOrders(id, pageNo, limit)
	}

	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		c.Status(http.StatusInternalServerError).JSON(entities.Response{
			StatusCode: 500,
			Message:    "cannot get the filtered result right now",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if len(result) == 0 {
		c.Status(http.StatusBadRequest).JSON(entities.Response{
			StatusCode: 400,
			Message:    "entered filter is not valid",
			Data:       nil,
			Errors:     "enter a valid filter",
		})
		return
	}

	c.Status(http.StatusOK).JSON(entities.Response{
		StatusCode: 200,
		Message:    "fetched the filtered orders successfully",
		Data:       result,
		Errors:     nil,
	})
}

func (orderHandler *OrderHandler) GetOrdersWithDetails(c *fiber.Ctx) {

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

	result, err := orderHandler.OrderUsecase.GetOrderswithUsers(pageNo, limit)
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		c.Status(http.StatusInternalServerError).JSON(entities.Response{
			StatusCode: 500,
			Message:    "cannot get orders with users",
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
