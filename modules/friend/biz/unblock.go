package friendbiz

import (
	"context"
	"cs_chat_app_server/common"
	friendmodel "cs_chat_app_server/modules/friend/model"
	"errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UnblockFriendStore interface {
	FindUser(ctx context.Context, filter map[string]interface{}) (*friendmodel.User, error)
	UpdateUser(ctx context.Context, filter map[string]interface{}, updatedUser *friendmodel.User) error
}

type unblockBiz struct {
	friendStore UnfriendFriendStore
}

func NewUnblockBiz(friendStore UnfriendFriendStore) *unblockBiz {
	return &unblockBiz{
		friendStore: friendStore,
	}
}

func (biz *unblockBiz) Unblock(ctx context.Context, userId string, blockedId string) error {
	id, _ := primitive.ObjectIDFromHex(userId)
	user, err := biz.friendStore.FindUser(ctx, map[string]interface{}{
		"_id": id,
	})

	if err != nil {
		return err
	}
	if user == nil {
		return common.ErrEntityNotFound("User", errors.New("user not found"))
	}
	for i := range user.BlockedUser {
		if user.BlockedUser[i] == blockedId {

			user.BlockedUser = append(user.BlockedUser[:i], user.BlockedUser[i+1:]...)
			log.Debug().Msgf("%v", user.BlockedUser)

			err = biz.friendStore.UpdateUser(ctx, map[string]interface{}{
				"_id": id,
			}, user)
			if err != nil {
				return err
			}
			break
		}
	}
	return nil
}
