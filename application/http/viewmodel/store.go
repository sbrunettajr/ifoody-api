package viewmodel

import (
	"time"

	"github.com/sbrunettajr/ifoody-api/domain/entity"
)

type CreateStoreRequest struct {
	FantasyName   string `json:"fantasy_name"`
	CorporateName string `json:"corporate_name"`
	CNPJ          string `json:"cnpj"`
}

func (vm CreateStoreRequest) ToEntity() entity.Store {
	return entity.Store{
		FantasyName:   vm.FantasyName,
		CorporateName: vm.CorporateName,
		CNPJ:          vm.CNPJ,
	}
}

type StoreResponse struct {
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	UUID          string    `json:"uuid"`
	FantasyName   string    `json:"fantasy_name"`
	CorporateName string    `json:"corporate_name"`
	CNPJ          string    `json:"cnpj"`
}

type FindAllStoresResponse []StoreResponse

func ParseFindAllStoresResponse(stores []entity.Store) FindAllStoresResponse {
	storesResponse := make(FindAllStoresResponse, 0, len(stores))
	for _, store := range stores {
		storeResponse := ParseFindByUUIDStoreResponse(store)
		storesResponse = append(storesResponse, storeResponse)
	}
	return storesResponse
}

func ParseFindByUUIDStoreResponse(store entity.Store) StoreResponse {
	return StoreResponse{
		CreatedAt:     store.CreatedAt,
		UpdatedAt:     store.UpdatedAt,
		UUID:          store.UUID,
		FantasyName:   store.FantasyName,
		CorporateName: store.CorporateName,
		CNPJ:          store.CNPJ,
	}
}
