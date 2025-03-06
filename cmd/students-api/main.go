package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Piyu-Pika/students-api/internal/config"
	"github.com/Piyu-Pika/students-api/internal/http/handlers/student"
	"github.com/Piyu-Pika/students-api/internal/storage/sqlite"
)

func main() {
	fmt.Println("Welcome to Students API")
	//LoadConfig
	cfg := config.MustLoad()
	fmt.Println(cfg)
	//database setup
	_,err:=sqlite.New(cfg)
	if err!=nil{
		log.Fatal(err)
	}
	slog.Info("Database setup successfully",slog.String("database",cfg.StoragePath))
	// router setup
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New())
	// setup server

	server := &http.Server{
		Addr:    cfg.HttpServer.Address,
		Handler: router,
	}

	slog.Info("Starting server", "address", cfg.HttpServer.Address)
	log.Println("Server started")

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	slog.Info("Server stopped")
	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("Error shutting down server", "error", err)
	}

	slog.Info("Server exited properly")
}
