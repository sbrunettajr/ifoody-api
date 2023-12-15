package repository

import (
	"context"

	"github.com/sbrunettajr/ifoody-api/domain/entity"
)

type ItemRepository interface {
	Create(context context.Context, item entity.Item) error
	FindAll(context context.Context) ([]entity.Item, error)
	FindByCategoryUUID(context context.Context, categoryUUID string) ([]entity.Item, error)
	FindByUUID(context context.Context, UUID string) (entity.Item, error)
	Delete(context context.Context, UUID string) error
	Update(context context.Context, item entity.Item) error
}
