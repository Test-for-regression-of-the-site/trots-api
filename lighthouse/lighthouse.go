// Package lighthouse provides lighthouse runners.
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
)

// Config describes lighthouse runner settings.
type Config struct {
	Exec string
	Args []string
	Env  []string
}

// Run lighthouse task.
func (cfg Config) Run(page string, report io.Writer) error {
	if cfg.Exec == "" {
		cfg.Exec = "lighthouse"
	}
	var _, errURL = url.Parse(page)
	if errURL != nil {
		return fmt.Errorf("page URL: %w", errURL)
	}
	var stderr = &bytes.Buffer{}
	var stdout = &bytes.Buffer{}
	var cmd = &exec.Cmd{
		Path: cfg.Exec,
		Args: cfg.args(
			cfg.Exec,
			page,
			"--output", "json",
		),
		Stdout: stdout,
		Stderr: stderr,
		Env:    cfg.Env,
	}
	_, _ = fmt.Fprintf(stderr, "%s\n\n", cmd)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("running lighthouse tool: %w\n%s", err, stderr)
	}
	var _, errCopy = stdout.WriteTo(report)
	if errCopy != nil {
		var msg = parseErrorMessage(stderr.Bytes())
		return fmt.Errorf("writing report: %w: %s", errCopy, msg)
	}
	if err := validateJSON(stdout.Bytes()); err != nil {
		return fmt.Errorf("report: %w", err)
	}
	return nil
}

func (cfg Config) args(extra ...string) []string {
	var n = len(cfg.Args)
	return append(cfg.Args[:n:n], extra...)
}

func parseErrorMessage(data []byte) string {
	var sep = []byte(" ")
	var msg = &strings.Builder{}
	var re = bytes.NewReader(data)
	var sc = bufio.NewScanner(re)
	for sc.Scan() {
		var line = sc.Bytes()
		if bytes.HasPrefix(line, sep) {
			continue
		}
		_, _ = msg.Write(line)
		_, _ = msg.WriteString(" ")
	}
	if msg.Len() == 0 {
		return string(data)
	}
	return msg.String()
}

func validateJSON(data []byte) error {
	var re = bytes.NewReader(data)
	var decoder = json.NewDecoder(re)
	for decoder.More() {
		var _, err = decoder.Token()
		if err != nil {
			return err
		}
	}
	return nil
}
