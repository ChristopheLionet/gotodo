package main

import (
	"gotodo/internal/api/routes"
	"log"
	"net/http"
)

func main() {
	routes.AddRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
