package viewmodel

import (
	"time"

	"github.com/sbrunettajr/ifoody-api/domain/entity"
)

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

func (vm CreateCategoryRequest) ToEntity() entity.Category {
	return entity.Category{
		Name: vm.Name,
	}
}

type categoryResponse struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
}

type findByStoreUUIDCategoriesResponse []categoryResponse

func ParseFindByStoreUUIDCategoriesResponse(categories []entity.Category) findByStoreUUIDCategoriesResponse {
	categoriesResponse := make(findByStoreUUIDCategoriesResponse, 0, len(categories))
	for _, category := range categories {
		categoryResponse := ParseFindByUUIDCategoryResponse(category)
		categoriesResponse = append(categoriesResponse, categoryResponse)
	}
	return categoriesResponse
}

func ParseFindByUUIDCategoryResponse(category entity.Category) categoryResponse {
	return categoryResponse{
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
		UUID:      category.UUID,
		Name:      category.Name,
	}
}
