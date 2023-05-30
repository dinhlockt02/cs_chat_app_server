package userrepo

import (
	"context"
	"cs_chat_app_server/common"
	friendrepo "cs_chat_app_server/modules/friend/repository"
	usermodel "cs_chat_app_server/modules/user/model"
)

type FindUserUserStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*usermodel.User, error)
}

type findUserRepo struct {
	userStore  FindUserUserStore
	friendRepo friendrepo.Repository
}

func NewFindUserRepo(
	userStore FindUserUserStore,
	friendRepo friendrepo.Repository,
) *findUserRepo {
	return &findUserRepo{
		userStore:  userStore,
		friendRepo: friendRepo,
	}
}

func (repo *findUserRepo) FindUser(ctx context.Context, requesterId string, filter map[string]interface{}) (*usermodel.User, error) {
	fuser, err := repo.friendRepo.FindUser(ctx, filter, friendrepo.WithRelation(requesterId))

	if err != nil {
		return nil, err
	}

	user, err := repo.userStore.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	user.Relation = fuser.Relation

	filter = map[string]interface{}{}

	err = common.AddIdFilter(filter, requesterId)
	if err != nil {
		return nil, err
	}

	requester, err := repo.friendRepo.FindUser(ctx, filter)

	requesterFriendMap := make(map[string]interface{}, len(requester.Friends))

	for _, friend := range requester.Friends {
		requesterFriendMap[friend] = struct{}{}
	}

	commonFriendCount := 0
	for _, friend := range fuser.Friends {
		if _, ok := requesterFriendMap[friend]; ok {
			commonFriendCount++
		}
	}

	user.CommonFriendCount = &commonFriendCount

	return user, nil
}
