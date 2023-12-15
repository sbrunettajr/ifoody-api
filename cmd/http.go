package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sbrunettajr/ifoody-api/application/http/controller"
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

	storeService := service.NewStoreService(dataManager)
	categoryService := service.NewCategoryService(dataManager, storeService)
	itemService := service.NewItemService(categoryService, dataManager, storeService)

	storeController := controller.NewStoreController(storeService)
	categoryController := controller.NewCategoryController(categoryService)
	itemController := controller.NewItemController(itemService)

	e := echo.New()

	v1Group := e.Group("/v1")

	// Implementação Realizada + Testes de Integração!

	v1Group.GET("", storeController.FindAll)
	v1Group.GET("/:store-uuid", storeController.FindAll)
	v1Group.POST("", storeController.Create) // OK

	v1Group.GET("/:store-uuid/categories", categoryController.FindByStoreUUID)
	v1Group.GET("/:store-uuid/categories/:category-uuid", categoryController.FindByStoreUUID)
	v1Group.POST("/:store-uuid/categories", categoryController.Create)

	v1Group.GET("/:store-uuid/items", itemController.FindAll)
	v1Group.GET("/:store-uuid/categories/:category-uuid/items", itemController.FindAll)
	v1Group.POST("/:store-uuid/items", itemController.Create)

	v1Group.GET("/:store-uuid/orders", itemController.Create)
	v1Group.GET("/:store-uuid/orders/:order-uuid", itemController.Create)
	v1Group.GET("/orders", itemController.Create)
	v1Group.GET("/orders/:order-uuid", itemController.Create)
	v1Group.POST("/orders", itemController.Create)

	// appGroup := v1Group.Group("/app")

	e.Logger.Fatal(e.Start(":5000"))
}
