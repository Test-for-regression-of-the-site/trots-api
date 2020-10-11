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
	Trots      = "trots"
	Yml        = "yaml"
	Lighthouse = "lighthouse"

	ServerAddressKey = "trots.server.address"
	TimeoutKey       = "trots.server.timeout"
	LighthouseImage  = "trots.lighthouse.image"

	Dot   = "."
	Dash  = "-"
	Slash = "/"

	SessionIdParameter = "sessionId"
	TestIdParameter    = "testId"

	TasksRoutePattern          = "/tasks"
	SessionIdParameterPattern  = "/{sessionId}"
	TestIdParameterPattern     = "/{testId}"
	TasksDashboardRoutePattern = "/tasks/dashboard"
)
