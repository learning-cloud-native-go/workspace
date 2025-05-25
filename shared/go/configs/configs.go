package configs

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
)

type BookAPI struct {
	Server Server
}

func NewBookAPI() *BookAPI {
	var c BookAPI
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}

// Types ----

type Server struct {
	Port           int           `env:"SERVER_PORT,required"`
	TimeoutRead    time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite   time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle    time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
	HandlerTimeout time.Duration `env:"SERVER_HANDLER_TIMEOUT,required"`
	Debug          bool          `env:"SERVER_DEBUG,required"`
}
