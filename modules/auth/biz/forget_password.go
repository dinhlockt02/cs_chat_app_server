package authbiz

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/mailer"
	authmodel "cs_chat_app_server/modules/auth/model"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

type ForgetPasswordAuthStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*authmodel.User, error)
}

type ForgetPasswordRedisStore interface {
	SetForgetPasswordCode(ctx context.Context, code string, email string) error
}

type forgetPasswordBiz struct {
	mailer     mailer.Mailer
	authstore  ForgetPasswordAuthStore
	redisStore ForgetPasswordRedisStore
}

func NewForgetPasswordBiz(
	mailer mailer.Mailer,
	authstore ForgetPasswordAuthStore,
	redisStore ForgetPasswordRedisStore,
) *forgetPasswordBiz {
	return &forgetPasswordBiz{
		mailer:     mailer,
		authstore:  authstore,
		redisStore: redisStore,
	}
}

func (biz *forgetPasswordBiz) Execute(ctx context.Context, email string) error {
	receiver, err := biz.authstore.Find(ctx, map[string]interface{}{
		"email": email,
	})

	if receiver == nil {
		log.Debug().Msgf("user %s not found", email)
		return nil
	}

	if err != nil {
		return common.ErrInternal(err)
	}

	code := biz.getCode(email)

	link := os.Getenv("RESET_PASSWORD_URL") + code

	err = biz.redisStore.SetForgetPasswordCode(ctx, code, email)
	if err != nil {
		return err
	}

	receiverName := ""
	if receiver.Name != nil {
		receiverName = *receiver.Name
	}
	go func() {
		err = biz.mailer.Send(authmodel.ForgetPasswordEmail, receiver.Email, receiverName, authmodel.ForgetPasswordEmailBody(link))
		if err != nil {
			log.Error().Err(err).Msg("forget password biz: " + err.Error())
		}
	}()
	return nil
}

func (biz *forgetPasswordBiz) getCode(user_id string) string {
	return fmt.Sprintf("%v:%v", time.Now().UnixNano(), user_id)
}
