package friendbiz

import (
	"context"
	"cs_chat_app_server/common"
	notimodel "cs_chat_app_server/components/notification/model"
	notirepo "cs_chat_app_server/components/notification/repository"
	friendmodel "cs_chat_app_server/modules/friend/model"
	friendrepo "cs_chat_app_server/modules/friend/repository"
	groupmdl "cs_chat_app_server/modules/group/model"
	grouprepo "cs_chat_app_server/modules/group/repository"
	groupstore "cs_chat_app_server/modules/group/store"
	"errors"
	"github.com/rs/zerolog/log"
)

type acceptRequestBiz struct {
	friendRepository friendrepo.Repository
	notification     notirepo.NotificationRepository
	groupRepository  grouprepo.Repository
}

func NewAcceptRequestBiz(
	friendRepository friendrepo.Repository,
	notification notirepo.NotificationRepository,
	groupRepository grouprepo.Repository,
) *acceptRequestBiz {
	return &acceptRequestBiz{
		friendRepository: friendRepository,
		notification:     notification,
		groupRepository:  groupRepository,
	}
}

func (biz *acceptRequestBiz) AcceptRequest(ctx context.Context, senderId string, receiverId string) error {

	// Find exists request
	existedRequest, err := biz.friendRepository.FindRequest(ctx, senderId, receiverId)
	if err != nil {
		return err
	}
	if existedRequest == nil {
		return common.ErrInvalidRequest(friendmodel.ErrRequestNotFound)
	}

	// Find sender
	filter := make(map[string]interface{})
	err = common.AddIdFilter(filter, senderId)
	sender, err := biz.friendRepository.FindUser(ctx, filter)
	if err != nil {
		return err
	}
	if sender == nil {
		return common.ErrEntityNotFound("User", errors.New("sender not found"))
	}

	// Find Receiver
	filter = make(map[string]interface{})
	err = common.AddIdFilter(filter, receiverId)
	receiver, err := biz.friendRepository.FindUser(ctx, filter)
	if receiver == nil {
		return common.ErrEntityNotFound("User", errors.New("receiver not found"))
	}

	// Update Sender
	sender.Friends = append(sender.Friends, receiverId)
	filter = make(map[string]interface{})
	err = common.AddIdFilter(filter, senderId)
	if err != nil {
		return err
	}

	err = biz.friendRepository.UpdateUser(ctx, filter, sender)
	if err != nil {
		return err
	}

	// Update Receiver
	receiver.Friends = append(receiver.Friends, senderId)
	filter = make(map[string]interface{})
	err = common.AddIdFilter(filter, receiverId)
	if err != nil {
		return err
	}
	err = biz.friendRepository.UpdateUser(ctx, filter, receiver)
	if err != nil {
		return err
	}

	// Delete request
	filter = make(map[string]interface{})
	err = common.AddIdFilter(filter, *existedRequest.Id)
	err = biz.friendRepository.DeleteRequest(ctx, filter)
	if err != nil {
		return err
	}

	// Find group
	group, err := biz.groupRepository.FindGroup(ctx,
		common.GetAndFilter(
			groupstore.GetUserIdInIdListFilter(receiverId, senderId),
			groupstore.GetTypeFilter(groupmdl.TypePersonal),
		),
	)
	if err != nil {
		return err
	}

	if group == nil {
		// Create group
		err = biz.groupRepository.CreateGroup(ctx, &groupmdl.Group{
			Members: []groupmdl.GroupUser{{
				Id:     senderId,
				Name:   sender.Name,
				Avatar: sender.Avatar,
			}, {
				Id:     receiverId,
				Name:   receiver.Name,
				Avatar: receiver.Avatar,
			}},
			Type:   groupmdl.TypePersonal,
			Active: common.GetPointer(true),
		})
		if err != nil {
			return err
		}
	} else {
		group.Active = common.GetPointer(true)
		err = biz.groupRepository.UpdateGroup(ctx, common.GetAndFilter(
			groupstore.GetUserIdInIdListFilter(receiverId, senderId),
			groupstore.GetTypeFilter(groupmdl.TypePersonal),
		), group)
		if err != nil {
			return err
		}
	}

	go func() {
		e := biz.notification.CreateAcceptFriendNotification(ctx, senderId, &notimodel.NotificationObject{
			Id:    senderId,
			Name:  sender.Name,
			Image: &sender.Avatar,
			Type:  notimodel.User,
		}, &notimodel.NotificationObject{
			Id:    receiverId,
			Name:  receiver.Name,
			Image: &receiver.Avatar,
			Type:  notimodel.User,
		})
		if e != nil {
			log.Err(e)
		}
	}()
	return nil
}
