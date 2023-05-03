package usermodel

import (
	"cs_chat_app_server/common"
	"errors"
	"time"
)

type User struct {
	common.MongoId        `json:",inline" bson:",inline,omitempty"`
	common.MongoUpdatedAt `json:",inline" bson:",inline,omitempty"`
	Name                  string     `json:"name" bson:"name"`
	Email                 string     `json:"email" bson:"email"`
	Password              string     `bson:"password" json:"-"`
	Avatar                string     `json:"avatar" bson:"avatar"`
	Address               string     `bson:"address" json:"address"`
	Phone                 string     `json:"phone" bson:"phone"`
	Gender                string     `json:"gender" bson:"gender"`
	Birthday              *time.Time `json:"birthday" bson:"birthday"`
}

func (User) EntityName() string {
	return "User"
}

func (User) CollectionName() string {
	return "users"
}

func (u *User) Process() {
	now := time.Now()
	u.UpdatedAt = &now
}

var ErrUserNotFound = errors.New("user not found")
var ErrUserBeBlocked = errors.New("user has been blocked")
