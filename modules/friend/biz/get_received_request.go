package friendbiz

import (
	"context"
	friendmodel "cs_chat_app_server/modules/friend/model"
)

type GetReceivedRequestFriendStore interface {
	FindRequests(ctx context.Context, filter map[string]interface{}) ([]friendmodel.Request, error)
}

type getReceivedRequestBiz struct {
	friendStore GetReceivedRequestFriendStore
}

func NewGetReceivedRequestBiz(friendStore GetReceivedRequestFriendStore) *getReceivedRequestBiz {
	return &getReceivedRequestBiz{
		friendStore: friendStore,
	}
}

func (biz *getReceivedRequestBiz) GetReceivedRequest(ctx context.Context, receiverId string) ([]friendmodel.Request, error) {
	requests, err := biz.friendStore.FindRequests(ctx, map[string]interface{}{
		"receiver.id": receiverId,
	})
	if err != nil {
		return nil, err
	}

	return requests, nil
}
