package service

import (
	"encoding/json"
	"github.com/Test-for-regression-of-the-site/trots-api/extensions"
	"github.com/Test-for-regression-of-the-site/trots-api/model"
	"github.com/Test-for-regression-of-the-site/trots-api/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func RunTest(request model.TestRequestPayload) {
	runTasks(primitive.NewObjectID().Hex(), 0, extensions.Chunks(request.Links, request.Parallel))
}

func GetTestReport(sessionId string, testId string) *map[string]interface{} {
	test, storageError := storage.GetTest(sessionId, testId)
	if storageError != nil {
		log.Printf("Storage error: %s", storageError)
		return nil
	}
	if test == nil {
		return nil
	}
	var report map[string]interface{}
	if storageError = json.Unmarshal(test.Report, &report); storageError != nil {
		log.Printf("Storage error: %s", storageError)
		return nil
	}
	return &report
}

func GetDashboard() *model.DashboardResponsePayload {
	log.Printf("Test")
	return nil
}
