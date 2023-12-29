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
	itemsFileService := service.NewItemsFileService(categoryService, dataManager, itemService, storeService)

	storeController := controller.NewStoreController(storeService)
	categoryController := controller.NewCategoryController(categoryService)
	itemController := controller.NewItemController(itemService)
	itemsFileController := controller.NewItemsFileController(itemsFileService)

	e := echo.New()

	prometheusMiddleware := middleware.NewPrometheusMiddleware(metricsService)

	e.Use(prometheusMiddleware.Process)

	v1 := e.Group("/v1")

	// Implementado + Teste = Finalizado

	v1.GET("/stores", storeController.FindAll)                // Implementado + Teste
	v1.GET("/stores/:store-uuid", storeController.FindByUUID) // Implementado + Teste
	v1.POST("/stores", storeController.Create)                // Implementado + Teste

	v1.GET("/stores/:store-uuid/categories", categoryController.FindByStoreUUID)           // Implementado + Teste
	v1.GET("/stores/:store-uuid/categories/:category-uuid", categoryController.FindByUUID) // Implementado + Teste
	v1.POST("/stores/:store-uuid/categories", categoryController.Create)                   // Implementado + Teste

	v1.GET("/stores/:store-uuid/items", itemController.FindAll) // Implementado + Teste
	v1.POST("/stores/:store-uuid/items", itemController.Create) // Implementado + Teste

	v1.GET("/stores/:store-uuid/items-file", itemsFileController.Download) // Implementado
	v1.POST("/stores/:store-uuid/items-file", itemsFileController.Upload)  // Implementado

	// v1.GET("/orders", orderController.FindAll)
	// v1.GET("/orders/:order-uuid", orderController.FindByUUID)
	// v1.POST("/orders", orderController.Create)

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	e.Logger.Fatal(e.Start(":5000"))
}
