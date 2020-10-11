package service

import (
	"github.com/Test-for-regression-of-the-site/trots-api/extensions"
	"github.com/Test-for-regression-of-the-site/trots-api/model"
	"github.com/google/uuid"
	"log"
)

func RunTest(request model.TestRequestPayload) {
	runTasks(uuid.New().String(), 0, extensions.Chunks(request.Links, request.Parallel))
}

func GetTestReport(sessionId string, testId string) *interface{} {
	log.Printf("Test")
	return nil
}

func GetDashboard() *model.DashboardResponsePayload {
	log.Printf("Test")
	return nil
}
