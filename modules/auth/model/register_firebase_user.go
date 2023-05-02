package authmodel

import (
	"cs_chat_app_server/common"
	"errors"
	"time"
)

type RegisterFirebaseUser struct {
	common.MongoId        `bson:",inline,omitempty"`
	common.MongoUpdatedAt `bson:",inline" json:",inline"`
	common.MongoCreatedAt `bson:",inline" json:",inline"`
	Email                 string `json:"email" bson:"email"`
	EmailVerified         bool   `json:"email_verified" bson:"email_verified"`
	ProfileUpdated        bool   `json:"profile_updated" bson:"profile_updated"`
}

func (RegisterFirebaseUser) CollectionName() string {
	return "users"
}

func (u *RegisterFirebaseUser) Process() error {
	var errs = make([]error, 0)

	if !common.EmailRegexp.Match([]byte(u.Email)) {
		errs = append(errs, errors.New("invalid email"))
	}

	now := time.Now()
	u.CreatedAt = &now
	u.UpdatedAt = &now

	if len(errs) > 0 {
		return common.ValidationError(errs)
	}

	u.EmailVerified = true
	u.ProfileUpdated = false

	return nil
}
