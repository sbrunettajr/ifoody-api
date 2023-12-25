//go:build integration
// +build integration

package test

import (
	"context"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sbrunettajr/ifoody-api/domain/entity"
	"github.com/sbrunettajr/ifoody-api/infra/repository"
)

func createStore(CNPJ string) (entity.Store, error) {
	store := entity.Store{
		CNPJ: CNPJ,
		UUID: uuid.NewString(),
	}

	context := context.Background()

	r := repository.NewDataManager(d)

	storeID, err := r.Store().Create(context, store)
	if err != nil {
		return entity.Store{}, err
	}

	store, err = r.Store().FindByID(context, storeID)
	if err != nil {
		return entity.Store{}, err
	}

	return store, nil
}

func createCategory(store entity.Store) (entity.Category, error) {
	category := entity.Category{
		Name:  uuid.NewString(),
		UUID:  uuid.NewString(),
		Store: store,
	}

	context := context.Background()

	r := repository.NewDataManager(d)

	categoryID, err := r.Category().Create(context, category)
	if err != nil {
		return entity.Category{}, err
	}

	category, err = r.Category().FindByID(context, categoryID)
	if err != nil {
		return entity.Category{}, err
	}

	return category, nil
}

func createItem(store entity.Store, category entity.Category) (entity.Item, error) {
	item := entity.Item{
		Name:        uuid.NewString(),
		Description: uuid.NewString(),
		Price:       rand.Float64(),
		Category:    category,
		Store:       store,
	}

	context := context.Background()

	r := repository.NewDataManager(d)

	itemID, err := r.Item().Create(context, item, nil)
	if err != nil {
		return entity.Item{}, err
	}

	item, err = r.Item().FindByID(context, itemID)
	if err != nil {
		return entity.Item{}, err
	}

	return item, nil
}

func newRequest(method, target string, body io.Reader) (*http.Request, *httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(method, target, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	return req, rec, ctx
}
