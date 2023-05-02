package authbiz

import (
	"context"
	"cs_chat_app_server/common"
	"errors"
	"strings"
)

type VerifyEmailAuthStore interface {
	UpdateEmailVerified(ctx context.Context, filter map[string]interface{}) error
}

type VerifyEmailRedisStore interface {
	GetVerifyEmailCode(ctx context.Context, code string) string
}

type verifyEmailBiz struct {
	authstore  VerifyEmailAuthStore
	redisStore VerifyEmailRedisStore
}

func NewVerifyEmail(
	authstore VerifyEmailAuthStore,
	redisStore VerifyEmailRedisStore,
) *verifyEmailBiz {
	return &verifyEmailBiz{
		authstore:  authstore,
		redisStore: redisStore,
	}
}

func (biz *verifyEmailBiz) Verify(ctx context.Context, code string) error {
	user_id := biz.redisStore.GetVerifyEmailCode(ctx, code)
	if strings.TrimSpace(user_id) == "" {
		return common.ErrInvalidRequest(errors.New("invalid code"))
	}
	id, err := common.ToObjectId(user_id)
	if err != nil {
		return common.ErrInternal(err)
	}

	err = biz.authstore.UpdateEmailVerified(ctx, map[string]interface{}{
		"_id": id,
	})
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
