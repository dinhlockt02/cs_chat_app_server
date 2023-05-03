package userrepo

import (
	"context"
	"cs_chat_app_server/common"
	friendmodel "cs_chat_app_server/modules/friend/model"
	usermodel "cs_chat_app_server/modules/user/model"
)

type FindUserUserStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*usermodel.User, error)
}

type FindUserFriendStore interface {
	FindUser(ctx context.Context, filter map[string]interface{}) (*friendmodel.User, error)
}

type findUserRepo struct {
	userStore   FindUserUserStore
	friendStore FindUserFriendStore
}

func NewFindUserRepo(
	userStore FindUserUserStore,
	friendStore FindUserFriendStore,
) *findUserRepo {
	return &findUserRepo{
		userStore:   userStore,
		friendStore: friendStore,
	}
}

func (repo *findUserRepo) FindUser(ctx context.Context, requesterId string, filter map[string]interface{}) (*usermodel.User, error) {

	{

		user, err := repo.friendStore.FindUser(ctx, filter)

		if err != nil {
			return nil, err
		}

		if user == nil {
			return nil, common.ErrEntityNotFound("User", usermodel.ErrUserNotFound)
		}

		for _, blockedId := range user.BlockedUser {
			if blockedId == requesterId {
				return nil, common.ErrForbidden(usermodel.ErrUserBeBlocked)
			}
		}
	}

	user, err := repo.userStore.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	return user, nil
}
