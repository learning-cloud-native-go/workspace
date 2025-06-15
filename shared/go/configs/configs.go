package configs

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
)

type BookAPI struct {
	Server Server
	DB     DB
}

func NewBookAPI() *BookAPI {
	var c BookAPI
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}

func NewDB() *DB {
	var c DB
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

type DB struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	Username string `env:"DB_USER,required"`
	Password string `env:"DB_PASS,required"`
	DBName   string `env:"DB_NAME,required"`
	Debug    bool   `env:"DB_DEBUG,required"`
}
