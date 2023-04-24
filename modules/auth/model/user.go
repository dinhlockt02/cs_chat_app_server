package authmodel

import (
	"cs_chat_app_server/common"
	"errors"
)

type User struct {
	Id             string `json:"id" bson:"_id"`
	Email          string `json:"email" bson:"email"`
	Password       string `json:"password" bson:"password,omitempty"`
	EmailVerified  bool   `json:"email_verified" bson:"email_verified"`
	ProfileUpdated bool   `json:"profile_updated" bson:"profile_updated"`
}

func (User) CollectionName() string {
	return "users"
}

var ErrUserNotFound = common.ErrEntityNotFound("User", errors.New("user not found"))
