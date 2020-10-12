package model

type TestType string

const (
	DesktopTest TestType = "desktop"
	MobileTest  TestType = "mobile"
)

type TestRequestPayload struct {
	TimeCreate int      `json:"timeCreate"`
	Links      []string `json:"links"`
	Parallel   int      `json:"parallel"`
	TestType   TestType `json:"testType"`
}

type ErrorResponsePayload struct {
	Error   error  `json:"-"`
	Status  int    `json:"-"`
	Message string `json:"message"`
}

type DashboardResponsePayload struct {
	ProcessEnd     bool                           `json:"processEnd"`
	ShortDashboard map[string][]TestReportPayload `json:"shortDashboard"`
}

type TestReportPayload struct {
	Id                string  `json:"Id"`
	Url               string  `json:"url"`
	Performance       float32 `json:"performance"`
	Accessibility     float32 `json:"accessibility"`
	BestPractices     float32 `json:"bestPractices"`
	Seo               float32 `json:"seo"`
	ProgressiveWebApp float32 `json:"progressiveWebApp"`
}
