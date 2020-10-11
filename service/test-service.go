package service

import (
	model "github.com/Test-for-regression-of-the-site/trots-api/model"
	"log"
)

func RunTest(request model.TestRequestPayload) *model.TestResponsePayload {
	log.Printf("Test")
	return &model.TestResponsePayload{Id: 123}
}
