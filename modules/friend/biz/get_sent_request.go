package friendbiz

import (
	"context"
	friendmodel "cs_chat_app_server/modules/friend/model"
)

type GetSentRequestFriendStore interface {
	FindRequests(ctx context.Context, filter map[string]interface{}) ([]friendmodel.Request, error)
}

type getSentRequestBiz struct {
	friendStore GetSentRequestFriendStore
}

func NewGetSentRequestBiz(friendStore GetSentRequestFriendStore) *getSentRequestBiz {
	return &getSentRequestBiz{
		friendStore: friendStore,
	}
}

func (biz *getSentRequestBiz) GetSentRequest(ctx context.Context, senderId string) ([]friendmodel.Request, error) {
	requests, err := biz.friendStore.FindRequests(ctx, map[string]interface{}{
		"sender.id": senderId,
	})
	if err != nil {
		return nil, err
	}

	return requests, nil
}
