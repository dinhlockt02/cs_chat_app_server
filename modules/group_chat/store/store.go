package gchatstore

import (
	"context"
	"cs_chat_app_server/common"
	gchatmdl "cs_chat_app_server/modules/group_chat/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	Create(ctx context.Context, groupChatItem *gchatmdl.GroupChatItem) error
	FindUser(ctx context.Context, filter map[string]interface{}) (*gchatmdl.User, error)
	List(
		ctx context.Context,
		filter map[string]interface{},
		paging gchatmdl.Paging,
	) ([]gchatmdl.GroupChatItem, error)
	FindMessage(
		ctx context.Context,
		filter map[string]interface{},
	) (*gchatmdl.GroupChatItem, error)
	FindGroup(
		ctx context.Context,
		filter map[string]interface{},
	) (*gchatmdl.Group, error)
}

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

func GetMessageTypeFilter(messageType string) map[string]interface{} {
	return map[string]interface{}{
		"type": messageType,
	}
}
