package friendrepo

import (
	"context"
	"cs_chat_app_server/common"
	friendmodel "cs_chat_app_server/modules/friend/model"
	requeststore "cs_chat_app_server/modules/request/store"
)

// FindUser is a method of finding a user
func (repo *friendRepository) FindUser(
	ctx context.Context,
	filter map[string]interface{},
	options ...FindUserOption,
) (*friendmodel.User, error) {
	user, err := repo.friendstore.FindUser(ctx, filter)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, common.ErrEntityNotFound("User", friendmodel.ErrUserNotFound)
	}

	for _, option := range options {
		err = option(ctx, repo, user)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

type FindUserOption func(ctx context.Context, repo *friendRepository, u *friendmodel.User) error

// WithRelation requires requesterId, which is the id of the requester,
//
// It will use the requesterId to query the relation of user with requester
// The relation includes
// [friendmodel.Self], [friendmodel.Friend], [friendmodel.Blocked],
// [friendmodel.Non], [friendmodel.Sent], [friendmodel.Received],
func WithRelation(requesterId string) FindUserOption {
	return func(ctx context.Context, repo *friendRepository, user *friendmodel.User) error {
		if *user.Id == requesterId {
			user.Relation = friendmodel.Self
			return nil
		}

		for _, friendId := range user.Friends {
			if friendId == requesterId {
				user.Relation = friendmodel.Friend
				return nil
			}
		}

		filter := make(map[string]interface{})
		err := common.AddIdFilter(filter, requesterId)

		if err != nil {
			return err
		}

		requester, err := repo.friendstore.FindUser(ctx, filter)

		if err != nil {
			return err
		}

		if requester == nil {
			return common.ErrEntityNotFound("User", friendmodel.ErrUserNotFound)
		}

		for _, blockedId := range requester.BlockedUser {
			if *user.Id == blockedId {
				user.Relation = friendmodel.Blocked
				return nil
			}
		}

		filter = common.GetOrFilter(
			common.GetAndFilter(
				requeststore.GetRequestSenderIdFilter(requesterId),
				requeststore.GetRequestReceiverIdFilter(*user.Id),
			),
			common.GetAndFilter(
				requeststore.GetRequestSenderIdFilter(*user.Id),
				requeststore.GetRequestReceiverIdFilter(requesterId),
			),
		)

		request, err := repo.requestStore.FindRequest(ctx, filter)
		if err != nil {
			return err
		}

		if request == nil {
			user.Relation = friendmodel.Non
			return nil
		}

		if request.Sender.Id == requesterId {
			user.Relation = friendmodel.Sent
			return nil
		}

		user.Relation = friendmodel.Received
		return nil
	}
}
