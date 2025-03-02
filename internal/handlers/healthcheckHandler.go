package handlers

import (
	"net/http"
)

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Healthy"))
	if err != nil {
		return
	}
}
