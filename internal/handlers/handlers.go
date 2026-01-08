package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "./index.html")
}

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	const maxUpload = 5 << 20
	r.Body = http.MaxBytesReader(w, r.Body, maxUpload)

	err := r.ParseMultipartForm(maxUpload)
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusRequestEntityTooLarge)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Ошибка при получении файла", http.StatusBadRequest)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file data", http.StatusInternalServerError)
		return
	}
	if len(data) > maxUpload {
		http.Error(w, "file too large", http.StatusRequestEntityTooLarge)
		return
	}

	convertedString, err := service.Conversion(string(data))
	if err != nil {
		http.Error(w, "Conversion failed", http.StatusUnprocessableEntity)
		return
	}

	timestamp := time.Now().UTC().Format("20060102_150405")
	fileName := timestamp + filepath.Ext(handler.Filename)

	outputFile, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "Failed to create output file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	_, err = outputFile.Write([]byte(convertedString))
	if err != nil {
		http.Error(w, "Failed to write to output file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(convertedString))
}
