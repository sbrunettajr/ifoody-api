//go:build integration
// +build integration

package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/sbrunettajr/ifoody-api/application/http/controller"
	"github.com/sbrunettajr/ifoody-api/application/http/viewmodel"
	"github.com/sbrunettajr/ifoody-api/domain/service"
	"github.com/sbrunettajr/ifoody-api/infra/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateCategory(t *testing.T) {
	t.Parallel()

	CNPJ := uuid.NewString()[:14]
	store, err := createStore(CNPJ)
	assert.NoError(t, err)

	name := uuid.NewString()
	requestBody := fmt.Sprintf(
		`
			{
				"name": "%s"
			}
		`,
		name,
	)

	_, rec, ctx := newRequest(http.MethodPost, "/v1/stores/:store-uuid/categories", strings.NewReader(requestBody))
	ctx.SetParamNames("store-uuid")
	ctx.SetParamValues(store.UUID)

	dataManager := repository.NewDataManager(d)
	storeService := service.NewStoreService(dataManager)
	categoryService := service.NewCategoryService(dataManager, storeService)
	categoryController := controller.NewCategoryController(categoryService)

	err = categoryController.Create(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Result().StatusCode)

	context := context.Background()
	categories, err := dataManager.Category().FindByStoreUUID(context, store.UUID)
	assert.NoError(t, err)
	assert.Equal(t, name, categories[0].Name)
}

func TestFindAllCategories(t *testing.T) {
	t.Parallel()

	CNPJ := uuid.NewString()[:14]
	store, err := createStore(CNPJ)
	assert.NoError(t, err)

	category, err := createCategory(store)
	assert.NoError(t, err)

	_, rec, ctx := newRequest(http.MethodGet, "/v1/stores/:store-uuid/categories", nil)
	ctx.SetParamNames("store-uuid")
	ctx.SetParamValues(store.UUID)

	dataManager := repository.NewDataManager(d)
	storeService := service.NewStoreService(dataManager)
	categoryService := service.NewCategoryService(dataManager, storeService)
	categoryController := controller.NewCategoryController(categoryService)

	err = categoryController.FindByStoreUUID(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

	body, err := io.ReadAll(rec.Result().Body)
	assert.NoError(t, err)

	var data viewmodel.FindByStoreUUIDCategoriesResponse
	err = json.Unmarshal(body, &data)
	assert.NoError(t, err)

	assert.True(t, len(data) > 0)
	assert.Equal(t, category.Name, data[0].Name)
}

func TestFindCategoryByUUID(t *testing.T) {
	t.Parallel()

	CNPJ := uuid.NewString()[:14]
	store, err := createStore(CNPJ)
	assert.NoError(t, err)

	category, err := createCategory(store)
	assert.NoError(t, err)

	_, rec, ctx := newRequest(http.MethodGet, "/v1/stores/:store-uuid/categories/category-uuid", nil)
	ctx.SetParamNames("store-uuid", "category-uuid")
	ctx.SetParamValues(store.UUID, category.UUID)

	dataManager := repository.NewDataManager(d)
	storeService := service.NewStoreService(dataManager)
	categoryService := service.NewCategoryService(dataManager, storeService)
	categoryController := controller.NewCategoryController(categoryService)

	err = categoryController.FindByUUID(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

	body, err := io.ReadAll(rec.Result().Body)
	assert.NoError(t, err)

	var data viewmodel.CategoryResponse
	err = json.Unmarshal(body, &data)
	assert.NoError(t, err)

	assert.Equal(t, category.UUID, data.UUID)
}
