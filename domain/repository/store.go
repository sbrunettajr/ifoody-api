package repository

import (
	"context"

	"github.com/sbrunettajr/ifoody-api/domain/entity"
)

type StoreRepository interface {
	Create(context context.Context, store entity.Store) error
	FindAll(context context.Context) ([]entity.Store, error)
	FindByFantasyName(context context.Context, fantasyName string) (entity.Store, error)
	FindByUUID(context context.Context, UUID string) (entity.Store, error)
}
