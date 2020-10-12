package configuration

import "time"

type LighthouseConfiguration struct {
	Image             string
	Tag               string
	ReportsTargetPath string
	ReportsSourcePath string
}

type MongoConfiguration struct {
	Address string
	Timeout time.Duration
}

type Configuration struct {
	Address    string
	Timeout    time.Duration
	Lighthouse LighthouseConfiguration
	Mongo      MongoConfiguration
}
