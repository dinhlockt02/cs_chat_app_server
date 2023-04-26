package authmiddleware

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/tokenprovider"
	"net/http"
)

type AuthMiddlewareMongoStore interface {
	FindOne(ctx context.Context, filter map[string]interface{}) (*User, error)
}

type authMiddlewareBiz struct {
	store         AuthMiddlewareMongoStore
	tokenProvider tokenprovider.TokenProvider
}

func NewAuthMiddlewareBiz(
	store AuthMiddlewareMongoStore,
	tokenProvider tokenprovider.TokenProvider,
) *authMiddlewareBiz {
	return &authMiddlewareBiz{
		store:         store,
		tokenProvider: tokenProvider,
	}
}

func (biz *authMiddlewareBiz) Authenticate(ctx context.Context, token string) (*User, error) {
	tokenPayload, err := biz.tokenProvider.Validate(token)
	if err != nil {
		return nil, common.NewFullErrorResponse(http.StatusUnauthorized,
			err,
			"unauthorized",
			err.Error(),
			"UnauthorizedError",
		)
	}

	id, err := common.ToObjectId(tokenPayload.UserId)
	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}
	user, err := biz.store.FindOne(ctx, map[string]interface{}{
		"_id": id,
	})
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	if user == nil {
		return nil, common.NewFullErrorResponse(http.StatusUnauthorized,
			nil,
			"unauthorized",
			"user not found",
			"UnauthorizedError")
	}
	return user, nil
}
