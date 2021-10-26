package main

import (
	"cloudflare-system-assessment/config"
	"cloudflare-system-assessment/internal/route"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	err := config.Load()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}
	initRoutes()
}

func initRoutes() {
	router := chi.NewRouter()
	route.SetupRoutes(router)

	fmt.Printf("Starting the server on port %s\n", config.Env.Port)
	err := http.ListenAndServe(config.Env.Port, router)
	if err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
		return
	}
}
