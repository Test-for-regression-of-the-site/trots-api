package server

import (
	"github.com/Test-for-regression-of-the-site/trots-api/constants"
	"github.com/Test-for-regression-of-the-site/trots-api/provider"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func Serve() {
	router := chi.NewRouter()
	router.Use(cors.Handler(constants.CorsOptions))
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(provider.Configuration.Timeout))
	router.Use(render.SetContentType(render.ContentTypeJSON))

	render.Respond = ErrorResponder

	router.Get(constants.TasksDashboardRoutePattern, getDashboard)
	router.Route(constants.TasksRoutePattern, func(router chi.Router) {
		router.Post(constants.Slash, tasksRoute)
		router.Route(constants.SessionIdParameterPattern, func(router chi.Router) {
			router.Use(handleTestReportsSessionId)
			router.Route(constants.TestIdParameterPattern, func(router chi.Router) {
				router.Use(handleTestReportsTestId)
				router.Get(constants.Slash, getTestReports)
			})
		})
	})

	log.Printf("Listening and serving on: %s", provider.Configuration.Address)
	if result := http.ListenAndServe(provider.Configuration.Address, router); result != nil {
		log.Printf(result.Error())
	}
}
