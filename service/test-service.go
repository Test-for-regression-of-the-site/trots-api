package service

import (
	"encoding/json"
	"github.com/Test-for-regression-of-the-site/trots-api/extensions"
	"github.com/Test-for-regression-of-the-site/trots-api/model"
	"github.com/Test-for-regression-of-the-site/trots-api/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"sync"
	"time"
)

var lock = sync.Mutex{}
var working = false

func RunTest(request model.TestRequestPayload) int64 {
	Lock()
	creationTime := time.Now().Unix()
	sessionId := model.SessionIdentifier{CreationTime: creationTime, Id: primitive.NewObjectID().Hex()}
	go runTasks(sessionId, request.TestType, 0, extensions.Chunks(request.Links, request.Parallel))
	return creationTime
}

func GetTestReport(sessionId, testId string) map[string]interface{} {
	test, storageError := storage.GetTest(sessionId, testId)
	if storageError != nil {
		log.Printf("Storage error: %s", storageError)
		return nil
	}
	if test == nil {
		return nil
	}
	reportData, storageError := storage.GetReport(test.ReportInformation.Id)
	if storageError != nil {
		log.Printf("Storage error: %s", storageError)
		return nil
	}
	if reportData == nil {
		return nil
	}
	var report map[string]interface{}
	if storageError = json.Unmarshal(reportData.Report, &report); storageError != nil {
		log.Printf("Storage error: %s", storageError)
		return nil
	}
	return report
}

func GetDashboard() *model.DashboardResponsePayload {
	sessions, storageError := storage.GetSessions()
	if storageError != nil {
		log.Printf("Storage error: %s", storageError)
		return nil
	}
	if sessions == nil {
		return nil
	}
	sessionDashboards := make(map[string]model.SessionReportPayload)
	for _, session := range *sessions {
		var testReports []model.TestReportPayload
		for _, test := range session.Tests {
			testReport := model.TestReportPayload{
				Id:                test.Id,
				Url:               test.Url,
				Accessibility:     test.ReportInformation.Accessibility,
				BestPractices:     test.ReportInformation.BestPractices,
				Performance:       test.ReportInformation.Performance,
				Seo:               test.ReportInformation.Seo,
				ProgressiveWebApp: test.ReportInformation.ProgressiveWebApp,
			}
			testReports = append(testReports, testReport)
		}
		sessionDashboards[session.Id.Hex()] = model.SessionReportPayload{CreationTime: session.CreationTime, Tests: testReports}
	}
	return &model.DashboardResponsePayload{ProcessEnd: !working, ShortDashboard: sessionDashboards}
}

func Lock() {
	lock.Lock()
	working = true
}

func Unlock() {
	working = false
	lock.Unlock()
}
