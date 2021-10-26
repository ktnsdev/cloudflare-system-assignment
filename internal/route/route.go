package route

import (
	"cloudflare-system-assessment/internal/controllers"
	"cloudflare-system-assessment/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(router *chi.Mux) {
	privateEndpoints := router.Group(nil)
	publicEndpoints := router.Group(nil)

	publicEndpoints.Get("/checkAlive", controllers.CheckAlive) // Check Alive
	publicEndpoints.Get("/auth/{username}", controllers.Auth)  // Authorisation
	publicEndpoints.Get("/verify", controllers.Verify)         // Verification

	privateEndpoints.Use(middleware.Middleware)                // Middleware
	privateEndpoints.Get("/README.txt", controllers.GetReadMe) // README.txt
	privateEndpoints.Get("/stats", controllers.GetStats)       // README.txt
}
