package server

import (
	"errors"
	"github.com/Test-for-regression-of-the-site/trots-api/model"
	"net/http"
)

type TestRequest struct {
	*model.TestRequestPayload
}

func (testRequest *TestRequest) Bind(request *http.Request) error {
	if testRequest.Links == nil {
		return errors.New("testRequest.Links is empty")
	}
	return nil
}
