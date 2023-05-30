package friendbiz

import (
	"context"
	"cs_chat_app_server/common"
	friendrepo "cs_chat_app_server/modules/friend/repository"
	requestmdl "cs_chat_app_server/modules/request/model"
	requeststore "cs_chat_app_server/modules/request/store"
)

type getReceivedRequestBiz struct {
	friendRepo friendrepo.Repository
}

func NewGetReceivedRequestBiz(friendRepo friendrepo.Repository) *getReceivedRequestBiz {
	return &getReceivedRequestBiz{
		friendRepo: friendRepo,
	}
}

func (biz *getReceivedRequestBiz) GetReceivedRequest(ctx context.Context, receiverId string) ([]requestmdl.Request, error) {

	filter := common.GetAndFilter(
		requeststore.GetRequestReceiverIdFilter(receiverId),
		requeststore.GetTypeFilterFilter(false),
	)

	requests, err := biz.friendRepo.FindRequests(ctx, filter)
	if err != nil {
		return nil, err
	}

	return requests, nil
}
