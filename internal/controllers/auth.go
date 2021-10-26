package controllers

import (
	"cloudflare-system-assessment/config"
	"cloudflare-system-assessment/internal/helper"
	"cloudflare-system-assessment/internal/middleware"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if username == "" {
		config.Logger(r.URL.Path, http.StatusBadRequest, "Bad Request")
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Bad Request"))
		if err != nil {
			fmt.Printf("Error writing response: %v\n", err)
			return
		}
		return
	}

	token, exp, err := createToken(username, time.Hour*24, r.URL.Path)
	if err != nil {
		fmt.Printf("Error at createToken: %v\n", err)
		config.Logger(r.URL.Path, http.StatusInternalServerError, "Internal Server Error")
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal Server Error"))
		if err != nil {
			fmt.Printf("Error writing response: %v\n", err)
			return
		}
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Path:    "/",
		Expires: exp,
	})

	config.Logger(r.URL.Path, http.StatusOK, "OK")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("OK"))
	if err != nil {
		fmt.Printf("Error writing response: %v\n", err)
		return
	}
	return
}

func createToken(username string, expTimeFromNow time.Duration, path string) (tokenString string, expirationTime time.Time, err error) {
	defer helper.CalculateElapsedTime(path)()
	tm := time.Now()
	now := tm
	exp := tm.Add(expTimeFromNow)

	claims := &middleware.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
			IssuedAt:  now.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	rsaPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(config.Env.JwtSecret))
	if err != nil {
		return "", exp, err
	}

	tokenString, err = token.SignedString(rsaPrivateKey)
	if err != nil {
		return "", exp, err
	}

	return tokenString, exp, nil
}
