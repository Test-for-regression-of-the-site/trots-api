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
