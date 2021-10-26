package middleware

import (
	"cloudflare-system-assessment/config"
	"cloudflare-system-assessment/internal/helper"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			fmt.Printf("Error at Reading Cookie Token: %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte("Unauthorized: token not found"))
			if err != nil {
				fmt.Printf("Error writing response: %v\n", err)
				return
			}
			return
		}

		valid, _, err := ValidateToken(token.Value, "middleware")
		if err != nil || !valid {
			fmt.Printf("Error at Reading Cookie Token: %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte("Unauthorized: invalid or expired token"))
			if err != nil {
				fmt.Printf("Error writing response: %v\n", err)
				return
			}
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ValidateToken(tokenString string, path string) (valid bool, claims *Claims, err error) {
	defer helper.CalculateElapsedTime(path)()
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(config.Env.JwtPublic))
	if err != nil {
		return false, &Claims{}, err
	}

	claims = &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		return false, &Claims{}, err
	}

	if !token.Valid {
		return false, &Claims{}, errors.New("invalid or expired token")
	}

	return true, claims, err
}
