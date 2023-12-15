package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sbrunettajr/ifoody-api/application/http/viewmodel"
	"github.com/sbrunettajr/ifoody-api/domain/service"
)

type storeController struct {
	storeService service.StoreService
}

func NewStoreController(
	storeService service.StoreService,
) storeController {
	return storeController{
		storeService: storeService,
	}
}

func (c storeController) Create(ctx echo.Context) error {
	var requestBody viewmodel.CreateStoreRequest
	if err := ctx.Bind(&requestBody); err != nil {
		return err
	}

	context := ctx.Request().Context()

	store := requestBody.ToEntity()

	err := c.storeService.Create(context, store)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusCreated)
}

func (c storeController) FindAll(ctx echo.Context) error {
	context := ctx.Request().Context()

	stores, err := c.storeService.FindAll(context)
	if err != nil {
		return err
	}

	responseBody := viewmodel.ParseFindAllStoresReponse(stores)

	return ctx.JSON(http.StatusOK, responseBody)
}
