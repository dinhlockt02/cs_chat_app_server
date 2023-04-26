package authbiz

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/hasher"
	"cs_chat_app_server/components/tokenprovider"
	authmodel "cs_chat_app_server/modules/auth/model"
	devicemodel "cs_chat_app_server/modules/device/model"
	"strings"
	"time"
)

type LoginAuthStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*authmodel.User, error)
}

type LoginDeviceStore interface {
	Create(ctx context.Context, data *devicemodel.Device) error
}

type loginBiz struct {
	tokenProvider  tokenprovider.TokenProvider
	authStore      LoginAuthStore
	deviceStore    LoginDeviceStore
	passwordHasher hasher.Hasher
}

func NewLoginBiz(
	tokenProvider tokenprovider.TokenProvider,
	authStore LoginAuthStore,
	deviceStore LoginDeviceStore,
	passwordHasher hasher.Hasher,
) *loginBiz {
	return &loginBiz{
		tokenProvider:  tokenProvider,
		authStore:      authStore,
		passwordHasher: passwordHasher,
		deviceStore:    deviceStore,
	}
}

func (biz *loginBiz) Login(ctx context.Context, data *authmodel.LoginUser, device *devicemodel.Device) (*authmodel.AuthToken, error) {

	if err := data.Process(); err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	if err := device.Process(); err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	existedUser, err := biz.authStore.Find(ctx, map[string]interface{}{
		"email": data.Email,
	})
	if err != nil {
		return nil, err
	}
	if existedUser == nil {
		return nil, common.ErrInvalidRequest(authmodel.ErrUserNotFound)
	}

	if strings.TrimSpace(existedUser.Password) == "" {
		return nil, common.ErrInvalidRequest(authmodel.ErrPasswordNotSet)
	}

	isMatch, err := biz.passwordHasher.Compare(data.Password, existedUser.Password)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	if !isMatch {
		return nil, common.ErrInvalidRequest(authmodel.ErrEmailOrPasswordNotMatch)
	}

	device.UserId = existedUser.Id
	err = biz.deviceStore.Create(ctx, device)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	refreshToken := &tokenprovider.Token{Token: *device.Id, CreatedAt: &now, ExpiredAt: nil}
	if err != nil {
		return nil, err
	}

	accessToken, err := biz.tokenProvider.Generate(
		&tokenprovider.TokenPayload{UserId: existedUser.Id},
		common.AccessTokenExpiry,
	)

	return &authmodel.AuthToken{
		AccessToken:    accessToken,
		RefreshToken:   refreshToken,
		EmailVerified:  existedUser.EmailVerified,
		ProfileUpdated: existedUser.ProfileUpdated,
	}, nil
}
