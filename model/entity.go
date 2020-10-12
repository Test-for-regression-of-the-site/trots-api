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
	Id string `bson:"id,omitempty"`
}

type ReportEntity struct {
	Id     string `bson:"id,omitempty"`
	Report []byte `bson:"report,omitempty"`
}
