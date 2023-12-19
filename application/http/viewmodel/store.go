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

type storeResponse struct {
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	UUID          string    `json:"uuid"`
	FantasyName   string    `json:"fantasy_name"`
	CorporateName string    `json:"corporate_name"`
	CNPJ          string    `json:"cnpj"`
}

type findAllStoresResponse []storeResponse

func ParseFindAllStoresReponse(stores []entity.Store) findAllStoresResponse {
	storesResponse := make(findAllStoresResponse, 0, len(stores))
	for _, store := range stores {
		storeResponse := storeResponse{
			CreatedAt:     store.CreatedAt,
			UpdatedAt:     store.UpdatedAt,
			UUID:          store.UUID,
			FantasyName:   store.FantasyName,
			CorporateName: store.CorporateName,
			CNPJ:          store.CNPJ,
		}
		storesResponse = append(storesResponse, storeResponse)
	}
	return storesResponse
}
