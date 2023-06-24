package authredis

import (
	"context"
	"cs_chat_app_server/common"
	"github.com/redis/go-redis/v9"
	"time"
)

const verifyPrefix = "verify-email:"
const forgetPasswordPrefix = "forget-password:"

type redisStore struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) *redisStore {
	return &redisStore{client: client}
}

func (s *redisStore) SetVerifyEmailCode(ctx context.Context, code string, user_id string) error {
	err := s.client.Set(ctx, verifyPrefix+code, user_id, 10*time.Minute).Err()
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}

func (s *redisStore) GetVerifyEmailCode(ctx context.Context, code string) string {
	val, _ := s.client.Get(ctx, verifyPrefix+code).Result()
	return val
}

func (s *redisStore) SetForgetPasswordCode(ctx context.Context, code string, email string) error {
	err := s.client.Set(ctx, forgetPasswordPrefix+code, email, 10*time.Minute).Err()
	if err != nil {
		return common.ErrInternal(err)
	}
	return nil
}

func (s *redisStore) GetForgetPasswordEmail(ctx context.Context, code string) (string, error) {
	val, err := s.client.Get(ctx, forgetPasswordPrefix+code).Result()
	return val, err
}
