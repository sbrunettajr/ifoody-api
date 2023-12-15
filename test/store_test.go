//go:build integration
// +build integration

package test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sbrunettajr/ifoody-api/application/http/controller"
	"github.com/sbrunettajr/ifoody-api/domain/service"
	"github.com/sbrunettajr/ifoody-api/infra/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateStore(t *testing.T) {
	t.Parallel()

	fantasyName := uuid.NewString()
	requestBody := fmt.Sprintf(
		`
			{
				"fantasy_name": "%s",
				"corporate_name": "%s"
			}
		`,
		fantasyName,
		uuid.NewString(),
	)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/v1/stores", strings.NewReader(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	dataManager := repository.NewDataManager(d)
	storeService := service.NewStoreService(dataManager)
	storeController := controller.NewStoreController(storeService)

	err := storeController.Create(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Result().StatusCode)

	context := context.Background()
	store, err := dataManager.Store().FindByFantasyName(context, fantasyName)
	assert.NoError(t, err)
	assert.Equal(t, fantasyName, store.FantasyName)
}
