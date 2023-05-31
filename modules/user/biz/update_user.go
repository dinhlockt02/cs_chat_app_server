package userbiz

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/pubsub"
	usermodel "cs_chat_app_server/modules/user/model"
)

type UpdateUserStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*usermodel.User, error)
	Update(ctx context.Context, filter map[string]interface{}, updatedUser *usermodel.UpdateUser) error
}

type updateUserBiz struct {
	updateUserStore UpdateUserStore
	ps              pubsub.PubSub
}

func NewUpdateUserBiz(
	updateUserStore UpdateUserStore,
	ps pubsub.PubSub,
) *updateUserBiz {
	return &updateUserBiz{
		updateUserStore: updateUserStore,
		ps:              ps,
	}
}

func (biz *updateUserBiz) Update(ctx context.Context, filter map[string]interface{}, data *usermodel.UpdateUser) error {

	if err := data.Process(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	existedUser, err := biz.updateUserStore.Find(ctx, filter)
	if err != nil {
		return err
	}

	if existedUser == nil {
		return common.ErrEntityNotFound(data.EntityName(), usermodel.ErrUserNotFound)
	}
	err = biz.updateUserStore.Update(ctx, filter, data)
	if err != nil {
		return err
	}

	_ = biz.ps.Publish(ctx, common.TopicUserUpdateProfile, *existedUser.Id)
	return nil
}
