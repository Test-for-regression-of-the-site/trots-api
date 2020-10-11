package lighthouse

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os/exec"
	"strings"
	"sync"
)

type ExecutionConfiguration struct {
	Image       string
	Arguments   []string
	Environment []string
}

type Task struct {
	sync.RWMutex
	Done         chan struct{}
	Running      bool
	Error        error
	Url          string
	ReportBuffer *bytes.Buffer
}

func ExecuteLighthouseTask(configuration ExecutionConfiguration, link string, reportWriter io.Writer) error {
	if configuration.Image == "" {
		configuration.Image = "lighthouse"
	}
	var _, urlError = url.Parse(link)
	if urlError != nil {
		return fmt.Errorf("URL error: %w", urlError)
	}
	var stdError = &bytes.Buffer{}
	var stdOut = &bytes.Buffer{}
	var command = &exec.Cmd{
		Path: configuration.Image,
		Args: configuration.formatArguments(
			configuration.Image,
			link,
			"--output", "json",
		),
		Stdout: stdOut,
		Stderr: stdError,
		Env:    configuration.Environment,
	}
	_, _ = fmt.Fprintf(stdError, "%s\n\n", command)
	if commandRunError := command.Run(); commandRunError != nil {
		return fmt.Errorf("Running lighthouse container error: %w\n%s", commandRunError, stdError)
	}
	var _, stdOutWritingError = stdOut.WriteTo(reportWriter)
	if stdOutWritingError != nil {
		return fmt.Errorf("STD out writing error: %w: %s", stdOutWritingError, parseErrorMessage(stdError.Bytes()))
	}
	if validationError := validateJson(stdOut.Bytes()); validationError != nil {
		return fmt.Errorf("Validation error: %w", validationError)
	}
	return nil
}

func (configuration ExecutionConfiguration) formatArguments(extra ...string) []string {
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
