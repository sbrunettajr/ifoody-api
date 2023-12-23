package test

import (
	"context"
	"io"
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

func newRequest(method, target string, body io.Reader) (*http.Request, *httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(method, target, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	return req, rec, ctx
}
