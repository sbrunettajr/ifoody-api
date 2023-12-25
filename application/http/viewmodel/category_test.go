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

func TestCreateCategoryRequest(t *testing.T) {
	t.Parallel()

	vm := CreateCategoryRequest{
		Name: uuid.NewString(),
	}

	category := vm.ToEntity()

	assert.Equal(t, vm.Name, category.Name)
}

func TestParseFindByStoreUUIDCategoriesResponse(t *testing.T) {
	t.Parallel()

	categories := []entity.Category{
		{
			UUID: uuid.NewString(),
			Name: uuid.NewString(),
		},
		{
			UUID: uuid.NewString(),
			Name: uuid.NewString(),
		},
	}

	categories[0].CreatedAt = time.Now()
	categories[0].UpdatedAt = time.Now()

	categories[1].CreatedAt = time.Now()
	categories[1].UpdatedAt = time.Now()

	response := ParseFindByStoreUUIDCategoriesResponse(categories)

	assert.Len(t, response, len(categories))

	for i, categoryResponse := range response {
		assert.Equal(t, categories[i].CreatedAt, categoryResponse.CreatedAt)
		assert.Equal(t, categories[i].UpdatedAt, categoryResponse.UpdatedAt)
		assert.Equal(t, categories[i].UUID, categoryResponse.UUID)
		assert.Equal(t, categories[i].Name, categoryResponse.Name)
	}
}

func TestParseFindByUUIDCategoryResponse(t *testing.T) {
	t.Parallel()

	category := entity.Category{
		UUID: uuid.NewString(),
		Name: uuid.NewString(),
	}
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	response := ParseFindByUUIDCategoryResponse(category)

	assert.Equal(t, category.CreatedAt, response.CreatedAt)
	assert.Equal(t, category.UpdatedAt, response.UpdatedAt)
	assert.Equal(t, category.UUID, response.UUID)
	assert.Equal(t, category.Name, response.Name)
}
