package pchatrepo

import (
	"context"
	"cs_chat_app_server/common"
	pchatmdl "cs_chat_app_server/modules/personal_chat/model"
	"errors"
)

type CreateMessagePersonalChatStore interface {
	FindUser(ctx context.Context, filter map[string]interface{}) (*pchatmdl.User, error)
	AddIdFilter(id string, filter map[string]interface{}) error
	Create(ctx context.Context, personalChatItem *pchatmdl.PersonalChatItem) error
}
type createMessageRepo struct {
	personalChatStore CreateMessagePersonalChatStore
}

func NewCreateMessageRepo(
	personalChatStore CreateMessagePersonalChatStore,
) *createMessageRepo {
	return &createMessageRepo{
		personalChatStore: personalChatStore,
	}
}

// TODO: Cache User Data

func (repo *createMessageRepo) Create(ctx context.Context,
	personalChatItem *pchatmdl.PersonalChatItem,
) error {
	err := repo.personalChatStore.Create(ctx, personalChatItem)
	if err != nil {
		return common.ErrInternal(err)
	}
	filter := map[string]interface{}{}
	err = repo.personalChatStore.AddIdFilter(personalChatItem.SenderId, filter)
	if err != nil {
		return err
	}
	user, err := repo.personalChatStore.FindUser(ctx, filter)
	if err != nil {
		return err
	}
	if user == nil {
		return common.ErrInternal(errors.New("User not found"))
	}
	personalChatItem.Sender = user

	filter = map[string]interface{}{}
	err = repo.personalChatStore.AddIdFilter(personalChatItem.ReceiverId, filter)
	if err != nil {
		return err
	}
	user, err = repo.personalChatStore.FindUser(ctx, filter)
	if err != nil {
		return err
	}
	if user == nil {
		return common.ErrInternal(errors.New("User not found"))
	}
	personalChatItem.Receiver = user

	return nil
}
