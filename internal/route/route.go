package route

import (
	"cloudflare-system-assessment/internal/controllers"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(router *chi.Mux) {
	router.Get("/checkAlive", controllers.CheckAlive) // Check Alive
	router.Get("/auth/{username}", controllers.Auth)  // Authorisation
	router.Get("/verify", controllers.Verify)         // Verification

	router.Get("/README.txt", controllers.GetReadMe) // README.txt
	router.Get("/stats", controllers.GetStats)       // Statistics
}
