package authmodel

import (
	"cs_chat_app_server/common"
	"errors"
	"strings"
	"time"
)

type UpdatePasswordUser struct {
	Password              string `bson:"password" json:"password"`
	common.MongoUpdatedAt `bson:",inline"`
}

func (UpdatePasswordUser) CollectionName() string {
	return "users"
}

func (u *UpdatePasswordUser) Process() error {
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
