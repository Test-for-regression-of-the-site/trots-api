package server

import (
	"github.com/Test-for-regression-of-the-site/trots-api/model"
	"net/http"
)

type TestRequest struct {
	*model.TestRequestPayload
}

func (testRequest *TestRequest) Bind(request *http.Request) error {
	return nil
}
