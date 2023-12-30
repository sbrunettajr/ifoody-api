package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/sbrunettajr/ifoody-api/domain/entity"
)

type orderItemMySQLRepository struct {
	db *sql.DB
}

func newOrderItemMySQLRepository(
	db *sql.DB,
) orderItemMySQLRepository {
	return orderItemMySQLRepository{
		db: db,
	}
}

func (r orderItemMySQLRepository) BulkInsert(context context.Context, orderID uint32, orderItems []entity.OrderItem, tx *sql.Tx) error {
	query := `
		INSERT 
		  INTO tb_order_item(uuid, quantity, item_id, order_id)
		VALUES %s  
	`

	if len(orderItems) == 0 {
		return nil
	}

	values := strings.TrimSuffix(strings.Repeat("(?, ?, ?, ?), ", len(orderItems)), ", ")
	query = fmt.Sprintf(query, values)

	var args []any
	for _, orderItem := range orderItems {
		args = append(
			args,
			orderItem.UUID,
			orderItem.Quantity,
			orderItem.Item.ID,
			orderID,
		)
	}

	_, err := tx.ExecContext(context, query, args...)
	if err != nil {
		return err
	}
	return nil
}
