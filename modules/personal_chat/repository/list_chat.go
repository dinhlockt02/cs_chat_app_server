package pchatrepo

import (
	"context"
	"cs_chat_app_server/common"
	pchatmdl "cs_chat_app_server/modules/personal_chat/model"
	"errors"
)

type ListMessagePersonalChatStore interface {
	List(
		ctx context.Context,
		filter map[string]interface{},
		paging pchatmdl.Paging,
	) ([]pchatmdl.PersonalChatItem, error)
	FindUser(ctx context.Context, filter map[string]interface{}) (*pchatmdl.User, error)
	AddIdFilter(id string, filter map[string]interface{}) error
}
type listMessageRepo struct {
	personalChatStore ListMessagePersonalChatStore
}

func NewListMessageRepo(
	personalChatStore ListMessagePersonalChatStore,
) *listMessageRepo {
	return &listMessageRepo{
		personalChatStore: personalChatStore,
	}
}

// TODO: Cache User Data

func (repo *listMessageRepo) List(ctx context.Context,
	filter map[string]interface{},
	paging pchatmdl.Paging,
) ([]pchatmdl.PersonalChatItem, error) {
	list, err := repo.personalChatStore.List(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	var users = make(map[string]*pchatmdl.User)

	for i, _ := range list {
		if _, ok := users[list[i].SenderId]; !ok {
			filter := map[string]interface{}{}
			err := repo.personalChatStore.AddIdFilter(list[i].SenderId, filter)
			if err != nil {
				return nil, err
			}
			user, err := repo.personalChatStore.FindUser(ctx, filter)
			if err != nil {
				return nil, err
			}
			if user == nil {
				return nil, common.ErrInternal(errors.New("User not found"))
			}
			users[list[i].SenderId] = user
		}

		if _, ok := users[list[i].ReceiverId]; !ok {
			filter := map[string]interface{}{}
			err := repo.personalChatStore.AddIdFilter(list[i].ReceiverId, filter)
			if err != nil {
				return nil, err
			}
			user, err := repo.personalChatStore.FindUser(ctx, filter)
			if err != nil {
				return nil, err
			}
			if user == nil {
				return nil, common.ErrInternal(errors.New("User not found"))
			}
			users[list[i].ReceiverId] = user
		}

		list[i].Sender = users[list[i].SenderId]
		list[i].Receiver = users[list[i].ReceiverId]
	}
	return list, nil
}
