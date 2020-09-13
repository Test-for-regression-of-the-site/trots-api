package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/cors"
)

func (api *API) initRoutes() {
	api.routes.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	api.routes.Get("/hello", func(rw http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(rw, "hello, %q", req.RemoteAddr)
	})

	var lt = newLighthouseTasks()
	lt.cfg = api.cfg.Lighthouse

	api.routes.Get("/tasks", lt.Status)
	api.routes.Post("/startTest", lt.AddTask)
}
