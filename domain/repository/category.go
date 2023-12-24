package repository

import (
	"context"

	"github.com/sbrunettajr/ifoody-api/domain/entity"
)

type CategoryRepository interface {
	Create(context context.Context, category entity.Category) (uint32, error)
	FindByID(context context.Context, ID uint32) (entity.Category, error)
	FindByStoreUUID(context context.Context, storeUUID string) ([]entity.Category, error)
	FindByUUID(context context.Context, UUID string) (entity.Category, error)
}
