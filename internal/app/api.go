package api

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/farbautie/ygg/config"
	"github.com/farbautie/ygg/internal/app/handlers"
	"github.com/farbautie/ygg/internal/services"
	"github.com/farbautie/ygg/pkg/server"
	"github.com/farbautie/ygg/pkg/storage"
)

func Run(config *config.Config) {
	localStorage, err := storage.NewLocalStorage(config.Storage.Local.Path)
	if err != nil {
		log.Fatal(err)
	}

	fileService := services.NewFileService(localStorage)
	fileHandler := handlers.NewFileHandler(fileService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /upload", fileHandler.UploadFile)

	log.Println("Listening on port", config.Port)
	srv := server.New(mux, server.Port(config.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Printf("Received signal %s, shutting down...", s)
	case err := <-srv.Notify():
		log.Printf("Server stopped with error: %s", err)
	}

	if err := srv.Shutdown(); err != nil {
		log.Printf("Error shutting down server: %s", err)
	}
}
