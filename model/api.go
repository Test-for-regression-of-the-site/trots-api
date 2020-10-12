package model

type TestRequestPayload struct {
	TimeCreate int      `json:"timeCreate"`
	Links      []string `json:"links"`
	Parallel   int      `json:"parallel"`
	TestType   string   `json:"testType"`
}

type ErrorResponsePayload struct {
	Error   error  `json:"-"`
	Status  int    `json:"-"`
	Message string `json:"message"`
}

type DashboardResponsePayload struct {
	ProcessEnd     bool                            `json:"processEnd"`
	ShortDashboard map[string]SessionReportPayload `json:"shortDashboard"`
}

type TestReportPayload struct {
	Id                string  `json:"id"`
	Url               string  `json:"url"`
	Performance       float32 `json:"performance"`
	Accessibility     float32 `json:"accessibility"`
	BestPractices     float32 `json:"bestPractices"`
	Seo               float32 `json:"seo"`
	ProgressiveWebApp float32 `json:"progressiveWebApp"`
}
type SessionReportPayload struct {
	CreationTime int64               `json:"creationTime"`
	Tests        []TestReportPayload `json:"tests"`
}

type TestResponsePayload struct {
	CreationTime int64 `json:"creationTime"`
}
