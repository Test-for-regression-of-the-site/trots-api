package service

import (
	"bytes"
	"github.com/Test-for-regression-of-the-site/trots-api/configuration"
	"github.com/Test-for-regression-of-the-site/trots-api/constants"
	"github.com/Test-for-regression-of-the-site/trots-api/model"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"log"
)

func RunTest(request model.TestRequestPayload) {
	id := uuid.New()
	completeTask := func(containerId string, reportBuffer *bytes.Buffer) {
		// TODO: Put into mongo
	}

	runTask := func(url string) {
		reportBuffer := &bytes.Buffer{}
		request := LighthouseTaskRequest{
			Configuration: configuration.LighthouseConfiguration{Image: viper.GetString(constants.LighthouseImage)},
			Id:            id.String(),
			Url:           url,
		}
		containerId, lighthouseError := executeLighthouseTask(request, reportBuffer)
		if lighthouseError != nil {
			log.Printf("Lighthouse error: %s", lighthouseError.Error())
			return
		}
		defer completeTask(containerId, reportBuffer)
	}

	for _, url := range request.Links {
		go runTask(url)
	}
}

func GetTestReport(sessionId string, testId string) *interface{} {
	log.Printf("Test")
	return nil
}

func GetDashboard() *model.DashboardResponsePayload {
	log.Printf("Test")
	return nil
}
