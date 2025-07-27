package payment

import (
	"bytes"
	paymentgateway "crypto-checkout-simulator/server/core/interfaces/payment-gateway"
	"crypto-checkout-simulator/server/core/modules/payment/dtos"
	"crypto-checkout-simulator/server/pkg/response"
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"log/slog"
	"net/http"
)

type Controller struct {
	validator      *validator.Validate
	paymentService *Service
	gateway        paymentgateway.Gateway
}

func NewController(v *validator.Validate, service *Service, gateway paymentgateway.Gateway) *Controller {
	return &Controller{v, service, gateway}
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

func (c *Controller) WebhookProcessor(ctx echo.Context) error {
	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		slog.Error("Error reading request body: %v", err)
		return ctx.String(http.StatusBadRequest, "Cannot read request body")
	}

	ctx.Request().Body = io.NopCloser(bytes.NewBuffer(body))

	isValid := c.gateway.ValidateEvent(true)
	if !isValid {
		log.Println("Invalid webhook signature")
		return response.FailureResponse(http.StatusBadRequest, "Invalid signature")
	}

	ctx.Request().Body = io.NopCloser(bytes.NewBuffer(body))

	var webhookEvent paymentgateway.CoinbaseWebhookEvent
	if err := json.NewDecoder(ctx.Request().Body).Decode(&webhookEvent); err != nil {
		log.Printf("Error decoding webhook event: %v", err)
		return response.FailureResponse(http.StatusBadRequest, "Cannot decode webhook event")
	}

	log.Printf("Successfully verified and received webhook event: %s, type: %s", webhookEvent.ID, webhookEvent.Event.Type)

	switch webhookEvent.Event.Type {
	case "charge:created":
		slog.Info("Payment created for charge:", webhookEvent.Event.Data)
		c.paymentService.WebhookPaymentCreatedHandler(webhookEvent.Event)
	case "charge:confirmed":
		slog.Info("Payment confirmed for charge:", webhookEvent.Event.Data)
		c.paymentService.WebhookPaymentConfirmedHandler(webhookEvent.Event)
	case "charge:failed":
		slog.Info("Payment failed for charge:", webhookEvent.Event.Data)
		c.paymentService.WebhookPaymentFailedHandler(webhookEvent.Event)

	default:
		slog.Info("Unhandled event type: %s", webhookEvent.Event.Type)
	}

	return response.SuccessResponse(ctx, 200, nil)
}
