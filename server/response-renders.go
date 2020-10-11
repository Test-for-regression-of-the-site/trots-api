package server

import (
	"github.com/Test-for-regression-of-the-site/trots-api/model"
	"net/http"
)

type TestResponse struct {
	*model.TestResponsePayload
}

func (error *TestResponse) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
