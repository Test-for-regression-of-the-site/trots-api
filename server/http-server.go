package server

import (
	"github.com/Test-for-regression-of-the-site/trots-api/constants"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

func Serve() {
	router := chi.NewRouter()
	router.Use(cors.Handler(constants.CorsOptions))
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(render.SetContentType(render.ContentTypeJSON))

	render.Respond = ErrorResponder

	router.Post(constants.TasksRoutePattern, tasksRoute)
	router.Route(constants.TasksRoutePattern, func(router chi.Router) {
		router.Route(constants.SessionIdParameterPattern, func(router chi.Router) {
			router.Use(handleTestReportsSessionId)
			router.Route(constants.TestIdParameterPattern, func(router chi.Router) {
				router.Use(handleTestReportsTestId)
				router.Get(constants.Slash, getTestReports)
			})
		})
	})

	log.Printf("Listening and serving on: %s", viper.GetString(constants.ServerAddressKey))
	if result := http.ListenAndServe(viper.GetString(constants.ServerAddressKey), router); result != nil {
		log.Printf(result.Error())
	}
}
