package authbiz

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/hasher"
	"cs_chat_app_server/components/mailer"
	authmodel "cs_chat_app_server/modules/auth/model"
)

type ResetPasswordAuthStore interface {
	ResetPassword(ctx context.Context, filter map[string]interface{}, data *authmodel.ResetPasswordBody) error
}

type ResetPasswordRedisStore interface {
	GetForgetPasswordEmail(ctx context.Context, code string) (string, error)
}

type resetPasswordBiz struct {
	mailer         mailer.Mailer
	authstore      ResetPasswordAuthStore
	redisStore     ResetPasswordRedisStore
	passwordHasher hasher.Hasher
}

func NewResetPasswordBiz(
	authstore ResetPasswordAuthStore,
	redisStore ResetPasswordRedisStore,
	passwordHasher hasher.Hasher,
) *resetPasswordBiz {
	return &resetPasswordBiz{
		authstore:      authstore,
		redisStore:     redisStore,
		passwordHasher: passwordHasher,
	}
}

func (biz *resetPasswordBiz) Execute(ctx context.Context, data *authmodel.ResetPasswordBody) error {
	if err := data.Process(); err != nil {
		return err
	}

	email, err := biz.redisStore.GetForgetPasswordEmail(ctx, data.Code)
	if err != nil {
		return common.ErrInternal(err)
	}

	if m := common.EmailRegexp.Match([]byte(email)); !m {
		return common.ErrInvalidRequest(authmodel.ErrInvalidCode)
	}

	hashedPassword, err := biz.passwordHasher.Hash(data.Password)
	if err != nil {
		return common.ErrInternal(err)
	}
	data.Password = hashedPassword

	return biz.authstore.ResetPassword(ctx, map[string]interface{}{
		"email": email,
	}, data)

}
