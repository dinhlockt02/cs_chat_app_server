package friendbiz

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/pubsub"
	friendmodel "cs_chat_app_server/modules/friend/model"
	friendrepo "cs_chat_app_server/modules/friend/repository"
)

type rejectRequestBiz struct {
	friendRepo friendrepo.Repository
	ps         pubsub.PubSub
}

func NewRejectRequestBiz(friendRepo friendrepo.Repository, ps pubsub.PubSub) *rejectRequestBiz {
	return &rejectRequestBiz{
		friendRepo: friendRepo,
		ps:         ps,
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
	biz.ps.Publish(ctx, common.TopicDenyFriendRequest, *existedRequest.Id)
	return nil
}
