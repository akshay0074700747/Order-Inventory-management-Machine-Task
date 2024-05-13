package dependencyinjection

import (
	"github.com/akshay0074700747/order-inventory_management/config"
	"github.com/akshay0074700747/order-inventory_management/db"
	repositoryadapters "github.com/akshay0074700747/order-inventory_management/repository/repository_adapters"
	usecaseadapters "github.com/akshay0074700747/order-inventory_management/usecase/usecase_adapters"
	"github.com/akshay0074700747/order-inventory_management/web/controller"
	"github.com/akshay0074700747/order-inventory_management/web/handlers"
	"github.com/akshay0074700747/order-inventory_management/web/middlewares"
)

// injecting all the dependencies from bottom up
func InjectDependencies(cfg config.Configurations) *controller.Controller {

	DB := db.ConnectDB(cfg)
	db.ExecTrigger(DB)
	db.InitializeAdmin(cfg.AdminEmail, cfg.AdminPassword, DB)

	userRepo := repositoryadapters.NewUserRepositoryAdapter(DB)
	productRepo := repositoryadapters.NewProductRepositoryAdapter(DB)
	orderRepo := repositoryadapters.NewOrderRepositoryAdapter(DB)

	userUsecase := usecaseadapters.NewUserUsecaseAdapter(userRepo)
	productUsecase := usecaseadapters.NewProductUsecaseAdapter(productRepo)
	orderUsecase := usecaseadapters.NewOrderUsecaseAdpater(orderRepo)

	userHandler := handlers.NewUserHandler(userUsecase, cfg.Secreet)
	productHandler := handlers.NewProductHandler(productUsecase)
	orderHandler := handlers.NewOrderHandler(orderUsecase)
	middleware := middlewares.NewMiddleware(cfg.Secreet)

	return controller.NewController(
		userHandler,
		productHandler,
		orderHandler,
		middleware,
	)
}
