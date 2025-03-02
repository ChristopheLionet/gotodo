package main

import (
	"gotodo/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/diagnostics/healthcheck", handlers.HandleHealthCheck)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
