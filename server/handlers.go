package server

import (
	"context"
	"github.com/Test-for-regression-of-the-site/trots-api/constants"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func handleTestReportsSessionId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if sessionId := chi.URLParam(request, constants.SessionIdParameter); sessionId != "" {
			next.ServeHTTP(writer, request.WithContext(context.WithValue(request.Context(), constants.SessionIdParameter, sessionId)))
			return
		}
		if renderError := render.Render(writer, request, NotFound); renderError != nil {
			log.Printf(renderError.Error())
		}
	})
}

func handleTestReportsTestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if testId := chi.URLParam(request, constants.TestIdParameter); testId != "" {
			next.ServeHTTP(writer, request.WithContext(context.WithValue(request.Context(), constants.TestIdParameter, testId)))
			return
		}
		if renderError := render.Render(writer, request, NotFound); renderError != nil {
			log.Printf(renderError.Error())
		}
	})
}
