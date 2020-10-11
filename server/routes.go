package server

import (
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
	_, _ = writer.Write([]byte("Test"))
}
