package friendbiz

import (
	"context"
	"cs_chat_app_server/common"
	friendmodel "cs_chat_app_server/modules/friend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FriendStore interface {
	CreateRequest(ctx context.Context, request *friendmodel.Request) error
	FindRequest(ctx context.Context, userId string, otherId string) (*friendmodel.Request, error)
	FindUser(ctx context.Context, filter map[string]interface{}) (*friendmodel.User, error)
}

type sendRequestBiz struct {
	friendStore FriendStore
}

func NewSendRequestBiz(friendStore FriendStore) *sendRequestBiz {
	return &sendRequestBiz{
		friendStore: friendStore,
	}
}

func (biz *sendRequestBiz) SendRequest(ctx context.Context, senderId string, receiverId string) error {
	existedRequest, err := biz.friendStore.FindRequest(ctx, senderId, receiverId)
	if err != nil {
		return err
	}
	if existedRequest != nil {
		return common.ErrInvalidRequest(friendmodel.ErrRequestExists)
	}
	id, _ := primitive.ObjectIDFromHex(senderId)
	sender, err := biz.friendStore.FindUser(ctx, map[string]interface{}{
		"_id": id,
	})
	if err != nil {
		return err
	}
	if sender == nil {
		return common.ErrEntityNotFound("User", nil)
	}

	for i := range sender.Friends {
		if sender.Friends[i] == receiverId {
			return common.ErrInvalidRequest(friendmodel.ErrHasBeenFriend)
		}
	}

	id, _ = primitive.ObjectIDFromHex(receiverId)
	receiver, err := biz.friendStore.FindUser(ctx, map[string]interface{}{
		"_id": id,
	})
	if err != nil {
		return err
	}
	if receiver == nil {
		return common.ErrEntityNotFound("User", nil)
	}
	senderRequestUser := friendmodel.RequestUser{
		Id:     senderId,
		Name:   sender.Name,
		Avatar: sender.Avatar,
	}
	receiverRequestUser := friendmodel.RequestUser{
		Id:     receiverId,
		Name:   receiver.Name,
		Avatar: receiver.Avatar,
	}
	request := friendmodel.Request{
		Sender:   senderRequestUser,
		Receiver: receiverRequestUser,
	}
	request.Process()
	err = biz.friendStore.CreateRequest(ctx, &request)
	if err != nil {
		return err
	}
	return nil
}
