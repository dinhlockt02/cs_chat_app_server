package authbiz

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/hasher"
	"cs_chat_app_server/components/tokenprovider"
	authmodel "cs_chat_app_server/modules/auth/model"
	devicemodel "cs_chat_app_server/modules/device/model"
	"time"
)

type RegisterAuthStore interface {
	CreateRegisterUser(ctx context.Context, data *authmodel.RegisterUser) (*authmodel.User, error)
	Find(ctx context.Context, filter map[string]interface{}) (*authmodel.User, error)
}

type RegisterDeviceStore interface {
	Create(ctx context.Context, data *devicemodel.Device) error
}

type registerBiz struct {
	tokenProvider  tokenprovider.TokenProvider
	passwordHasher hasher.Hasher
	deviceStore    RegisterDeviceStore
	authStore      RegisterAuthStore
}

func NewRegisterBiz(
	tokenProvider tokenprovider.TokenProvider,
	passwordHasher hasher.Hasher,
	deviceStore RegisterDeviceStore,
	authStore RegisterAuthStore,
) *registerBiz {
	return &registerBiz{
		tokenProvider:  tokenProvider,
		passwordHasher: passwordHasher,
		deviceStore:    deviceStore,
		authStore:      authStore,
	}
}

func (biz *registerBiz) Register(ctx context.Context, data *authmodel.RegisterUser, device *devicemodel.Device) (*authmodel.AuthToken, error) {

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
	if existedUser != nil {
		return nil, common.ErrInvalidRequest(authmodel.ErrUserExists)
	}

	hashedPassword, err := biz.passwordHasher.Hash(data.Password)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	data.Password = hashedPassword

	user, err := biz.authStore.CreateRegisterUser(ctx, data)
	if err != nil {
		return nil, err
	}

	device.UserId = user.Id
	err = biz.deviceStore.Create(ctx, device)

	now := time.Now()
	refreshToken := &tokenprovider.Token{Token: *device.Id, CreatedAt: &now, ExpiredAt: nil}
	if err != nil {
		return nil, err
	}

	accessToken, err := biz.tokenProvider.Generate(
		&tokenprovider.TokenPayload{Id: *device.Id},
		common.AccessTokenExpiry,
	)

	return &authmodel.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
