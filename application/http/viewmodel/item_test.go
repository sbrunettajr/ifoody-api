//go:build unit
// +build unit

package viewmodel

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/sbrunettajr/ifoody-api/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateItemRequest_ToEntity(t *testing.T) {
	t.Parallel()

	vm := CreateItemRequest{
		Code:         uuid.NewString(),
		Name:         uuid.NewString(),
		Description:  uuid.NewString(),
		Price:        rand.Float64(),
		CategoryUUID: uuid.NewString(),
	}

	item := vm.ToEntity()

	assert.Equal(t, vm.Code, item.Code)
	assert.Equal(t, vm.Name, item.Name)
	assert.Equal(t, vm.Description, item.Description)
	assert.Equal(t, vm.Price, item.Price)
	assert.Equal(t, vm.CategoryUUID, item.Category.UUID)
}

func TestParseFindAllItemsResponse(t *testing.T) {
	t.Parallel()

	items := []entity.Item{
		{
			Code:        uuid.NewString(),
			UUID:        uuid.NewString(),
			Name:        uuid.NewString(),
			Description: uuid.NewString(),
			Price:       rand.Float64(),
		},
		{
			Code:        uuid.NewString(),
			UUID:        uuid.NewString(),
			Name:        uuid.NewString(),
			Description: uuid.NewString(),
			Price:       rand.Float64(),
		},
	}

	items[0].CreatedAt = time.Now()
	items[0].UpdatedAt = time.Now()

	items[1].CreatedAt = time.Now()
	items[1].UpdatedAt = time.Now()

	response := ParseFindAllItemsResponse(items)

	assert.Len(t, response, len(items))

	for i, itemResponse := range response {
		assert.Equal(t, items[i].CreatedAt, itemResponse.CreatedAt)
		assert.Equal(t, items[i].UpdatedAt, itemResponse.UpdatedAt)
		assert.Equal(t, items[i].UUID, itemResponse.UUID)
		assert.Equal(t, items[i].Code, itemResponse.Code)
		assert.Equal(t, items[i].Name, itemResponse.Name)
		assert.Equal(t, items[i].Description, itemResponse.Description)
		assert.Equal(t, items[i].Price, itemResponse.Price)
	}
}
