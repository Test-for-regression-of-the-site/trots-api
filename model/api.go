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

type TestResponsePayload struct {
	Id int `json:"id"`
}

type ErrorResponsePayload struct {
	Error   error  `json:"-"`
	Status  int    `json:"-"`
	Message string `json:"message"`
}

type ShortDashboardItemPayload struct {
	Id                string `json:"id"`
	Url               string `json:"url"`
	Performance       int    `json:"performance"`
	BestPractices     int    `json:"bestPractices"`
	Seo               int    `json:"seo"`
	ProgressiveWebApp int    `json:"progressiveWebApp"`
}

type ShortDashboardPayload struct {
	Items []ShortDashboardItemPayload `json:"items"`
}

type DashboardPayload struct {
	Uuid ShortDashboardPayload `json:"UUID"`
}

type ReportsPayload struct {
	ProcessEnd     bool             `json:"processEnd"`
	ShortDashboard DashboardPayload `json:"shortDashboard"`
}
