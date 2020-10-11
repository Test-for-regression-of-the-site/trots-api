package configuration

type LighthouseConfiguration struct {
	Image string
}

type Configuration struct {
	Address    string
	Lighthouse LighthouseConfiguration
}
