package userbiz

import (
	"context"
	usermodel "cs_chat_app_server/modules/user/model"
)

type FindUserUserRepository interface {
	FindUser(ctx context.Context, requesterId string, filter map[string]interface{}) (*usermodel.User, error)
}
type findUserBiz struct {
	userRepo FindUserUserRepository
}

func NewFindUserBiz(
	userRepo FindUserUserRepository,
) *findUserBiz {
	return &findUserBiz{
		userRepo: userRepo,
	}
}

func (biz *findUserBiz) FindUser(ctx context.Context, requesterId string, filter map[string]interface{}) (*usermodel.User, error) {

	return biz.userRepo.FindUser(ctx, requesterId, filter)
}
