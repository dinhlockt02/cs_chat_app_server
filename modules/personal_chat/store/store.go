package pchatstore

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

func (s *mongoStore) AddIdFilter(id string, filter map[string]interface{}) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return common.ErrInvalidRequest(err)
	}
	filter["_id"] = _id
	return nil
}

func GetSenderFilter(senderId string) map[string]interface{} {
	return map[string]interface{}{
		"sender": senderId,
	}
}

func GetReceiverFilter(receiverId string) map[string]interface{} {
	return map[string]interface{}{
		"receiver": receiverId,
	}
}

func GetMessageTypeFilter(messageType string) map[string]interface{} {
	return map[string]interface{}{
		"type": messageType,
	}
}
