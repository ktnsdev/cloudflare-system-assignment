package controllers

import (
	"cloudflare-system-assessment/config"
	"fmt"
	"net/http"
)

func CheckAlive(w http.ResponseWriter, r *http.Request) {
	config.Logger(r.URL.Path, http.StatusOK, "OK")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("The server is alive."))
	if err != nil {
		fmt.Printf("Error writing response: %v\n", err)
		return
	}
}
