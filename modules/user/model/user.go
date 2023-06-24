package usermodel

import (
	"cs_chat_app_server/common"
	friendmodel "cs_chat_app_server/modules/friend/model"
	"errors"
	"time"
)

type User struct {
	common.MongoId        `json:",inline" bson:",inline,omitempty"`
	common.MongoUpdatedAt `json:",inline" bson:",inline,omitempty"`
	Name                  string               `json:"name" bson:"name"`
	Email                 string               `json:"email" bson:"email"`
	Password              string               `bson:"password" json:"-"`
	Avatar                string               `json:"avatar" bson:"avatar"`
	Phone                 string               `json:"phone" bson:"phone"`
	Gender                string               `json:"gender" bson:"gender"`
	Birthday              *time.Time           `json:"birthday" bson:"birthday"`
	Bio                   string               `json:"bio" bson:"bio"`
	CommonFriend          []string             `json:"-"`
	CommonFriendCount     *int                 `json:"common_friend_count,omitempty"`
	Relation              friendmodel.Relation `json:"relation"`
	EmailVerified         *bool                `json:"email_verified,omitempty" bson:"email_verified,omitempty"`
	ProfileUpdated        *bool                `json:"profile_updated,omitempty" bson:"profile_updated,omitempty"`
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
