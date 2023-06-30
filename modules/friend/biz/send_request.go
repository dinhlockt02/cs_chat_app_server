package friendbiz

import (
	"context"
	"cs_chat_app_server/common"
	notimodel "cs_chat_app_server/components/notification/model"
	notirepo "cs_chat_app_server/components/notification/repository"
	friendmodel "cs_chat_app_server/modules/friend/model"
	friendrepo "cs_chat_app_server/modules/friend/repository"
	requestmdl "cs_chat_app_server/modules/request/model"
	"errors"
	"github.com/rs/zerolog/log"
)

type sendRequestBiz struct {
	friendRepo   friendrepo.Repository
	notification notirepo.NotificationServiceRepository
}

func NewSendRequestBiz(
	friendRepo friendrepo.Repository,
	notification notirepo.NotificationServiceRepository,
) *sendRequestBiz {
	return &sendRequestBiz{
		friendRepo:   friendRepo,
		notification: notification,
	}
}

func (biz *sendRequestBiz) SendRequest(ctx context.Context, senderId string, receiverId string) error {
	// Find exists request
	existedRequest, err := biz.friendRepo.FindRequest(ctx, senderId, receiverId)
	if err != nil {
		return err
	}
	if existedRequest != nil {
		return common.ErrInvalidRequest(friendmodel.ErrRequestExists)
	}

	// Find sender
	filter := make(map[string]interface{})
	err = common.AddIdFilter(filter, senderId)
	sender, err := biz.friendRepo.FindUser(ctx, filter)
	if err != nil {
		return err
	}
	if sender == nil {
		return common.ErrEntityNotFound("User", errors.New("sender not found"))
	}

	// Find Receiver
	filter = make(map[string]interface{})
	err = common.AddIdFilter(filter, receiverId)
	receiver, err := biz.friendRepo.FindUser(ctx, filter)
	if receiver == nil {
		return common.ErrEntityNotFound("User", errors.New("receiver not found"))
	}

	senderRequestUser := requestmdl.RequestUser{
		Id:     senderId,
		Name:   sender.Name,
		Avatar: sender.Avatar,
	}
	receiverRequestUser := requestmdl.RequestUser{
		Id:     receiverId,
		Name:   receiver.Name,
		Avatar: receiver.Avatar,
	}
	request := requestmdl.Request{
		Sender:   senderRequestUser,
		Receiver: receiverRequestUser,
	}
	request.Process()
	err = biz.friendRepo.CreateRequest(ctx, &request)
	if err != nil {
		return err
	}

	go func() {
		e := biz.notification.CreateReceiveFriendRequestNotification(context.Background(), receiverId, &notimodel.NotificationObject{
			Id:    receiverId,
			Name:  receiver.Name,
			Image: &receiver.Avatar,
			Type:  notimodel.User,
		},
			&notimodel.NotificationObject{
				Id:    *request.Id,
				Name:  "",
				Image: nil,
				Type:  notimodel.Request,
			},
			&notimodel.NotificationObject{
				Id:    senderId,
				Name:  sender.Name,
				Image: &sender.Avatar,
				Type:  notimodel.User,
			})
		if e != nil {
			log.Err(e)
		}
	}()

	return nil
}
