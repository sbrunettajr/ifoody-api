package viewmodel

type CreateOrderOrderItemRequest struct {
	UUID     string `json:"uuid"`
	Quantity int    `json:"quantity"`
}

type CreateOrderOrderItemsRequest []CreateOrderOrderItemRequest
