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

type SendVerifyEmailAuthStore interface {
	Find(ctx context.Context, filter map[string]interface{}) (*authmodel.User, error)
}

type SendVerifyEmailRedisStore interface {
	SetVerifyEmailCode(ctx context.Context, code string, user_id string) error
}

type sendVerifyEmail struct {
	mailer     mailer.Mailer
	authstore  SendVerifyEmailAuthStore
	redisStore SendVerifyEmailRedisStore
}

func NewSendVerifyEmail(
	mailer mailer.Mailer,
	authstore SendVerifyEmailAuthStore,
	redisStore SendVerifyEmailRedisStore,
) *sendVerifyEmail {
	return &sendVerifyEmail{
		mailer:     mailer,
		authstore:  authstore,
		redisStore: redisStore,
	}
}

func (biz *sendVerifyEmail) Send(ctx context.Context, receiver_id string, isConcurrent bool) error {
	id, _ := common.ToObjectId(receiver_id)
	receiver, err := biz.authstore.Find(ctx, map[string]interface{}{
		"_id": id,
	})

	if err != nil {
		return common.ErrInternal(err)
	}

	code := biz.getCode(receiver_id)
	link := os.Getenv("VERIFY_EMAIL_URL") + code

	err = biz.redisStore.SetVerifyEmailCode(ctx, code, receiver_id)
	if err != nil {
		return err
	}

	if isConcurrent {
		go func() {
			err = biz.mailer.Send(authmodel.VerifyEmailTitle, receiver.Email, "", authmodel.VerifyEmailBody(link))
			if err != nil {
				log.Error().Msg(err.(error).Error())
			}
		}()
	} else {
		err = biz.mailer.Send(authmodel.VerifyEmailTitle, receiver.Email, "", authmodel.VerifyEmailBody(link))
		if err != nil {
			return common.ErrInternal(err)
		}
	}
	return nil
}

func (biz *sendVerifyEmail) getCode(user_id string) string {
	return fmt.Sprintf("%v:%v", time.Now().UnixNano(), user_id)
}
