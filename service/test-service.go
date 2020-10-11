package service

import (
	"bytes"
	"github.com/Test-for-regression-of-the-site/trots-api/configuration"
	"github.com/Test-for-regression-of-the-site/trots-api/lighthouse"
	model "github.com/Test-for-regression-of-the-site/trots-api/model"
	"github.com/google/uuid"
	"log"
)

func RunTest(request model.TestRequestPayload) *model.TestResponsePayload {
	var id = uuid.New().String()

	saveTask := func(containerId string, reportBuffer *bytes.Buffer) {
		// TODO: Put into mongo
	}

	runTask := func(url string) {
		reportBuffer := &bytes.Buffer{}
		lighthouseExecutionConfiguration := configuration.LighthouseExecutionConfiguration{
			Image:       "",
			Arguments:   nil,
			Environment: nil,
		}
		containerId, lighthouseError := lighthouse.ExecuteLighthouseTask(lighthouseExecutionConfiguration, url, reportBuffer)
		if lighthouseError != nil {
			return
		}
		defer saveTask(containerId, reportBuffer)
	}

	for _, url := range request.Links {
		go runTask(url)
	}

	return &model.TestResponsePayload{Id: id}
}

func GetTestReports(sessionId string, testId string) *model.ReportsPayload {
	log.Printf("Test")
	return nil
}

func GetDashboard() *model.DashboardPayload {
	log.Printf("Test")
	return nil
}
