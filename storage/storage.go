package storage

import (
	"bytes"
	"context"
	"github.com/Test-for-regression-of-the-site/trots-api/configuration"
	"github.com/Test-for-regression-of-the-site/trots-api/provider"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoStorage struct {
	client  *mongo.Client
	context *context.Context
}

var storage = connect(provider.Configuration.Mongo)

func PutReport(sessionId string, testId string, report *bytes.Buffer) {
	collection := storage.client.Database("testing").Collection("numbers")
	document := bson.M{"sessionId": sessionId, "testId": testId, "report": report.Bytes()}
	if _, mongoError := collection.InsertOne(*storage.context, document); mongoError != nil {
		log.Printf("Mongo error: %s", mongoError)
		return
	}
}

func connect(configuration configuration.MongoConfiguration) *MongoStorage {
	client, mongoError := mongo.NewClient(options.Client().ApplyURI(configuration.Address))
	if mongoError != nil {
		log.Panicf("Mongo error: %s", mongoError)
	}
	mongoContext, cancel := context.WithTimeout(context.Background(), configuration.Timeout)
	defer cancel()
	if mongoError = client.Connect(mongoContext); mongoError != nil {
		log.Panicf("Mongo error: %s", mongoError)
	}
	storage := &MongoStorage{client: client, context: &mongoContext}
	defer storage.disconnect()
	log.Printf("Mongo client connected to: %s", configuration.Address)
	return storage
}

func (storage *MongoStorage) disconnect() {
	if disconnectError := storage.client.Disconnect(*storage.context); disconnectError != nil {
		log.Printf("Mongo error: %s", disconnectError)
	}
}
