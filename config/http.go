package config

import "os"

type HttpConfig struct {
	Host string
	Port string
}

func LoadHttpConfig() *HttpConfig {

	if os.Getenv("HTTP_PORT") == "" {
		os.Setenv("HTTP_PORT", "3000")
	}

	return &HttpConfig{
		Host: os.Getenv("HTTP_HOST"),
		Port: os.Getenv("HTTP_PORT"),
	}
}
