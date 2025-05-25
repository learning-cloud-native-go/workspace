package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"configs"
)

func main() {
	c := configs.NewBookAPI()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      mux,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Println(fmt.Sprintf("Starting server :%d", c.Server.Port))
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!")
}
