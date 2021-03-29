package main

import (
	"log"
	"main/cmd/internal/handlers"
	"net/http"
)

func main() {

	mux := routes()

	log.Println("Starting channel Listener")
	go handlers.ListenToWsChannel()
	log.Println("Starting webservre on port 8080")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
