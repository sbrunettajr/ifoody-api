package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sbrunettajr/ifoody-api/domain/service"
)

type itemsFileController struct {
	itemsFileService service.ItemsFileService
}

func NewItemsFileController(
	itemsFileService service.ItemsFileService,
) itemsFileController {
	return itemsFileController{
		itemsFileService: itemsFileService,
	}
}

func (c itemsFileController) Upload(ctx echo.Context) error {
	file, err := ctx.FormFile("items")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	context := ctx.Request().Context()
	storeUUID := ctx.Param("store-uuid")

	err = c.itemsFileService.Upload(context, storeUUID, src)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}
