package service

import (
	"bytes"
	"github.com/Test-for-regression-of-the-site/trots-api/storage"
	"github.com/google/uuid"
	"log"
)

func runTasks(sessionId string, chunkIndex int, chunks [][]string) {
	for _, url := range chunks[chunkIndex] {
		runTask := func() {
			testId := uuid.New()
			buffer := &bytes.Buffer{}
			request := LighthouseTaskRequest{
				SessionId: sessionId,
				TestId:    testId.String(),
				Url:       url,
			}
			lighthouseError := executeLighthouseTask(request, buffer)
			if lighthouseError != nil {
				log.Printf("Lighthouse error: %s", lighthouseError.Error())
				return
			}

			runNextTask := func() {
				completeTask(sessionId, testId.String(), buffer)
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
	storage.PutReport(sessionId, testId, report)
}
