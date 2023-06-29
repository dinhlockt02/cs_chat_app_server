package groupbiz

import (
	"context"
	"cs_chat_app_server/common"
	notirepo "cs_chat_app_server/components/notification/repository"
	friendmodel "cs_chat_app_server/modules/friend/model"
	groupmdl "cs_chat_app_server/modules/group/model"
	grouprepo "cs_chat_app_server/modules/group/repository"
	requestmdl "cs_chat_app_server/modules/request/model"
	requeststore "cs_chat_app_server/modules/request/store"
	"errors"
)

type sendGroupRequestBiz struct {
	groupRepo    grouprepo.Repository
	notification notirepo.NotificationServiceRepository
}

func NewSendGroupRequestBiz(groupRepo grouprepo.Repository) *sendGroupRequestBiz {
	return &sendGroupRequestBiz{groupRepo: groupRepo}
}

// SendRequest send a group invitation request to user.
func (biz *sendGroupRequestBiz) SendRequest(ctx context.Context, requester string, user string, group *groupmdl.Group) error {

	// TODO: Allow send request only if requester is a member of group

	// Find exists request
	requesterFilter := requeststore.GetRequestReceiverIdFilter(user)
	groupFilter := requeststore.GetRequestGroupIdFilter(*group.Id)
	ft := common.GetAndFilter(requesterFilter, groupFilter)
	existedRequest, err := biz.groupRepo.FindRequest(ctx, ft)
	if err != nil {
		return err
	}
	if existedRequest != nil {
		return common.ErrInvalidRequest(friendmodel.ErrRequestExists)
	}

	// Find sender
	filter := make(map[string]interface{})
	err = common.AddIdFilter(filter, requester)
	sender, err := biz.groupRepo.FindUser(ctx, filter)
	if err != nil {
		return err
	}
	if sender == nil {
		return common.ErrEntityNotFound("User", errors.New("sender not found"))
	}

	// Find Receiver
	filter = make(map[string]interface{})
	err = common.AddIdFilter(filter, user)
	receiver, err := biz.groupRepo.FindUser(ctx, filter)
	if receiver == nil {
		return common.ErrEntityNotFound("User", errors.New("receiver not found"))
	}

	senderRequestUser := requestmdl.RequestUser{
		Id:     requester,
		Name:   sender.Name,
		Avatar: sender.Avatar,
	}
	receiverRequestUser := requestmdl.RequestUser{
		Id:     user,
		Name:   receiver.Name,
		Avatar: receiver.Avatar,
	}
	groupRequest := requestmdl.RequestGroup{
		Id:       *group.Id,
		Name:     group.Name,
		ImageUrl: *group.ImageUrl,
	}
	req := requestmdl.Request{
		Sender:   senderRequestUser,
		Receiver: receiverRequestUser,
		Group:    &groupRequest,
	}
	req.Process()
	err = biz.groupRepo.CreateRequest(ctx, &req)
	if err != nil {
		return err
	}

	go func() {
		// TODO: Push notification group request
		//e := biz.notification.CreateReceiveFriendRequestNotification(
		//	context.Background(), receiverId, &notimodel.NotificationObject{
		//		Id:    receiverId,
		//		Name:  receiver.Name,
		//		Image: &receiver.Avatar,
		//		Type:  notimodel.User,
		//	}, &notimodel.NotificationObject{
		//		Id:    senderId,
		//		Name:  sender.Name,
		//		Image: &sender.Avatar,
		//		Type:  notimodel.User,
		//	})
		//if e != nil {
		//	log.Err(e)
		//}
	}()

	return nil
}
