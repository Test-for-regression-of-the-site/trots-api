package constants

import (
	"github.com/go-chi/cors"
	"time"
)

var (
	CorsOptions = cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}

	LighthouseOptions = []string{
		"lighthouse",
		"--chrome-flags=\"--headless --disable-gpu\"",
		"--output", "json",
		"--output-path", LighthouseReportsDirectory + Slash + LighthouseReportFile,
	}
)

const (
	Trots      = "trots"
	Yml        = "yaml"
	Lighthouse = "lighthouse"

	ServerAddressKey = "trots.server.address"
	TimeoutKey       = "trots.server.timeout"
	LighthouseImage  = "trots.lighthouse.image"
	LighthouseTag    = "trots.lighthouse.tag"
	MongoAddress     = "trots.mongo.address"
	MongoTimeout     = "trots.mongo.timeout"

	Dot    = "."
	Dash   = "-"
	Slash  = "/"
	Colon  = ":"
	Latest = "latest"

	DockerReadWriteMode      = "rw"
	DockerSysAdminCapability = "SYS_ADMIN"

	SessionIdParameter = "sessionId"
	TestIdParameter    = "testId"

	TasksRoutePattern          = "/tasks"
	SessionIdParameterPattern  = "/{sessionId}"
	TestIdParameterPattern     = "/{testId}"
	TasksDashboardRoutePattern = "/tasks/dashboard"

	LighthouseReportsDirectory    = "/home/chrome/reports"
	LighthouseReportFile          = "report.json"
	LighthouseHostReportDirectory = "reports"
	LighthouseReportWaiting       = 10 * time.Second
)
