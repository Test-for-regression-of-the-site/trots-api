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
	ProcessEnd     bool                   `json:"processEnd"`
	ShortDashboard []SessionReportPayload `json:"shortDashboard"`
}

type TestReportPayload struct {
	Id                string `json:"Id"`
	Url               string `json:"url"`
	Performance       int    `json:"performance"`
	Accessibility     int    `json:"accessibility"`
	BestPractices     int    `json:"bestPractices"`
	Seo               int    `json:"seo"`
	ProgressiveWebApp int    `json:"progressiveWebApp"`
}

type SessionReportPayload struct {
	TestReports map[string][]TestReportPayload `json:""`
}
