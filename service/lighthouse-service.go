package service

import (
	"context"
	"fmt"
	"github.com/Test-for-regression-of-the-site/trots-api/configuration"
	"github.com/Test-for-regression-of-the-site/trots-api/constants"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io"
)

var dockerClient = connect()

type LighthouseTaskRequest struct {
	Configuration configuration.LighthouseConfiguration
	Id            string
	Url           string
}

func executeLighthouseTask(request LighthouseTaskRequest, reportWriter io.Writer) (string, error) {
	containerConfig := &container.Config{Image: request.Configuration.Image}
	containerId, dockerError := dockerClient.ContainerCreate(
		context.Background(),
		containerConfig,
		nil,
		nil,
		constants.Lighthouse+constants.Dash+request.Id,
	)
	if dockerError != nil {
		panic(dockerError)
	}
	dockerClient.ContainerWait(context.Background(), containerId.ID, container.WaitConditionNextExit)
	return containerId.ID, nil
}

func connect() *client.Client {
	dockerClient, dockerError := client.NewEnvClient()
	if dockerError != nil {
		fmt.Println("Unable to create docker dockerClient")
		panic(dockerError)
	}
	return dockerClient
}
