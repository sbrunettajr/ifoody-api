//go:build unit
// +build unit

package viewmodel

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/sbrunettajr/ifoody-api/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateStoreRequest_ToEntity(t *testing.T) {
	t.Parallel()

	vm := CreateStoreRequest{
		FantasyName:   uuid.NewString(),
		CorporateName: uuid.NewString(),
		CNPJ:          uuid.NewString(),
	}

	store := vm.ToEntity()

	assert.Equal(t, vm.FantasyName, store.FantasyName)
	assert.Equal(t, vm.CorporateName, store.CorporateName)
	assert.Equal(t, vm.CNPJ, store.CNPJ)
}

func TestParseFindAllStoresResponse(t *testing.T) {
	t.Parallel()

	stores := []entity.Store{
		{
			UUID:          uuid.NewString(),
			FantasyName:   uuid.NewString(),
			CorporateName: uuid.NewString(),
			CNPJ:          uuid.NewString(),
		},
		{
			UUID:          uuid.NewString(),
			FantasyName:   uuid.NewString(),
			CorporateName: uuid.NewString(),
			CNPJ:          uuid.NewString(),
		},
	}

	stores[0].CreatedAt = time.Now()
	stores[0].UpdatedAt = time.Now()

	stores[1].CreatedAt = time.Now()
	stores[1].UpdatedAt = time.Now()

	response := ParseFindAllStoresResponse(stores)

	assert.Len(t, response, len(stores))

	for i, storeResponse := range response {
		assert.Equal(t, stores[i].CreatedAt, storeResponse.CreatedAt)
		assert.Equal(t, stores[i].UpdatedAt, storeResponse.UpdatedAt)
		assert.Equal(t, stores[i].UUID, storeResponse.UUID)
		assert.Equal(t, stores[i].FantasyName, storeResponse.FantasyName)
		assert.Equal(t, stores[i].CorporateName, storeResponse.CorporateName)
		assert.Equal(t, stores[i].CNPJ, storeResponse.CNPJ)
	}
}

func TestParseFindByUUIDStoreResponse(t *testing.T) {
	t.Parallel()

	store := entity.Store{
		UUID:          uuid.NewString(),
		FantasyName:   uuid.NewString(),
		CorporateName: uuid.NewString(),
		CNPJ:          uuid.NewString(),
	}

	store.CreatedAt = time.Now()
	store.UpdatedAt = time.Now()

	response := ParseFindByUUIDStoreResponse(store)

	assert.Equal(t, store.CreatedAt, response.CreatedAt)
	assert.Equal(t, store.UpdatedAt, response.UpdatedAt)
	assert.Equal(t, store.UUID, response.UUID)
	assert.Equal(t, store.FantasyName, response.FantasyName)
	assert.Equal(t, store.CorporateName, response.CorporateName)
	assert.Equal(t, store.CNPJ, response.CNPJ)
}
