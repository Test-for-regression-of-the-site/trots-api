package service

import (
	"github.com/Test-for-regression-of-the-site/trots-api/extensions"
	"github.com/Test-for-regression-of-the-site/trots-api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func RunTest(request model.TestRequestPayload) {
	runTasks(primitive.NewObjectID().Hex(), 0, extensions.Chunks(request.Links, request.Parallel))
}

func GetTestReport(sessionId string, testId string) *interface{} {
	log.Printf("Test")
	return nil
}

func GetDashboard() *model.DashboardResponsePayload {
	log.Printf("Test")
	return nil
}
