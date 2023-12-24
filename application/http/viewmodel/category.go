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

type CategoryResponse struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
}

type FindByStoreUUIDCategoriesResponse []CategoryResponse

func ParseFindByStoreUUIDCategoriesResponse(categories []entity.Category) FindByStoreUUIDCategoriesResponse {
	categoriesResponse := make(FindByStoreUUIDCategoriesResponse, 0, len(categories))
	for _, category := range categories {
		categoryResponse := ParseFindByUUIDCategoryResponse(category)
		categoriesResponse = append(categoriesResponse, categoryResponse)
	}
	return categoriesResponse
}

func ParseFindByUUIDCategoryResponse(category entity.Category) CategoryResponse {
	return CategoryResponse{
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
		UUID:      category.UUID,
		Name:      category.Name,
	}
}
