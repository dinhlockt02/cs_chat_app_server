package friendbiz

import (
	"context"
	"cs_chat_app_server/common"
	friendrepo "cs_chat_app_server/modules/friend/repository"
	requestmdl "cs_chat_app_server/modules/request/model"
	requeststore "cs_chat_app_server/modules/request/store"
)

type getSentRequestBiz struct {
	friendRepo friendrepo.Repository
}

func NewGetSentRequestBiz(friendRepo friendrepo.Repository) *getSentRequestBiz {
	return &getSentRequestBiz{
		friendRepo: friendRepo,
	}
}

func (biz *getSentRequestBiz) GetSentRequest(ctx context.Context, senderId string) ([]requestmdl.Request, error) {
	requests, err := biz.friendRepo.FindRequests(ctx, common.GetAndFilter(
		requeststore.GetRequestSenderIdFilter(senderId),
		requeststore.GetTypeFilterFilter(false),
	))
	if err != nil {
		return nil, err
	}

	return requests, nil
}
