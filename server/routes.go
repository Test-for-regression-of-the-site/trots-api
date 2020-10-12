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
	service.RunTest(*testRequest.TestRequestPayload)
	writer.WriteHeader(http.StatusOK)
}

func getTestReports(writer http.ResponseWriter, request *http.Request) {
	sessionId := request.Context().Value(constants.SessionIdParameter).(string)
	testId := request.Context().Value(constants.TestIdParameter).(string)
	report := service.GetTestReport(sessionId, testId)
	if report != nil {
		render.JSON(writer, request, report)
	}
}

func getDashboard(writer http.ResponseWriter, request *http.Request) {
	dashboard := service.GetDashboard()
	if dashboard != nil {
		if renderError := render.Render(writer, request, &DashboardResponse{dashboard}); renderError != nil {
			log.Printf(renderError.Error())
		}
	}
}
