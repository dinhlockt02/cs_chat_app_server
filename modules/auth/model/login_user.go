package authmodel

import (
	"cs_chat_app_server/common"
	"errors"
	"strings"
)

type LoginUser struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" json:"password"`
}

func (LoginUser) CollectionName() string {
	return "users"
}

func (u *LoginUser) Process() error {
	var errs = make([]error, 0)

	if !common.EmailRegexp.Match([]byte(u.Email)) {
		errs = append(errs, errors.New("invalid email"))
	}

	if len(strings.TrimSpace(u.Password)) < 8 {
		errs = append(errs, errors.New("password must be at least 8 character"))
	}

	if len(strings.TrimSpace(u.Password)) > 50 {
		errs = append(errs, errors.New("password must be at most 50 character"))
	}

	if len(errs) > 0 {
		return common.ValidationError(errs)
	}
	return nil
}
