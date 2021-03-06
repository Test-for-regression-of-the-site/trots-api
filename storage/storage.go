package storage

import (
	"context"
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

var storage = connect()

func PutTest(sessionId model.SessionIdentifier, test model.TestEntity, report *model.ReportEntity) {
	mongoContext, cancel := context.WithTimeout(context.Background(), provider.Configuration.Mongo.Timeout)
	defer cancel()
	id, mongoError := primitive.ObjectIDFromHex(sessionId.Id)
	if mongoError != nil {
		log.Printf("Mongo error: %s", mongoError)
		return
	}
	mongoError = storage.client.UseSession(mongoContext, func(mongoSession mongo.SessionContext) error {
		if mongoError := mongoSession.StartTransaction(); mongoError != nil {
			log.Printf("Mongo error: %s", mongoError)
			return mongoError
		}
		sessionCollection := storage.client.Database(constants.Trots).Collection(constants.Session)
		reportCollection := storage.client.Database(constants.Trots).Collection(constants.Report)
		if _, mongoError := reportCollection.InsertOne(mongoContext, report); mongoError != nil {
			log.Printf("Mongo error: %s", mongoError)
			return mongoError
		}
		session, mongoError := GetSession(sessionId.Id)
		if mongoError != nil {
			log.Printf("Mongo error: %s", mongoError)
			return mongoError
		}
		if session == nil {
			session = &model.SessionEntity{
				Id:           id,
				CreationTime: sessionId.CreationTime,
				Tests:        []model.TestEntity{test},
			}
			if _, mongoError := sessionCollection.InsertOne(mongoContext, session); mongoError != nil {
				log.Printf("Mongo error: %s", mongoError)
			}
			return mongoError
		}
		session.Tests = append(session.Tests, test)
		filter := bson.D{{constants.MongoId, id}}
		update := bson.D{{constants.MongoSet, bson.D{{constants.Tests, session.Tests}}}}
		if _, mongoError := sessionCollection.UpdateOne(mongoContext, filter, update); mongoError != nil {
			log.Printf("Mongo error: %s", mongoError)
			return mongoError
		}
		if mongoError := mongoSession.CommitTransaction(mongoContext); mongoError != nil {
			log.Printf("Mongo error: %s", mongoError)
			return mongoError
		}
		return nil
	})
	if mongoError != nil {
		log.Printf("Mongo error: %s", mongoError)
	}
}

func GetSession(sessionId string) (*model.SessionEntity, error) {
	id, mongoError := primitive.ObjectIDFromHex(sessionId)
	if mongoError != nil {
		log.Printf("Mongo error: %s", mongoError)
		return nil, mongoError
	}
	collection := storage.client.Database(constants.Trots).Collection(constants.Session)
	mongoContext, cancel := context.WithTimeout(context.Background(), provider.Configuration.Mongo.Timeout)
	defer cancel()
	cursor, mongoError := collection.Find(mongoContext, bson.D{{constants.MongoId, id}})
	if mongoError != nil {
		log.Printf("Mongo error: %s", mongoError)
		return nil, mongoError
	}
	if !cursor.Next(mongoContext) {
		return nil, nil
	}
	var session model.SessionEntity
	if mongoError := cursor.Decode(&session); mongoError != nil {
		log.Printf("Mongo error: %s", mongoError)
		return nil, mongoError
	}
	return &session, nil
}

func GetTest(sessionId, testId string) (*model.TestEntity, error) {
	session, mongoError := GetSession(sessionId)
	if mongoError != nil {
		log.Printf("Mongo error: %s", mongoError)
		return nil, mongoError
	}
	if session == nil {
		return nil, nil
	}
	for _, test := range session.Tests {
		if test.Id == testId {
			return &test, nil
		}
	}
	return nil, nil
}

func GetReport(reportId string) (*model.ReportEntity, error) {
	id, mongoError := primitive.ObjectIDFromHex(reportId)
	if mongoError != nil {
		log.Printf("Mongo error: %s", mongoError)
		return nil, mongoError
	}
	collection := storage.client.Database(constants.Trots).Collection(constants.Report)
	mongoContext, cancel := context.WithTimeout(context.Background(), provider.Configuration.Mongo.Timeout)
	defer cancel()
	cursor, mongoError := collection.Find(mongoContext, bson.D{{constants.MongoId, id}})
	if mongoError != nil {
		log.Printf("Mongo error: %s", mongoError)
		return nil, mongoError
	}
	if !cursor.Next(mongoContext) {
		return nil, nil
	}
	var report model.ReportEntity
	if mongoError := cursor.Decode(&report); mongoError != nil {
		log.Printf("Mongo error: %s", mongoError)
		return nil, mongoError
	}
	return &report, nil
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
		if mongoError = cursor.Decode(&session); mongoError != nil {
			log.Printf("Mongo error: %s", mongoError)
			return nil, mongoError
		}
		sessions = append(sessions, session)
	}
	return &sessions, nil
}

func connect() *MongoStorage {
	client, mongoError := mongo.NewClient(options.Client().ApplyURI(provider.Configuration.Mongo.Address))
	if mongoError != nil {
		log.Panicf("Mongo error: %s", mongoError)
	}
	mongoContext, cancel := context.WithTimeout(context.Background(), provider.Configuration.Mongo.Timeout)
	defer cancel()
	if mongoError = client.Connect(mongoContext); mongoError != nil {
		log.Panicf("Mongo error: %s", mongoError)
	}
	log.Printf("Mongo client connected to: %s", provider.Configuration.Mongo.Address)
	return &MongoStorage{client: client}
}

func Disconnect() {
	if disconnectError := storage.client.Disconnect(context.Background()); disconnectError != nil {
		log.Printf("Mongo error: %s", disconnectError)
	}
}
