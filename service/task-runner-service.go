package service

import (
	"bytes"
	"encoding/json"
	"github.com/Test-for-regression-of-the-site/trots-api/model"
	"github.com/Test-for-regression-of-the-site/trots-api/storage"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"sync"
)

type TaskRunnerRequest struct {
	SessionId  model.SessionIdentifier
	TestType   string
	ChunkIndex int
	Chunks     [][]string
	Trotling   model.Trotling
}

func runTasks(request *TaskRunnerRequest) {
	urls := request.Chunks[request.ChunkIndex]
	var waitGroup sync.WaitGroup

	runTask := func(group *sync.WaitGroup, url string) {
		defer group.Done()
		testId := primitive.NewObjectID().Hex()
		buffer := &bytes.Buffer{}
		lighthouseTaskRequest := LighthouseTaskRequest{
			SessionId: request.SessionId.Id,
			TestId:    testId,
			Url:       url,
			TestType:  request.TestType,
			Trotling:  request.Trotling,
		}
		lighthouseError := executeLighthouseTask(lighthouseTaskRequest, buffer)
		if lighthouseError != nil {
			log.Printf("Lighthouse error: %s", lighthouseError.Error())
			return
		}
		completeTask(request.SessionId, testId, url, buffer)
	}

	for _, url := range urls {
		waitGroup.Add(1)
		go runTask(&waitGroup, url)
	}

	waitGroup.Wait()
	nextIndex := request.ChunkIndex + 1
	if nextIndex == len(request.Chunks) {
		Unlock()
		return
	}
	runTasks(&TaskRunnerRequest{
		SessionId:  request.SessionId,
		TestType:   request.TestType,
		ChunkIndex: nextIndex,
		Chunks:     request.Chunks,
		Trotling:   request.Trotling,
	})
}

func completeTask(sessionId model.SessionIdentifier, testId, url string, reportContent *bytes.Buffer) {
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
