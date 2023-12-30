package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sbrunettajr/ifoody-api/application/http/viewmodel"
	"github.com/sbrunettajr/ifoody-api/domain/service"
)

type orderController struct {
	orderService service.OrderService
}

func NewOrderController(
	orderService service.OrderService,
) orderController {
	return orderController{
		orderService: orderService,
	}
}

func (c orderController) Create(ctx echo.Context) error {
	var requestBody viewmodel.CreateOrderRequest
	if err := ctx.Bind(&requestBody); err != nil {
		return err
	}

	context := ctx.Request().Context()

	order := requestBody.ToEntity()

	err := c.orderService.Create(context, order)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusCreated)
}
