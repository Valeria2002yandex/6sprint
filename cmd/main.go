package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	logger := log.New(
		os.Stdout,
		"[SERVER] ",
		log.LstdFlags|log.Lmicroseconds,
	)

	srv := server.NewServer(logger)

	logger.Println("Starting server on", srv.Server.Addr)
	err := srv.Server.ListenAndServe()
	if err != nil {
		logger.Fatal("Server failed to run:", err)
	}
}
