package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sbrunettajr/ifoody-api/application/http/viewmodel"
	"github.com/sbrunettajr/ifoody-api/domain/service"
)

type itemController struct {
	itemService service.ItemService
}

func NewItemController(
	itemService service.ItemService,
) itemController {
	return itemController{
		itemService: itemService,
	}
}

func (c itemController) Create(ctx echo.Context) error {
	var requestBody viewmodel.CreateItemRequest
	if err := ctx.Bind(&requestBody); err != nil {
		return err
	}

	context := ctx.Request().Context()

	item := requestBody.ToEntity()

	err := c.itemService.Create(context, item)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusCreated)
}

func (c itemController) FindAll(ctx echo.Context) error {
	context := ctx.Request().Context()

	items, err := c.itemService.FindAll(context)
	if err != nil {
		return err
	}

	responseBody := viewmodel.ParseFindAllItemsResponse(items)

	return ctx.JSON(http.StatusOK, responseBody)
}
