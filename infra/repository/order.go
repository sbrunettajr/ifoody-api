package repository

import (
	"context"
	"database/sql"

	"github.com/sbrunettajr/ifoody-api/domain/entity"
	"github.com/sbrunettajr/ifoody-api/domain/repository"
)

var _ repository.OrderRepository = orderMySQLRepository{}

type orderMySQLRepository struct {
	db *sql.DB
}

func newOrderMySQLRepository(
	db *sql.DB,
) orderMySQLRepository {
	return orderMySQLRepository{
		db: db,
	}
}

func (r orderMySQLRepository) Create(context context.Context, order entity.Order, tx *sql.Tx) (uint32, error) {
	query := `
		INSERT 
		  INTO tb_order(uuid, store_id)
		VALUES (?, ?);  
	`

	result, err := tx.ExecContext(
		context,
		query,
		order.UUID,
		order.Store.ID,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint32(id), nil
}
