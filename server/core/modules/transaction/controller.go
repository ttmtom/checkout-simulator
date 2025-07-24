package transaction

import (
	"crypto-checkout-simulator/server/pkg/response"
	"github.com/labstack/echo/v4"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Checkout(ctx echo.Context) error {
	return response.FailureResponse(500)
}
