package friendrepo

import (
	"context"
	friendmodel "cs_chat_app_server/modules/friend/model"
	friendstore "cs_chat_app_server/modules/friend/store"
	requestmdl "cs_chat_app_server/modules/request/model"
	requeststore "cs_chat_app_server/modules/request/store"
)

type Repository interface {
	FindRequest(
		ctx context.Context,
		sender string,
		receiver string,
	) (*requestmdl.Request, error)
	FindUser(
		ctx context.Context,
		filter map[string]interface{},
		options ...FindUserOption,
	) (*friendmodel.User, error)
	UpdateUser(
		ctx context.Context,
		filter map[string]interface{},
		user *friendmodel.User,
	) error
	DeleteRequest(
		ctx context.Context,
		filter map[string]interface{},
	) error
	FindRequests(
		ctx context.Context,
		filter map[string]interface{},
	) ([]requestmdl.Request, error)
	CreateRequest(
		ctx context.Context,
		request *requestmdl.Request,
	) error
}

type friendRepository struct {
	friendstore  friendstore.Store
	requestStore requeststore.Store
}

func NewFriendRepository(
	friendstore friendstore.Store,
	requestStore requeststore.Store,
) *friendRepository {
	return &friendRepository{
		friendstore:  friendstore,
		requestStore: requestStore,
	}
}
