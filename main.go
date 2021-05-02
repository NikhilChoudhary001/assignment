package main

import (
	"assignment/controller"
	"assignment/mapStore"
	"assignment/router"

	"log"
	"net/http"

	// external

	"go.uber.org/zap"
)

//Entry point of the program
func main() {
	logger, _ := zap.NewProduction() // Create Uber's Zap logger
	h := &controller.Handler{
		Repository: mapStore.NewMapStore(),
		Logger:     logger,
	}
	r := router.InitializeRoutes(h) // configure routes

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe() // Run the http server
}