package service

import (
	"fmt"
	"github.com/Test-for-regression-of-the-site/trots-api/configuration"
	"github.com/Test-for-regression-of-the-site/trots-api/constants"
	docker "github.com/fsouza/go-dockerclient"
	"io"
	"log"
	"os"
	"path/filepath"
)

var dockerClient = connect()

type LighthouseTaskRequest struct {
	Configuration configuration.LighthouseConfiguration
	Id            string
	Url           string
}

func executeLighthouseTask(request LighthouseTaskRequest, reportWriter io.Writer) (string, error) {
	directory := "reports/" + request.Id
	if directoryError := os.Mkdir(directory, os.ModePerm); directoryError != nil {
		log.Printf("Directory error: %s", directoryError)
		return "", directoryError
	}
	directoryPath, directoryError := filepath.Abs(directory)
	if directoryError != nil {
		log.Printf("Directory error: %s", directoryError)
		return "", directoryError
	}
	imageOptions := docker.PullImageOptions{
		Tag:          "latest",
		Repository:   request.Configuration.Image,
		OutputStream: log.Writer(),
	}
	dockerError := dockerClient.PullImage(imageOptions, docker.AuthConfiguration{})
	if dockerError != nil {
		log.Printf("Docker error: %s", dockerError)
		return "", dockerError
	}
	containerConfig := &docker.Config{
		Image:        request.Configuration.Image + ":" + "latest",
		Tty:          true,
		AttachStdout: true,
		AttachStderr: true,
		Cmd: []string{
			"lighthouse",
			"--chrome-flags=\"--headless --disable-gpu\"",
			"--output", "json",
			"--output-path", "/home/chrome/reports/report.json",
			"https://justinribeiro.com",
		},
	}
	hostConfig := &docker.HostConfig{
		Binds:      []string{directoryPath + ":/home/chrome/reports:rw"},
		CapAdd:     []string{"SYS_ADMIN"},
		AutoRemove: true,
	}
	containerOptions := docker.CreateContainerOptions{
		HostConfig: hostConfig,
		Config:     containerConfig,
		Name:       constants.Lighthouse + constants.Dash + request.Id,
	}
	containerId, dockerError := dockerClient.CreateContainer(containerOptions)
	if dockerError != nil {
		log.Printf("Docker error: %s", dockerError)
		return "", dockerError
	}
	if dockerError = dockerClient.StartContainer(containerId.ID, hostConfig); dockerError != nil {
		log.Printf("Docker error: %s", dockerError)
		return "", dockerError
	}
	if _, dockerError = dockerClient.WaitContainer(containerId.ID); dockerError != nil {
		log.Printf("Docker error: %s", dockerError)
		return "", dockerError
	}
	return containerId.ID, nil
}

func connect() *docker.Client {
	dockerClient, dockerError := docker.NewClientFromEnv()
	if dockerError != nil {
		fmt.Println("Unable to create docker dockerClient")
		panic(dockerError)
	}
	return dockerClient
}
