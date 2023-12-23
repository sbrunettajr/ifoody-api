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

func TestCreateStore(t *testing.T) {
	t.Parallel()

	CNPJ := uuid.NewString()[:14]
	requestBody := fmt.Sprintf(
		`
			{
				"fantasy_name": "%s",
				"corporate_name": "%s",
				"cnpj": "%s"
			}
		`,
		uuid.NewString(),
		uuid.NewString(),
		CNPJ,
	)

	_, rec, ctx := newRequest(http.MethodPost, "/v1/stores", strings.NewReader(requestBody))
	dataManager := repository.NewDataManager(d)

	metricsService := service.NewMetricsService()

	storeService := service.NewStoreService(dataManager, metricsService)
	storeController := controller.NewStoreController(storeService)

	err := storeController.Create(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Result().StatusCode)

	context := context.Background()
	store, err := dataManager.Store().FindByCNPJ(context, CNPJ)
	assert.NoError(t, err)
	assert.Equal(t, CNPJ, store.CNPJ)
}

func TestFindAllStores(t *testing.T) {
	t.Parallel()

	CNPJ := uuid.NewString()[:14]
	_, err := createStore(CNPJ)
	assert.NoError(t, err)

	_, rec, ctx := newRequest(http.MethodGet, "/v1/stores", nil)

	dataManager := repository.NewDataManager(d)

	metricsService := service.NewMetricsService()

	storeService := service.NewStoreService(dataManager, metricsService)
	storeController := controller.NewStoreController(storeService)

	err = storeController.FindAll(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

	body, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)

	var data viewmodel.FindAllStoresResponse
	err = json.Unmarshal(body, &data)
	assert.NoError(t, err)

	assert.True(t, len(data) > 0)
}

func TestFindStoreByUUID(t *testing.T) {
	t.Parallel()

	CNPJ := uuid.NewString()[:14]
	store, err := createStore(CNPJ)
	assert.NoError(t, err)

	_, rec, ctx := newRequest(http.MethodGet, "/v1/stores/:store-uuid", nil)
	ctx.SetParamNames("store-uuid")
	ctx.SetParamValues(store.UUID)

	dataManager := repository.NewDataManager(d)

	metricsService := service.NewMetricsService()

	storeService := service.NewStoreService(dataManager, metricsService)
	storeController := controller.NewStoreController(storeService)

	err = storeController.FindByUUID(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

	body, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)

	var data viewmodel.StoreResponse
	err = json.Unmarshal(body, &data)
	assert.NoError(t, err)

	assert.Equal(t, CNPJ, data.CNPJ)
}
