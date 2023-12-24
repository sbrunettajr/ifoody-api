package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/sbrunettajr/ifoody-api/application/http/controller"
	"github.com/sbrunettajr/ifoody-api/application/http/viewmodel"
	"github.com/sbrunettajr/ifoody-api/domain/service"
	"github.com/sbrunettajr/ifoody-api/infra/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateItem(t *testing.T) {
	t.Parallel()

	CNPJ := uuid.NewString()[:14]
	store, err := createStore(CNPJ)
	assert.NoError(t, err)

	category, err := createCategory(store)
	assert.NoError(t, err)

	name := uuid.NewString()
	requestBody := fmt.Sprintf(
		`
			{
				"name": "%s",
				"description": "%s",
				"price": 50,
				"category_uuid": "%s"
			}	
		`,
		name,
		uuid.NewString(),
		category.UUID,
	)

	_, rec, ctx := newRequest(http.MethodPost, "/v1/stores/:store-uuid/items", strings.NewReader(requestBody))
	ctx.SetParamNames("store-uuid")
	ctx.SetParamValues(store.UUID)

	dataManager := repository.NewDataManager(d)
	storeService := service.NewStoreService(dataManager)
	categoryService := service.NewCategoryService(dataManager, storeService)
	itemService := service.NewItemService(categoryService, dataManager, storeService)
	itemController := controller.NewItemController(itemService)

	err = itemController.Create(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Result().StatusCode)

	context := context.Background()
	items, err := dataManager.Item().FindByCategoryUUID(context, category.UUID)
	assert.NoError(t, err)
	assert.Equal(t, name, items[0].Name)
}

func TestFindAll(t *testing.T) {
	t.Parallel()

	CNPJ := uuid.NewString()[:14]
	store, err := createStore(CNPJ)
	assert.NoError(t, err)

	category01, err := createCategory(store)
	assert.NoError(t, err)

	_, err = createItem(store, category01)
	assert.NoError(t, err)

	category02, err := createCategory(store)
	assert.NoError(t, err)

	_, err = createItem(store, category02)
	assert.NoError(t, err)

	testCases := []struct {
		desc         string
		quantity     int
		categoryUUID string
	}{
		{
			desc:     "ShouldFindItemsByStoreUUID",
			quantity: 2,
		},
		{
			desc:         "ShouldFindItemsByCategoryUUID",
			quantity:     1,
			categoryUUID: category02.UUID,
		},
	}

	dataManager := repository.NewDataManager(d)
	storeService := service.NewStoreService(dataManager)
	categoryService := service.NewCategoryService(dataManager, storeService)
	itemService := service.NewItemService(categoryService, dataManager, storeService)
	itemController := controller.NewItemController(itemService)

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			q := make(url.Values)
			q.Set("category-uuid", tC.categoryUUID)

			_, rec, ctx := newRequest(http.MethodGet, "/v1/stores/store-uuid/items?"+q.Encode(), nil)
			ctx.SetParamNames("store-uuid")
			ctx.SetParamValues(store.UUID)

			err := itemController.FindAll(ctx)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

			body, err := io.ReadAll(rec.Result().Body)
			assert.NoError(t, err)

			var data viewmodel.FindAllItemsResponse
			err = json.Unmarshal(body, &data)
			assert.NoError(t, err)

			assert.Equal(t, tC.quantity, len(data))
		})
	}
}
