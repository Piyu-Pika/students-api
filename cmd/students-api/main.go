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
)

func main() {
	fmt.Println("Welcome to Students API")
	//LoadConfig
	cfg := config.MustLoad()
	fmt.Println(cfg)
	//database setup
	// router setup
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Students API"))
		fmt.Fprintf(w, "Hello World")
	})
	// setup server

	server := &http.Server{
		Addr:    cfg.HttpServer.Address,
		Handler: router,
	}
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
	err:=server.Shutdown(ctx)
	if err != nil {
		slog.Fatal(err)
	}

}
