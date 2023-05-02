package authmodel

import (
	"cs_chat_app_server/common"
	"errors"
	"strings"
	"time"
)

type RegisterUser struct {
	Id             *string    `json:"id" bson:"_id,omitempty"`
	CreatedAt      *time.Time `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt      *time.Time `bson:"updated_at" json:"updated_at,omitempty"`
	Email          string     `json:"email" bson:"email"`
	Password       string     `json:"password" bson:"password"`
	EmailVerified  bool       `json:"email_verified" bson:"email_verified"`
	ProfileUpdated bool       `json:"profile_updated" bson:"profile_updated"`
}

func (RegisterUser) CollectionName() string {
	return "users"
}

func (u *RegisterUser) Process() error {
	var errs = make([]error, 0)

	if !common.EmailRegexp.Match([]byte(u.Email)) {
		errs = append(errs, errors.New("invalid mailer"))
	}

	if len(strings.TrimSpace(u.Password)) < 8 {
		errs = append(errs, errors.New("password must be at least 8 character"))
	}

	if len(strings.TrimSpace(u.Password)) > 50 {
		errs = append(errs, errors.New("password must be at most 50 character"))
	}

	now := time.Now()
	u.CreatedAt = &now
	u.UpdatedAt = &now

	if len(errs) > 0 {
		return common.ValidationError(errs)
	}
	return nil
}

var ErrUserExists = errors.New("user existed")
