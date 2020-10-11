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
		var task = &lighthouse.Task{
			Done:         make(chan struct{}),
			Url:          url,
			ReportBuffer: &bytes.Buffer{},
		}

		go func() {
			defer close(task.Done)

			task.Lock()
			defer task.Unlock()

			task.Running = true

			defer func() {
				task.Running = false
			}()

			configuration := lighthouse.ExecutionConfiguration{
				Image:       "",
				Arguments:   nil,
				Environment: nil,
			}

			task.Error = lighthouse.ExecuteLighthouseTask(configuration, task.Url, task.ReportBuffer)
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
