package storage

import (
	"context"
	"github.com/Test-for-regression-of-the-site/trots-api/configuration"
	"github.com/Test-for-regression-of-the-site/trots-api/constants"
	"github.com/Test-for-regression-of-the-site/trots-api/model"
	"github.com/Test-for-regression-of-the-site/trots-api/provider"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoStorage struct {
	client *mongo.Client
}

var storage = connect(provider.Configuration.Mongo)

func PutTest(sessionId string, test model.TestEntity) {
	collection := storage.client.Database(constants.Trots).Collection(constants.Session)
	mongoContext, cancel := context.WithTimeout(context.Background(), provider.Configuration.Mongo.Timeout)
	defer cancel()
	session, mongoError := GetSession(sessionId)
	if mongoError != nil {
		log.Printf("Mongo error: %s", mongoError)
		return
	}
	if session == nil {
		id, mongoError := primitive.ObjectIDFromHex(sessionId)
		if mongoError != nil {
			log.Printf("Mongo error: %s", mongoError)
			return
		}
		session = &model.SessionEntity{
			Id:    id,
			Tests: []model.TestEntity{test},
		}
	}
	session.Tests = append(session.Tests, test)
	if _, mongoError := collection.InsertOne(mongoContext, session); mongoError != nil {
		log.Printf("Mongo error: %s", mongoError)
	}
}

func GetSession(sessionId string) (*model.SessionEntity, error) {
	collection := storage.client.Database(constants.Trots).Collection(constants.Session)
	mongoContext, cancel := context.WithTimeout(context.Background(), provider.Configuration.Mongo.Timeout)
	defer cancel()
	cursor, mongoError := collection.Find(mongoContext, bson.D{{"_id", sessionId}})
	if mongoError != nil {
		log.Printf("Mongo error: %s", mongoError)
		return nil, mongoError
	}
	if !cursor.TryNext(mongoContext) {
		return nil, nil
	}
	var session model.SessionEntity
	if mongoError := cursor.Decode(session); mongoError != nil {
		log.Printf("Mongo error: %s", mongoError)
		return nil, mongoError
	}
	return &session, nil
}

func GetTest(sessionId string, testId string) (*model.TestEntity, error) {
	session, mongoError := GetSession(sessionId)
	if mongoError != nil {
		log.Printf("Mongo error: %s", mongoError)
		return nil, mongoError
	}
	for _, test := range session.Tests {
		if test.Id == testId {
			return &test, nil
		}
	}
	return nil, nil
}

func GetSessions() (*[]model.SessionEntity, error) {
	collection := storage.client.Database(constants.Trots).Collection(constants.Session)
	mongoContext, cancel := context.WithTimeout(context.Background(), provider.Configuration.Mongo.Timeout)
	defer cancel()
	cursor, mongoError := collection.Find(mongoContext, bson.D{})
	if mongoError != nil {
		log.Printf("Mongo error: %s", mongoError)
		return nil, mongoError
	}
	var sessions []model.SessionEntity
	for cursor.Next(mongoContext) {
		var session model.SessionEntity
		if mongoError = cursor.Decode(session); mongoError != nil {
			log.Printf("Mongo error: %s", mongoError)
			return nil, mongoError
		}
		sessions = append(sessions, session)
	}
	return &sessions, nil
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
	log.Printf("Mongo client connected to: %s", configuration.Address)
	return &MongoStorage{client: client}
}

func Disconnect() {
	if disconnectError := storage.client.Disconnect(context.Background()); disconnectError != nil {
		log.Printf("Mongo error: %s", disconnectError)
	}
}
