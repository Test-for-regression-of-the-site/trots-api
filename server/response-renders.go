package server

import (
	"github.com/Test-for-regression-of-the-site/trots-api/model"
	"net/http"
)

type DashboardResponse struct {
	*model.DashboardResponsePayload
}

func (dashboardResponse *DashboardResponse) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

type TestResponse struct {
	*model.TestResponsePayload
}

func (testResponse *TestResponse) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
