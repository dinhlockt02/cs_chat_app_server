package userrepo

import (
	"context"
	friendmodel "cs_chat_app_server/modules/friend/model"
	usermodel "cs_chat_app_server/modules/user/model"
)

type FindUserUserStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*usermodel.User, error)
}

type FindUserFriendRepository interface {
	FindUser(ctx context.Context, requesterId string, filter map[string]interface{}) (*friendmodel.User, error)
}

type findUserRepo struct {
	userStore  FindUserUserStore
	friendRepo FindUserFriendRepository
}

func NewFindUserRepo(
	userStore FindUserUserStore,
	friendRepo FindUserFriendRepository,
) *findUserRepo {
	return &findUserRepo{
		userStore:  userStore,
		friendRepo: friendRepo,
	}
}

func (repo *findUserRepo) FindUser(ctx context.Context, requesterId string, filter map[string]interface{}) (*usermodel.User, error) {
	fuser, err := repo.friendRepo.FindUser(ctx, requesterId, filter)

	if err != nil {
		return nil, err
	}

	user, err := repo.userStore.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	user.Relation = fuser.Relation

	return user, nil
}
