package authbiz

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/hasher"
	authmodel "cs_chat_app_server/modules/auth/model"
)

type UpdatePasswordAuthStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*authmodel.User, error)
	Update(ctx context.Context, filter map[string]interface{}, passwordUser *authmodel.UpdatePasswordUser) error
}

type updatePasswordBiz struct {
	authStore      UpdatePasswordAuthStore
	passwordHasher hasher.Hasher
}

func NewUpdatePasswordBiz(
	authStore UpdatePasswordAuthStore,
	passwordHasher hasher.Hasher,
) *updatePasswordBiz {
	return &updatePasswordBiz{
		authStore:      authStore,
		passwordHasher: passwordHasher,
	}
}

func (biz *updatePasswordBiz) Update(ctx context.Context, filter map[string]interface{}, data *authmodel.UpdatePasswordUser) error {
	if err := data.Process(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	existedUser, err := biz.authStore.Find(ctx, filter)
	if err != nil {
		return err
	}
	if existedUser == nil {
		return common.ErrEntityNotFound("User", authmodel.ErrUserNotFound)
	}

	if equal, err := biz.passwordHasher.Compare(data.OldPassword, existedUser.Password); err != nil {
		return common.ErrInternal(err)
	} else if !equal {
		return common.ErrForbidden(authmodel.ErrPasswordNotMatch)
	}

	hashedPassword, err := biz.passwordHasher.Hash(data.NewPassword)
	if err != nil {
		return common.ErrInternal(err)
	}

	data.NewPassword = hashedPassword

	err = biz.authStore.Update(ctx, filter, data)
	if err != nil {
		return err
	}

	return nil
}
