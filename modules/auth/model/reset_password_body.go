package authmodel

import (
	"cs_chat_app_server/common"
	"errors"
	"strings"
	"time"
)

type ResetPasswordBody struct {
	Password  string     `bson:"password" json:"password"`
	Code      string     `json:"code"`
	UpdatedAt *time.Time `bson:"updated_at"`
}

func (*ResetPasswordBody) CollectionName() string {
	return common.UserCollectionName
}

func (u *ResetPasswordBody) Process() error {
	var errs = make([]error, 0)

	if len(strings.TrimSpace(u.Password)) < 8 {
		errs = append(errs, errors.New("password must be at least 8 character"))
	}

	if len(strings.TrimSpace(u.Password)) > 50 {
		errs = append(errs, errors.New("password must be at most 50 character"))
	}

	now := time.Now()
	u.UpdatedAt = &now

	if len(errs) > 0 {
		return common.ValidationError(errs)
	}
	return nil
}
