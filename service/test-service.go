package service

import (
	model "github.com/Test-for-regression-of-the-site/trots-api/model"
	"log"
)

func RunTest(request model.TestRequestPayload) *model.TestResponsePayload {
	log.Printf("Test")
	return &model.TestResponsePayload{Id: 123}
}

func GetTestReports(sessionId string, testId string) *model.ReportsPayload {
	log.Printf("Test")
	return nil
}

func GetDashboard() *model.DashboardPayload {
	log.Printf("Test")
	return nil
}
