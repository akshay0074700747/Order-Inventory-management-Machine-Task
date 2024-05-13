package controller

import (
	"fmt"
	"log"

	"github.com/akshay0074700747/order-inventory_management/web/handlers"
	"github.com/akshay0074700747/order-inventory_management/web/middlewares"
	"github.com/gofiber/fiber"
)

type Controller struct {
	controllers *fiber.App
}

func NewController(userHandlers *handlers.UserHandler,
	productHandlers *handlers.ProductHandler,
	orderHandlers *handlers.OrderHandler,
	middleware *middlewares.Middleware) *Controller {

	//initialising the fiber server
	app := fiber.New()

	//setting up routes
	app.Post("/signup", userHandlers.SignUp)
	app.Post("/login", userHandlers.Login)

	user := app.Group("/user", middleware.GlobalMiddleware)

	user.Get("/products", productHandlers.GetProducts)
	user.Get("/products/search/:namePrefix", productHandlers.SearchProducts)

	user.Post("/order", orderHandlers.PlaceOrder)
	user.Get("/dashboard", orderHandlers.GetOrdersByuser)

	admin := user.Group("/admin", middleware.AdminMiddleware)

	admin.Get("/users/statistics", userHandlers.GetMostOrderedUsers)
	admin.Post("/product/add", productHandlers.AddProduct)

	admin.Patch("/product/update", productHandlers.UpdatedProduct)
	admin.Delete("/product/delete/:productID", productHandlers.DeleteProduct)

	admin.Patch("/product/increment", productHandlers.IncrementStock)
	admin.Patch("/product/decrement", productHandlers.DecrementStock)

	admin.Get("/products/statistics/ordered", productHandlers.GetMostOrderedProducts)
	admin.Get("/product/statistics/stocks", productHandlers.GetProductswithLeastStocks)

	admin.Get("/product/statistics/trending", productHandlers.GetTrendingProducts)
	admin.Get("/orders/filtered", orderHandlers.GetOrdersFiltered)

	admin.Get("/orders/sorted", orderHandlers.GetOrdersSorted)
	admin.Get("/orders/detailed", orderHandlers.GetOrdersWithDetails)

	return &Controller{
		controllers: app,
	}
}

// starts the server in the specied port
func (ctrl *Controller) Start(port string) {

	fmt.Printf("Server is Starting on port %s ...", port)

	//starting up the server
	if err := ctrl.controllers.Listen(port); err != nil {
		log.Fatal(err)
	}
}
