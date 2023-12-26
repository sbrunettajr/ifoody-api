package controller

import (
	"mime"
	"net/http"
	"strconv"

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

func (c itemsFileController) Download(ctx echo.Context) error {
	context := ctx.Request().Context()
	queryParam := ctx.QueryParam("is-template")

	var isTemplate bool
	var err error
	if queryParam != "" {
		isTemplate, err = strconv.ParseBool(queryParam)
		if err != nil {
			return err
		}
	}

	bytes, err := c.itemsFileService.Download(context, isTemplate)
	if err != nil {
		return err
	}

	filename := "items.xlsx"
	contextType := mime.TypeByExtension(".xlsx")

	ctx.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename="+filename)

	return ctx.Blob(http.StatusOK, contextType, bytes)
}
