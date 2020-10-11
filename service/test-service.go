package service

import (
	"bytes"
	"github.com/Test-for-regression-of-the-site/trots-api/lighthouse"
	model "github.com/Test-for-regression-of-the-site/trots-api/model"
	"github.com/google/uuid"
	"log"
)

func RunTest(request model.TestRequestPayload) *model.TestResponsePayload {
	var id = uuid.New().String()

	saveTask := func(reportBuffer *bytes.Buffer) {
		// TODO: Put into mongo
	}

	runTask := func(url string) {
		configuration := lighthouse.ExecutionConfiguration{
			Image:       "",
			Arguments:   nil,
			Environment: nil,
		}
		reportBuffer := &bytes.Buffer{}
		lighthouseError := lighthouse.ExecuteLighthouseTask(configuration, url, reportBuffer)
		if lighthouseError != nil {
			return
		}
		defer saveTask(reportBuffer)
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
