package friendrepo

import (
	"context"
	"cs_chat_app_server/common"
	friendmodel "cs_chat_app_server/modules/friend/model"
)

type FindUserFriendStore interface {
	FindRequest(ctx context.Context, userId string, otherId string) (*friendmodel.Request, error)
	FindUser(ctx context.Context, filter map[string]interface{}) (*friendmodel.User, error)
}

type findUserRepository struct {
	friendstore FindUserFriendStore
}

func NewFindUserRepository(friendstore FindUserFriendStore) *findUserRepository {
	return &findUserRepository{friendstore: friendstore}
}

// FindUser is a method of finding a user
// It requires requesterId, which is the id of the requester,
// and filter which will be used for filtered user.
//
// It will use the requesterId to query the relation of user with requester
// The relation is an enum provided in [friendmodel] package
func (repo *findUserRepository) FindUser(ctx context.Context, requesterId string, filter map[string]interface{}) (*friendmodel.User, error) {
	user, err := repo.friendstore.FindUser(ctx, filter)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, common.ErrEntityNotFound("User", friendmodel.ErrUserNotFound)
	}

	if *user.Id == requesterId {
		user.Relation = friendmodel.Self
		return user, nil
	}

	for _, friendId := range user.Friends {
		if friendId == requesterId {
			user.Relation = friendmodel.Friend
			return user, nil
		}
	}

	for _, blockedId := range user.BlockedUser {
		if blockedId == requesterId {
			return user, common.ErrForbidden(friendmodel.ErrUserBeBlocked)
		}
	}

	filter = make(map[string]interface{})
	err = common.AddIdToFilter(filter, requesterId)

	if err != nil {
		return nil, err
	}

	requester, err := repo.friendstore.FindUser(ctx, filter)

	if err != nil {
		return nil, err
	}

	if requester == nil {
		return nil, common.ErrEntityNotFound("User", friendmodel.ErrUserNotFound)
	}

	for _, blockedId := range requester.BlockedUser {
		if *user.Id == blockedId {
			user.Relation = friendmodel.Blocked
			return user, nil
		}
	}

	request, err := repo.friendstore.FindRequest(ctx, requesterId, *user.Id)
	if err != nil {
		return nil, err
	}

	if request == nil {
		user.Relation = friendmodel.Non
		return user, nil
	}

	if request.Sender.Id == requesterId {
		user.Relation = friendmodel.Sent
		return user, nil
	}

	user.Relation = friendmodel.Received
	return user, nil
}
