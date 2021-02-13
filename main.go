package main

import (
	"net/http"
	"pismo-challenge-go/handlers"
	"pismo-challenge-go/server"
)

func main() {
	server := server.CreateServer()

	server.HandleFunc("/accounts", handlers.CreateAccount)
	http.ListenAndServe(":8080", server)
}
