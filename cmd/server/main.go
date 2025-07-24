package main

import (
	"crypto-checkout-simulator/config"
	"crypto-checkout-simulator/pkg/logger"
	"crypto-checkout-simulator/server"
)

func main() {
	logger.Init()
	c := config.New()
	s := server.New(c)

	err := s.Start()
	if err != nil {
		panic(err)
	}
}
