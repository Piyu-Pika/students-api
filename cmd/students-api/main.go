package main

import (
	"fmt"
	"log"
	"net/http"

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
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	
}
