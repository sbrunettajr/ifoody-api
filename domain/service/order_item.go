package service

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/sbrunettajr/ifoody-api/domain/entity"
	"github.com/sbrunettajr/ifoody-api/domain/repository"
)

type OrderItemService struct {
	dataManager repository.DataManager
	itemService ItemService
}

func NewOrderItemService(
	dataManager repository.DataManager,
	itemService ItemService,
) OrderItemService {
	return OrderItemService{
		dataManager: dataManager,
		itemService: itemService,
	}
}

func (s OrderItemService) BulkInsert(context context.Context, orderID uint32, orderItems []entity.OrderItem, tx *sql.Tx) error {
	for i := 0; i < len(orderItems); i++ {
		orderItems[i].UUID = uuid.NewString()

		item, err := s.itemService.FindByUUID(context, orderItems[i].Item.UUID)
		if err != nil {
			return err
		}
		orderItems[i].Item = item
		orderItems[i].Price = item.Price * float64(orderItems[i].Quantity)
	}

	err := s.dataManager.OrderItem().BulkInsert(context, orderID, orderItems, tx)
	if err != nil {
		return err
	}

	return nil
}
