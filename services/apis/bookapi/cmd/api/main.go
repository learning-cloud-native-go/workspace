package main

import (
	"fmt"
	"net/http"

	"configs"
	"logger"

	"bookapi/app/router"
)

func main() {
	c := configs.NewBookAPI()
	l := logger.New(c.Server.Debug)

	r := router.New(l)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	l.Info().Msgf("Starting server %v", s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		l.Fatal().Err(err).Msg("Server startup failure")
	}
}
