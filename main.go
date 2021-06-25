package main

import (
	"crud/config"
	"crud/routes"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)



func main() {
	log.Println("Starting the HTTP server on port 8090")
	router := mux.NewRouter().StrictSlash(true)
	routes.RegisterUserRoutes(router)
	log.Fatal(http.ListenAndServe(":8090", router))

	/* graceful exit */
	go gracefulShutdown()
	forever := make(chan int)
	<-forever
}

func gracefulShutdown() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s
		fmt.Println("Sutting down gracefully.")
		// clean up here
		config.Disconnect()
		os.Exit(0)
	}()
}
