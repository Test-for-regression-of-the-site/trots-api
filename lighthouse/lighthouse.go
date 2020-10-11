package lighthouse

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Test-for-regression-of-the-site/trots-api/configuration"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"io"
	"net/url"
	"os/exec"
	"strings"
)

func ExecuteLighthouseTask(configuration configuration.LighthouseExecutionConfiguration, link string, reportWriter io.Writer) (string, error) {
	if configuration.Image == "" {
		configuration.Image = "lighthouse"
	}

	dockerClient, dockerError := client.NewEnvClient()
	if dockerError != nil {
		fmt.Println("Unable to create docker dockerClient")
		panic(dockerError)
	}

	hostBinding := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: "8000",
	}
	containerPort, dockerError := nat.NewPort("tcp", "80")
	if dockerError != nil {
		panic("Unable to get the port")
	}

	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}

	containerConfig := &container.Config{Image: configuration.Image}
	hostConfig := &container.HostConfig{PortBindings: portBinding}
	containerId, dockerError := dockerClient.ContainerCreate(context.Background(), containerConfig, hostConfig, nil, "lighthouse")
	if dockerError != nil {
		panic(dockerError)
	}

	dockerClient.ContainerWait(context.Background(), containerId.ID)

	var _, urlError = url.Parse(link)
	if urlError != nil {
		return "", fmt.Errorf("URL error: %w", urlError)
	}
	var stdError = &bytes.Buffer{}
	var stdOut = &bytes.Buffer{}
	var command = &exec.Cmd{
		Path:   configuration.Image,
		Args:   formatArguments(configuration, configuration.Image, link, "--output", "json"),
		Stdout: stdOut,
		Stderr: stdError,
		Env:    configuration.Environment,
	}
	_, _ = fmt.Fprintf(stdError, "%s\n\n", command)
	if commandRunError := command.Run(); commandRunError != nil {
		return "", fmt.Errorf("Running lighthouse container error: %w\n%s", commandRunError, stdError)
	}
	var _, stdOutWritingError = stdOut.WriteTo(reportWriter)
	if stdOutWritingError != nil {
		return "", fmt.Errorf("STD out writing error: %w: %s", stdOutWritingError, parseErrorMessage(stdError.Bytes()))
	}
	if validationError := validateJson(stdOut.Bytes()); validationError != nil {
		return "", fmt.Errorf("Validation error: %w", validationError)
	}

	return containerId.ID, nil
}

func formatArguments(configuration configuration.LighthouseExecutionConfiguration, extra ...string) []string {
	var length = len(configuration.Arguments)
	return append(configuration.Arguments[:length:length], extra...)
}

func parseErrorMessage(data []byte) string {
	var message = &strings.Builder{}
	var reader = bytes.NewReader(data)
	var scanner = bufio.NewScanner(reader)
	for scanner.Scan() {
		var line = scanner.Bytes()
		if bytes.HasPrefix(line, []byte(" ")) {
			continue
		}
		_, _ = message.Write(line)
		_, _ = message.WriteString(" ")
	}
	if message.Len() == 0 {
		return string(data)
	}
	return message.String()
}

func validateJson(data []byte) error {
	var reader = bytes.NewReader(data)
	var decoder = json.NewDecoder(reader)
	for decoder.More() {
		var _, err = decoder.Token()
		if err != nil {
			return err
		}
	}
	return nil
}
