package friendbiz

import (
	"context"
	"cs_chat_app_server/common"
	friendmodel "cs_chat_app_server/modules/friend/model"
	friendrepo "cs_chat_app_server/modules/friend/repository"
)

type rejectRequestBiz struct {
	friendRepo friendrepo.Repository
}

func NewRejectRequestBiz(friendRepo friendrepo.Repository) *rejectRequestBiz {
	return &rejectRequestBiz{
		friendRepo: friendRepo,
	}
}

func (biz *rejectRequestBiz) RejectRequest(ctx context.Context, senderId string, receiverId string) error {
	existedRequest, err := biz.friendRepo.FindRequest(ctx, senderId, receiverId)
	if err != nil {
		return err
	}
	if existedRequest == nil {
		return common.ErrInvalidRequest(friendmodel.ErrRequestNotFound)
	}
	filter := make(map[string]interface{})
	err = common.AddIdFilter(filter, *existedRequest.Id)
	if err != nil {
		return err
	}
	err = biz.friendRepo.DeleteRequest(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
