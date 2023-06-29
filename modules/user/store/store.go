package userstore

import "go.mongodb.org/mongo-driver/mongo"

type Store interface {
}

type mongoStore struct {
	database *mongo.Database
}

func NewMongoStore(database *mongo.Database) *mongoStore {
	return &mongoStore{database: database}
}

func GetGroupsFilter(id string) map[string]interface{} {
	return map[string]interface{}{
		"groups": id,
	}
}
