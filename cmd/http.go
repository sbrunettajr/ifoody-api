package main

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sbrunettajr/ifoody-api/application/http/controller"
	"github.com/sbrunettajr/ifoody-api/application/http/middleware"
	"github.com/sbrunettajr/ifoody-api/domain/service"
	"github.com/sbrunettajr/ifoody-api/infra/db"
	"github.com/sbrunettajr/ifoody-api/infra/repository"
)

func main() {

	db, err := db.NewDB()
	if err != nil {
		panic(err) // Fix: use panic?
	}

	dataManager := repository.NewDataManager(db)

	metricsService := service.NewMetricsService()

	storeService := service.NewStoreService(dataManager)
	categoryService := service.NewCategoryService(dataManager, storeService)
	itemService := service.NewItemService(categoryService, dataManager, storeService)

	storeController := controller.NewStoreController(storeService)
	categoryController := controller.NewCategoryController(categoryService)
	itemController := controller.NewItemController(itemService)

	e := echo.New()

	prometheusMiddleware := middleware.NewPrometheusMiddleware(metricsService)

	e.Use(prometheusMiddleware.Process)

	v1 := e.Group("/v1")

	// Implementação Realizada + Testes de Integração!

	v1.GET("/stores", storeController.FindAll)
	v1.GET("/stores/:store-uuid", storeController.FindAll)
	v1.POST("/stores", storeController.Create) // OK

	v1.GET("/:store-uuid/categories", categoryController.FindByStoreUUID)
	v1.GET("/:store-uuid/categories/:category-uuid", categoryController.FindByStoreUUID)
	v1.POST("/:store-uuid/categories", categoryController.Create)

	v1.GET("/:store-uuid/items", itemController.FindAll)
	v1.GET("/:store-uuid/categories/:category-uuid/items", itemController.FindAll)
	v1.POST("/:store-uuid/items", itemController.Create)

	v1.GET("/:store-uuid/orders", itemController.Create)
	v1.GET("/:store-uuid/orders/:order-uuid", itemController.Create)
	v1.GET("/orders", itemController.Create)
	v1.GET("/orders/:order-uuid", itemController.Create)
	v1.POST("/orders", itemController.Create)

	// appGroup := v1.Group("/app")

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	e.Logger.Fatal(e.Start(":5000"))
}
