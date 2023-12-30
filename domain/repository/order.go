package repository

import (
	"context"
	"database/sql"

	"github.com/sbrunettajr/ifoody-api/domain/entity"
)

type OrderRepository interface {
	Create(context context.Context, order entity.Order, tx *sql.Tx) (uint32, error)
}
