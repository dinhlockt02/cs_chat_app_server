package friendbiz

import (
	"context"
	"cs_chat_app_server/common"
	friendmodel "cs_chat_app_server/modules/friend/model"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AcceptRequestFriendStore interface {
	FindUser(ctx context.Context, filter map[string]interface{}) (*friendmodel.User, error)
	UpdateUser(ctx context.Context, filter map[string]interface{}, updatedUser *friendmodel.User) error
	DeleteRequest(ctx context.Context, requestId string) error
	FindRequest(ctx context.Context, userId string, otherId string) (*friendmodel.Request, error)
}

type acceptRequestBiz struct {
	friendStore AcceptRequestFriendStore
}

func NewAcceptRequestBiz(friendStore AcceptRequestFriendStore) *acceptRequestBiz {
	return &acceptRequestBiz{
		friendStore: friendStore,
	}
}

func (biz *acceptRequestBiz) AcceptRequest(ctx context.Context, senderId string, receiverId string) error {
	existedRequest, err := biz.friendStore.FindRequest(ctx, senderId, receiverId)
	if err != nil {
		return err
	}
	if existedRequest == nil {
		return common.ErrInvalidRequest(friendmodel.ErrRequestNotFound)
	}
	id, _ := primitive.ObjectIDFromHex(senderId)
	sender, err := biz.friendStore.FindUser(ctx, map[string]interface{}{
		"_id": id,
	})

	if err != nil {
		return err
	}
	if sender == nil {
		return common.ErrEntityNotFound("User", errors.New("sender not found"))
	}
	id, _ = primitive.ObjectIDFromHex(receiverId)
	receiver, err := biz.friendStore.FindUser(ctx, map[string]interface{}{
		"_id": id,
	})
	if receiver == nil {
		return common.ErrEntityNotFound("User", errors.New("receiver not found"))
	}
	sender.Friends = append(sender.Friends, receiverId)
	receiver.Friends = append(receiver.Friends, senderId)

	id, _ = primitive.ObjectIDFromHex(senderId)
	err = biz.friendStore.UpdateUser(ctx, map[string]interface{}{
		"_id": id,
	}, sender)
	if err != nil {
		return err
	}

	id, _ = primitive.ObjectIDFromHex(receiverId)
	err = biz.friendStore.UpdateUser(ctx, map[string]interface{}{
		"_id": id,
	}, receiver)
	if err != nil {
		return err
	}
	err = biz.friendStore.DeleteRequest(ctx, *existedRequest.Id)
	if err != nil {
		return err
	}
	return nil
}
