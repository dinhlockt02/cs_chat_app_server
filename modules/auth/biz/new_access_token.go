package authbiz

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/tokenprovider"
)

type accessTokenBiz struct {
	tokenProvider tokenprovider.TokenProvider
}

func NewAccessTokenBiz(
	tokenprovider tokenprovider.TokenProvider,
) *accessTokenBiz {
	return &accessTokenBiz{
		tokenProvider: tokenprovider,
	}
}

func (biz *accessTokenBiz) New(ctx context.Context, refreshToken string) (*tokenprovider.Token, error) {

	accessToken, err := biz.tokenProvider.Generate(
		&tokenprovider.TokenPayload{Id: refreshToken},
		common.AccessTokenExpiry,
	)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
