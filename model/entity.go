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
	ReportInformation ReportInformation `bson:"reportInformation,omitempty"`
}

type ReportInformation struct {
	Id                string `bson:"id,omitempty"`
	Performance       int    `bson:"performance,omitempty"`
	Accessibility     int    `bson:"accessibility,omitempty"`
	BestPractices     int    `bson:"bestPractices,omitempty"`
	Seo               int    `bson:"seo,omitempty"`
	ProgressiveWebApp int    `bson:"progressiveWebApp,omitempty"`
}

type ReportEntity struct {
	Id     primitive.ObjectID `bson:"_id,omitempty"`
	Report []byte             `bson:"report,omitempty"`
}
