package friendstore

import (
	"cs_chat_app_server/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoStore struct {
	database *mongo.Database
}

func NewMongoStore(database *mongo.Database) *mongoStore {
	return &mongoStore{database: database}
}

func (s *mongoStore) NewFilterById(id string) (map[string]interface{}, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}
	return map[string]interface{}{
		"_id": _id,
	}, nil
}
