package friendbiz

import (
	"context"
	"cs_chat_app_server/common"
	friendmodel "cs_chat_app_server/modules/friend/model"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FindFriendFriendStore interface {
	FindFriend(ctx context.Context, filter map[string]interface{}) ([]friendmodel.FriendUser, error)
	FindUser(ctx context.Context, filter map[string]interface{}) (*friendmodel.User, error)
}

type findFriendBiz struct {
	friendStore FindFriendFriendStore
}

func NewFindFriendBiz(friendStore FindFriendFriendStore) *findFriendBiz {
	return &findFriendBiz{friendStore: friendStore}
}

func (biz *findFriendBiz) FindFriend(ctx context.Context, userId string) ([]friendmodel.FriendUser, error) {
	id, _ := primitive.ObjectIDFromHex(userId)
	user, err := biz.friendStore.FindUser(ctx, map[string]interface{}{
		"_id": id,
	})

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, common.ErrEntityNotFound("User", errors.New("user not found"))
	}

	var ids = make([]primitive.ObjectID, 0, len(user.Friends))
	for _, friend := range user.Friends {
		id, err = primitive.ObjectIDFromHex(friend)
		if err != nil {
			return nil, common.ErrInternal(err)
		}
		ids = append(ids, id)
	}
	friends, err := biz.friendStore.FindFriend(ctx, map[string]interface{}{
		"_id": map[string]interface{}{
			"$in": ids,
		},
	})
	if err != nil {
		return nil, err
	}
	return friends, nil
}
