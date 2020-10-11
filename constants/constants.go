package constants

import "github.com/go-chi/cors"

var (
	CorsOptions = cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}
)

const (
	Trots               = "trots"
	Yml                 = "yaml"
	Dot                 = "."
	ServerAddressKey    = "trots.server.address"
	TasksRoute          = "/tasks"
	TasksDashboardRoute = "/tasks/dashboard"
)
