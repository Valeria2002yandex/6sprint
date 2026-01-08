package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	srv := server.NewServer(logger)

	if err := srv.Server.ListenAndServe(); err != nil {
		logger.Fatal("Server failed to run:", err)
	}
}
