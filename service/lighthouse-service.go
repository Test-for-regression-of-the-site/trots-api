package service

import (
	"github.com/Test-for-regression-of-the-site/trots-api/constants"
	"github.com/Test-for-regression-of-the-site/trots-api/model"
	"github.com/Test-for-regression-of-the-site/trots-api/provider"
	docker "github.com/fsouza/go-dockerclient"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var dockerClient = connect()

type LighthouseTaskRequest struct {
	SessionId string
	TestId    string
	Url       string
	TestType  string
	Trotling  model.Trotling
}

func executeLighthouseTask(request LighthouseTaskRequest, reportWriter io.Writer) error {
	reportsTargetPath, directoryError := makeReportsDirectory(request)
	if directoryError != nil {
		return directoryError
	}
	containerError := launchLighthouse(reportsTargetPath, request)
	if containerError != nil {
		return containerError
	}
	time.Sleep(constants.LighthouseReportWaiting)
	return handleReport(reportsTargetPath, reportWriter)
}

func makeReportsDirectory(request LighthouseTaskRequest) (string, error) {
	directory := provider.Configuration.Lighthouse.ReportsTargetPath + constants.Slash + request.SessionId + constants.Slash + request.TestId
	if directoryError := os.MkdirAll(directory, os.ModePerm); directoryError != nil {
		log.Printf("Directory error: %s", directoryError)
		return "", directoryError
	}
	directoryPath, directoryError := filepath.Abs(filepath.FromSlash(directory))
	if directoryError != nil {
		log.Printf("Directory error: %s", directoryError)
		return "", directoryError
	}
	return directoryPath, nil
}

func launchLighthouse(reportsTargetPath string, request LighthouseTaskRequest) error {
	imageOptions := docker.PullImageOptions{
		Tag:          provider.Configuration.Lighthouse.Tag,
		Repository:   provider.Configuration.Lighthouse.Image,
		OutputStream: log.Writer(),
	}
	dockerError := dockerClient.PullImage(imageOptions, docker.AuthConfiguration{})
	if dockerError != nil {
		log.Printf("Docker error: %s", dockerError)
		return dockerError
	}
	options := []string{
		constants.Lighthouse,
		constants.LightHouseFlagChrome,
		constants.LightHouseFlagOutput, constants.LightHouseFlagJson,
		constants.LightHouseEmulatedFormFactor, request.TestType,
		constants.LightHouseFlagOutputPath, reportsTargetPath + constants.Slash + constants.LighthouseReportFile,
	}

	if request.Trotling.Simulate {
		options = append(options, constants.LightHouseTrotlingMethod, request.Trotling.Method)
		options = append(options, constants.LightHouseTrotlingDownloadThroughputKbps, request.Trotling.DownloadKbps)
		options = append(options, constants.LightHouseTrotlingUploadThroughputKbps, request.Trotling.UploadKbps)
		options = append(options, constants.LightHouseTrotlingThroughputKbps, request.Trotling.Kbps)
	}

	options = append(options, request.Url)

	log.Printf("Lighthouse arguments: %v", options)

	containerConfig := &docker.Config{
		Image:        provider.Configuration.Lighthouse.Image + constants.Colon + provider.Configuration.Lighthouse.Tag,
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          options,
	}
	reportsSourcePath := provider.Configuration.Lighthouse.ReportsSourcePath + constants.Slash + request.SessionId + constants.Slash + request.TestId
	hostConfig := &docker.HostConfig{
		Binds:  []string{reportsSourcePath + constants.Colon + reportsTargetPath + constants.Colon + constants.DockerReadWriteMode},
		CapAdd: []string{constants.DockerSysAdminCapability},
	}
	containerOptions := docker.CreateContainerOptions{
		Config:     containerConfig,
		Name:       constants.Lighthouse + constants.Dash + request.TestId,
		HostConfig: hostConfig,
	}
	containerId, dockerError := dockerClient.CreateContainer(containerOptions)
	if dockerError != nil {
		log.Printf("Docker error: %s", dockerError)
		return dockerError
	}
	defer func() {
		dockerClient.RemoveContainer(docker.RemoveContainerOptions{
			ID:    containerId.ID,
			Force: true,
		})
	}()
	if dockerError = dockerClient.StartContainer(containerId.ID, hostConfig); dockerError != nil {
		log.Printf("Docker error: %s", dockerError)
		return dockerError
	}
	if _, dockerError = dockerClient.WaitContainer(containerId.ID); dockerError != nil {
		log.Printf("Docker error: %s", dockerError)
		return dockerError
	}
	logsOptions := docker.LogsOptions{
		Follow:       true,
		ErrorStream:  log.Writer(),
		OutputStream: log.Writer(),
		Stdout:       true,
		Stderr:       true,
		Container:    containerId.ID,
	}
	if dockerError = dockerClient.Logs(logsOptions); dockerError != nil {
		log.Printf("Docker error: %s", dockerError)
		return dockerError
	}
	return nil
}

func handleReport(reportsTargetPath string, reportWriter io.Writer) error {
	file, readingError := ioutil.ReadFile(filepath.FromSlash(reportsTargetPath + constants.Slash + constants.LighthouseReportFile))
	if readingError != nil {
		log.Printf("Reading error: %s", readingError)
		return readingError
	}
	if _, writingError := reportWriter.Write(file); writingError != nil {
		log.Printf("Writing error: %s", writingError)
		return writingError
	}
	if removingError := os.RemoveAll(reportsTargetPath); removingError != nil {
		log.Printf("Removing error: %s", removingError)
		return removingError
	}
	parent, _ := filepath.Split(reportsTargetPath)
	files, removingError := ioutil.ReadDir(parent)
	if removingError != nil {
		log.Printf("Removing error: %s", removingError)
		return removingError
	}
	if len(files) == 0 {
		if removingError := os.RemoveAll(parent); removingError != nil {
			log.Printf("Removing error: %s", removingError)
			return removingError
		}
	}
	return nil
}

func connect() *docker.Client {
	dockerClient, dockerError := docker.NewClientFromEnv()
	if dockerError != nil {
		log.Printf("Unable to create docker dockerClient")
		panic(dockerError)
	}
	return dockerClient
}
