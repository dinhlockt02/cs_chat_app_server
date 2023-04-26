package userbiz

import (
	"context"
	"cs_chat_app_server/common"
	usermodel "cs_chat_app_server/modules/user/model"
)

type UpdateUserStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*usermodel.User, error)
	Update(ctx context.Context, filter map[string]interface{}, updatedUser *usermodel.UpdateUser) error
}

type updateUserBiz struct {
	updateUserStore UpdateUserStore
}

func NewUpdateUserBiz(updateUserStore UpdateUserStore) *updateUserBiz {
	return &updateUserBiz{updateUserStore: updateUserStore}
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
	return nil
}
