package server

import (
	"github.com/Test-for-regression-of-the-site/trots-api/constants"
	"github.com/Test-for-regression-of-the-site/trots-api/service"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func tasksRoute(writer http.ResponseWriter, request *http.Request) {
	testRequest := &TestRequest{}
	if bindingError := render.Bind(request, testRequest); bindingError != nil {
		if renderError := render.Render(writer, request, InvalidRequest(bindingError)); renderError != nil {
			log.Printf(renderError.Error())
		}
		return
	}
	if renderError := render.Render(writer, request, response); renderError != nil {
		log.Printf(renderError.Error())
	}
}

func getTestReports(writer http.ResponseWriter, request *http.Request) {
	sessionId := request.Context().Value(constants.SessionIdParameter).(string)
	testId := request.Context().Value(constants.TestIdParameter).(string)
	service.GetTestReport(sessionId, testId)
}
