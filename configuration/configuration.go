package configuration

type LighthouseExecutionConfiguration struct {
	Image       string
	Arguments   []string
	Environment []string
}

type Configuration struct {
	Address    string
	Lighthouse LighthouseExecutionConfiguration
}
