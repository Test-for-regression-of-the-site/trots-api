package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionEntity struct {
	Id    primitive.ObjectID `bson:"_id,omitempty"`
	Tests []TestEntity       `bson:"tests,omitempty"`
}

type TestEntity struct {
	Id                string            `bson:"id,omitempty"`
	Url               string            `bson:"url,omitempty"`
	ReportInformation ReportInformation `bson:"reportInformation,omitempty"`
}

type ReportInformation struct {
	Id                string  `bson:"id,omitempty"`
	Performance       float32 `bson:"performance,omitempty"`
	Accessibility     float32 `bson:"accessibility,omitempty"`
	BestPractices     float32 `bson:"bestPractices,omitempty"`
	Seo               float32 `bson:"seo,omitempty"`
	ProgressiveWebApp float32 `bson:"progressiveWebApp,omitempty"`
}

type ReportEntity struct {
	Id     primitive.ObjectID `bson:"_id,omitempty"`
	Report []byte             `bson:"report,omitempty"`
}
