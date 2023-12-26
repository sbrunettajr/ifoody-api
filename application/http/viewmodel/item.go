package viewmodel

import (
	"time"

	"github.com/sbrunettajr/ifoody-api/domain/entity"
)

type CreateItemRequest struct {
	Code         string  `json:"code"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	CategoryUUID string  `json:"category_uuid"`
}

func (vm CreateItemRequest) ToEntity() entity.Item {
	return entity.Item{
		Code:        vm.Code,
		Name:        vm.Name,
		Description: vm.Description,
		Price:       vm.Price,
		Category: entity.Category{
			UUID: vm.CategoryUUID,
		},
	}
}

type itemResponse struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UUID        string    `json:"uuid"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
}

type FindAllItemsResponse []itemResponse

func ParseFindAllItemsResponse(items []entity.Item) FindAllItemsResponse {
	itemsResponse := make(FindAllItemsResponse, 0, len(items))
	for _, item := range items {
		itemResponse := itemResponse{
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			UUID:        item.UUID,
			Code:        item.Code,
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
		}
		itemsResponse = append(itemsResponse, itemResponse)
	}
	return itemsResponse
}
