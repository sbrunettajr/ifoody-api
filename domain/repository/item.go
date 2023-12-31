package repository

import (
	"context"

	"github.com/sbrunettajr/ifoody-api/domain/entity"
)

type ItemRepository interface {
	BulkInsert(context context.Context, items []entity.Item) error
	Create(context context.Context, item entity.Item) (uint32, error)
	FindByCategoryUUID(context context.Context, categoryUUID string) ([]entity.Item, error)
	FindByID(context context.Context, itemID uint32) (entity.Item, error)
	FindByStoreUUID(context context.Context, storeUUID string) ([]entity.Item, error)
	FindByStoreUUIDWithRelations(context context.Context, storeUUID string) ([]entity.Item, error)
	FindByUUID(context context.Context, UUID string) (entity.Item, error)
	Delete(context context.Context, UUID string) error
	Update(context context.Context, item entity.Item) error
}
