package repository

import (
	"context"
	"database/sql"

	"github.com/sbrunettajr/ifoody-api/domain/entity"
)

type OrderItemRepository interface {
	BulkInsert(context context.Context, orderID uint32, orderItems []entity.OrderItem, tx *sql.Tx) error
}
