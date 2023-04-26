package authmiddleware

import (
	"cs_chat_app_server/common"
)

type User struct {
	common.MongoId `bson:",inline"`
}

func (User) CollectionName() string {
	return "users"
}

func (u *User) GetId() string {
	return *u.Id
}
