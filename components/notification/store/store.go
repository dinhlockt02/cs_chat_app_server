package notistore

import "go.mongodb.org/mongo-driver/mongo"

type mongoStore struct {
	database *mongo.Database
}

func NewMongoStore(database *mongo.Database) *mongoStore {
	return &mongoStore{database: database}
}
