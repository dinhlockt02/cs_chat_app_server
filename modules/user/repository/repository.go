package userrepo

import (
	"context"
	usermodel "cs_chat_app_server/modules/user/model"
	userstore "cs_chat_app_server/modules/user/store"
)

type Repository interface {
	FindUser(
		ctx context.Context,
		filter map[string]interface{},
	) (*usermodel.User, error)
}

type UserRepository struct {
	userStore userstore.Store
}

func NewUserRepository(
	userStore userstore.Store,
) *UserRepository {
	return &UserRepository{
		userStore: userStore,
	}
}
