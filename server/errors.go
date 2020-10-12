package server

import (
	"github.com/Test-for-regression-of-the-site/trots-api/model"
	"github.com/go-chi/render"
	"net/http"
)

type ErrorResponse struct {
	*model.ErrorResponsePayload
}

func (errorResponse *ErrorResponse) Render(writer http.ResponseWriter, request *http.Request) error {
	render.Status(request, errorResponse.Status)
	return nil
}

func ErrorResponder(writer http.ResponseWriter, request *http.Request, payload interface{}) {
	if httpError, ok := payload.(error); ok {
		if _, ok := request.Context().Value(render.StatusCtxKey).(int); !ok {
			writer.WriteHeader(http.StatusBadRequest)
		}
		render.DefaultResponder(writer, request, render.M{"message": httpError.Error()})
		return
	}

	render.DefaultResponder(writer, request, payload)
}

func InvalidRequest(requestError error) render.Renderer {
	return &ErrorResponse{&model.ErrorResponsePayload{Error: requestError, Status: http.StatusBadRequest, Message: requestError.Error()}}
}

var NotFound = &ErrorResponse{&model.ErrorResponsePayload{Status: http.StatusNotFound, Message: "Not found"}}
