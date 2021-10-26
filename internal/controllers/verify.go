package controllers

import (
	"cloudflare-system-assessment/config"
	"cloudflare-system-assessment/internal/middleware"
	"fmt"
	"net/http"
)

func Verify(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		fmt.Printf("Error at Reading Cookie Token: %v\n", err)
		config.Logger(r.URL.Path, http.StatusUnauthorized, "Unauthorized: token not found")
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("Unauthorized: token not found"))
		if err != nil {
			fmt.Printf("Error writing response: %v\n", err)
			return
		}
		return
	}

	valid, claims, err := middleware.ValidateToken(token.Value, r.URL.Path)
	if err != nil || !valid {
		fmt.Printf("Error at Reading Cookie Token: %v\n", err)
		config.Logger(r.URL.Path, http.StatusUnauthorized, "Unauthorized: invalid or expired token")
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("Unauthorized: invalid or expired token"))
		if err != nil {
			fmt.Printf("Error writing response: %v\n", err)
			return
		}
		return
	}

	config.Logger(r.URL.Path, http.StatusOK, "OK")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(claims.Sub))
	if err != nil {
		fmt.Printf("Error writing response: %v\n", err)
		return
	}
}
