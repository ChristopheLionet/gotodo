package routes

import (
	"gotodo/internal/api/handlers"
	"net/http"
)

func AddRoutes() {
	http.HandleFunc("/diagnostics/healthcheck", handlers.HandleHealthCheck)
}
