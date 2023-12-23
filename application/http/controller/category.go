package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sbrunettajr/ifoody-api/application/http/viewmodel"
	"github.com/sbrunettajr/ifoody-api/domain/entity"
	"github.com/sbrunettajr/ifoody-api/domain/service"
)

type categoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(
	categoryService service.CategoryService,
) categoryController {
	return categoryController{
		categoryService: categoryService,
	}
}

func (c categoryController) Create(ctx echo.Context) error {
	var requestBody viewmodel.CreateCategoryRequest
	if err := ctx.Bind(&requestBody); err != nil {
		return err
	}

	context := ctx.Request().Context()
	storeUUID := ctx.Param("store-uuid")

	category := requestBody.ToEntity()
	category.Store = entity.Store{
		UUID: storeUUID,
	}

	err := c.categoryService.Create(context, category)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusCreated)
}

func (c categoryController) FindByStoreUUID(ctx echo.Context) error {
	context := ctx.Request().Context()
	storeUUID := ctx.Param("store-uuid")

	categories, err := c.categoryService.FindByStoreUUID(context, storeUUID)
	if err != nil {
		return err
	}

	responseBody := viewmodel.ParseFindByStoreUUIDCategoriesResponse(categories)

	return ctx.JSON(http.StatusOK, responseBody)
}

func (c categoryController) FindByUUID(ctx echo.Context) error {
	context := ctx.Request().Context()
	categoryUUID := ctx.Param("category-uuid")

	category, err := c.categoryService.FindByUUID(context, categoryUUID)
	if err != nil {
		return err
	}

	responseBody := viewmodel.ParseFindByUUIDCategoryResponse(category)

	return ctx.JSON(http.StatusOK, responseBody)
}
