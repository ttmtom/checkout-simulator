package payment

import (
	"crypto-checkout-simulator/server/core/modules/payment/dtos"
	"crypto-checkout-simulator/server/pkg/response"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type Controller struct {
	validator      *validator.Validate
	paymentService *Service
}

func NewController(v *validator.Validate, service *Service) *Controller {
	return &Controller{v, service}
}

func (c *Controller) Checkout(ctx echo.Context) error {
	checkout := new(dtos.CheckoutDto)
	slog.Info(checkout.Email, checkout.Amount)

	if err := ctx.Bind(checkout); err != nil {
		return response.FailureResponse(http.StatusBadRequest, err.Error())
	}

	if err := c.validator.Struct(checkout); err != nil {
		return response.FailureResponse(http.StatusBadRequest, echo.Map{
			"Message": "Invalid input",
			"Error":   err.Error(),
		})
	}

	res, err := c.paymentService.Checkout(checkout.Email, checkout.Amount)

	if err != nil {
		return response.FailureResponse(http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(ctx, 200, res)
}
