package viewmodel

import "github.com/sbrunettajr/ifoody-api/domain/entity"

type CreateOrderRequest struct {
	StoreUUID  string                       `json:"store_uuid"`
	OrderItems CreateOrderOrderItemsRequest `json:"order_items"`
}

func (vm CreateOrderRequest) ToEntity() entity.Order {
	orderItems := make([]entity.OrderItem, 0)

	for _, orderItemVM := range vm.OrderItems {
		orderItem := entity.OrderItem{
			Quantity: orderItemVM.Quantity,
			Item: entity.Item{
				UUID: orderItemVM.UUID,
			},
		}
		orderItems = append(orderItems, orderItem)
	}

	return entity.Order{
		OrderItems: orderItems,
		Store: entity.Store{
			UUID: vm.StoreUUID,
		},
	}
}
