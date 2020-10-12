package service

import (
	"bytes"
	"github.com/Test-for-regression-of-the-site/trots-api/model"
	"github.com/Test-for-regression-of-the-site/trots-api/storage"
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
				completeTask(sessionId, testId, buffer)
				nextChunkIndex := chunkIndex + 1
				if nextChunkIndex < len(chunks) {
					runTasks(sessionId, nextChunkIndex, chunks)
				}
			}

			defer runNextTask()
		}
		go runTask()
	}
}

func completeTask(sessionId string, testId string, report *bytes.Buffer) {
	test := model.TestEntity{
		Id:     testId,
		Report: report.Bytes(),
	}
	storage.PutTest(sessionId, test)
}
