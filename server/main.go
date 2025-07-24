package server

import (
	"crypto-checkout-simulator/config"
	"crypto-checkout-simulator/server/adapter/storage/pg"
	"crypto-checkout-simulator/server/core/modules"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct {
	echo   *echo.Echo
	config *config.Config
}

func New(c *config.Config) *Server {
	storage, err := pg.New(c.Database)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	v := validator.New()

	mc := modules.InitModuleContainer(&storage, v)
	tc := mc.GetPaymentModule().GetController()

	{
		e.GET("/health/", func(e echo.Context) error {
			return e.String(http.StatusOK, "OK")
		})
	}

	checkout := e.Group("/checkout/")
	{
		checkout.POST("", tc.Checkout)
	}

	return &Server{
		e,
		c,
	}
}

func (s *Server) Start() *error {
	err := s.echo.Start(fmt.Sprintf("%s:%s", s.config.Http.Host, s.config.Http.Port))

	return &err
}
