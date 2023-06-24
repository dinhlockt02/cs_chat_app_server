package groupstore

import (
	"context"
	groupmdl "cs_chat_app_server/modules/group/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	Create(ctx context.Context, group *groupmdl.Group) error
	List(
		ctx context.Context,
		filter map[string]interface{},
	) ([]groupmdl.Group, error)

	FindUser(
		ctx context.Context,
		filter map[string]interface{},
	) (*groupmdl.User, error)
	UpdateUser(
		ctx context.Context,
		filter map[string]interface{},
		updatedUser *groupmdl.User,
	) error
	FindGroup(
		ctx context.Context,
		filter map[string]interface{},
	) (*groupmdl.Group, error)
	UpdateGroup(
		ctx context.Context,
		filter map[string]interface{},
		updatedGroup *groupmdl.Group,
	) error
	UpdateGroupMember(
		ctx context.Context,
		filter map[string]interface{},
		updatedMember *groupmdl.GroupUser) error
	FindUsers(ctx context.Context, filter map[string]interface{}) ([]groupmdl.User, error)
	UpdateLatestMessage(
		ctx context.Context,
		filter map[string]interface{},
		latestMessage *groupmdl.GroupMessage) error
}

type mongoStore struct {
	database *mongo.Database
}

func NewMongoStore(database *mongo.Database) *mongoStore {
	return &mongoStore{database: database}
}
