package repository

import (
	"context"

	"github.com/sbrunettajr/ifoody-api/domain/entity"
)

type StoreRepository interface {
	Create(context context.Context, store entity.Store) (uint32, error)
	FindAll(context context.Context) ([]entity.Store, error)
	FindByCNPJ(context context.Context, CNPJ string) (entity.Store, error)
	FindByID(context context.Context, ID uint32) (entity.Store, error)
	FindByUUID(context context.Context, UUID string) (entity.Store, error)
}
