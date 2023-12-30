package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sbrunettajr/ifoody-api/domain/entity"
	"github.com/sbrunettajr/ifoody-api/domain/repository"
)

type OrderService struct {
	dataManager      repository.DataManager
	orderItemService OrderItemService
	storeService     StoreService
}

func NewOrderService(
	dataManager repository.DataManager,
	orderItemService OrderItemService,
	storeService StoreService,
) OrderService {
	return OrderService{
		dataManager:      dataManager,
		orderItemService: orderItemService,
		storeService:     storeService,
	}
}

func (s OrderService) Create(context context.Context, order entity.Order) error {
	order.UUID = uuid.NewString()

	store, err := s.storeService.FindByUUID(context, order.Store.UUID)
	if err != nil {
		return err
	}
	order.Store = store

	tx, err := s.dataManager.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	orderID, err := s.dataManager.Order().Create(context, order, tx)
	if err != nil {
		return err
	}

	err = s.orderItemService.BulkInsert(context, orderID, order.OrderItems, tx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
