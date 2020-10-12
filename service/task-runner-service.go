package service

import (
	"bytes"
	"encoding/json"
	"github.com/Test-for-regression-of-the-site/trots-api/model"
	"github.com/Test-for-regression-of-the-site/trots-api/storage"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func runTasks(sessionId string, chunkIndex int, chunks [][]string) {
	for _, url := range chunks[chunkIndex] {
		runTask := func() {
			testId := primitive.NewObjectID().Hex()
			buffer := &bytes.Buffer{}
			request := LighthouseTaskRequest{
				SessionId: sessionId,
				TestId:    testId,
				Url:       url,
			}
			lighthouseError := executeLighthouseTask(request, buffer)
			if lighthouseError != nil {
				log.Printf("Lighthouse error: %s", lighthouseError.Error())
				return
			}

			runNextTask := func() {
				completeTask(sessionId, testId, url, buffer)
				nextChunkIndex := chunkIndex + 1
				if nextChunkIndex >= len(chunks) {
					Unlock()
					return
				}
				runTasks(sessionId, nextChunkIndex, chunks)
			}

			defer runNextTask()
		}
		go runTask()
	}
}

func completeTask(sessionId string, testId string, url string, reportContent *bytes.Buffer) {
	reportId := primitive.NewObjectID()
	var report map[string]interface{}
	if jsonError := json.Unmarshal(reportContent.Bytes(), &report); jsonError != nil {
		log.Printf("Storage error: %s", jsonError)
		return
	}
	categories := cast.ToStringMap(report["categories"])
	testEntity := model.TestEntity{
		Id:  testId,
		Url: url,
		ReportInformation: model.ReportInformation{
			Id:                reportId.Hex(),
			Accessibility:     cast.ToFloat32(cast.ToStringMap(categories["accessibility"])["score"]),
			Performance:       cast.ToFloat32(cast.ToStringMap(categories["performance"])["score"]),
			BestPractices:     cast.ToFloat32(cast.ToStringMap(categories["best-practices"])["score"]),
			Seo:               cast.ToFloat32(cast.ToStringMap(categories["seo"])["score"]),
			ProgressiveWebApp: cast.ToFloat32(cast.ToStringMap(categories["pwa"])["score"]),
		},
	}
	reportEntity := &model.ReportEntity{
		Id:     reportId,
		Report: reportContent.Bytes(),
	}
	storage.PutTest(sessionId, testEntity, reportEntity)
}
