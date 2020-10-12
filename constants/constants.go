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
)

const (
	Trots      = "trots"
	Yml        = "yaml"
	Lighthouse = "lighthouse"
	Session    = "session"
	Report     = "report"
	Tests      = "tests"

	ServerAddressKey               = "trots.server.address"
	TimeoutKey                     = "trots.server.timeout"
	LighthouseImageKey             = "trots.lighthouse.image"
	LighthouseTagKey               = "trots.lighthouse.tag"
	LighthouseReportsTargetPathKey = "trots.lighthouse.reportsTargetPath"
	LighthouseReportsSourcePathKey = "trots.lighthouse.reportsSourcePath"
	MongoAddressKey                = "trots.mongo.address"
	MongoTimeoutKey                = "trots.mongo.timeout"

	MongoId  = "_id"
	MongoSet = "$set"

	Dot   = "."
	Dash  = "-"
	Slash = "/"
	Colon = ":"

	DockerReadWriteMode      = "rw"
	DockerSysAdminCapability = "SYS_ADMIN"

	SessionIdParameter = "sessionId"
	TestIdParameter    = "testId"

	TasksRoutePattern          = "/tasks"
	SessionIdParameterPattern  = "/{sessionId}"
	TestIdParameterPattern     = "/{testId}"
	TasksDashboardRoutePattern = "/tasks/dashboard"

	LighthouseReportFile    = "report.json"
	LighthouseReportWaiting = 10 * time.Second

	LightHouseFlagChrome         = "--chrome-flags=\"--headless --no-sandbox --disable-gpu\""
	LightHouseFlagOutput         = "--output"
	LightHouseFlagOutputPath     = "--output-path"
	LightHouseFlagJson           = "json"
	LightHouseEmulatedFormFactor = "--emulated-form-factor"
)
