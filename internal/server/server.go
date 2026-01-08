package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

// NewServer создаёт новый экземпляр сервера
func NewServer(logger *log.Logger) *Server {
	// Создаём HTTP-роутер
	router := http.NewServeMux()

	// Регистрируем хендлеры
	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/upload", handlers.ConvertHandler)

	// Создаём экземпляр http.Server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Возвращаем структуру сервера
	return &Server{
		Logger: logger,
		Server: server,
	}
}
