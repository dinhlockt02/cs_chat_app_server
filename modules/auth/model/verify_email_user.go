package authmodel

import (
	"cs_chat_app_server/common"
	"time"
)

type EmailVerifiedUser struct {
	EmailVerified         bool `json:"email_verified" bson:"email_verified"`
	common.MongoUpdatedAt `bson:",inline"`
}

func (EmailVerifiedUser) CollectionName() string {
	return "users"
}

func (u *EmailVerifiedUser) Process() {
	now := time.Now()
	u.UpdatedAt = &now
	u.EmailVerified = true
}
