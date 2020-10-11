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

	for _, url := range request.Links {
		var task = &lighthouse.Task{Url: url, ReportBuffer: &bytes.Buffer{}}
		go func() {
			configuration := lighthouse.ExecutionConfiguration{
				Image:       "",
				Arguments:   nil,
				Environment: nil,
			}
			task.Error = lighthouse.ExecuteLighthouseTask(configuration, task.Url, task.ReportBuffer)
			defer func() {

			}()
		}()
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
