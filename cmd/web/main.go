package main

import (
	"log"
	"net/http"
)

func main() {

	mux := routes()

	log.Println("Starting webservre on port 8080")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
