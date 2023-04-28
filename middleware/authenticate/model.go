package authmiddleware

import (
	"cs_chat_app_server/common"
)

type Device struct {
	common.MongoId `bson:",inline" json:",inline"`
	UserId         string `json:"user_id" bson:"user_id"`
}

func (Device) CollectionName() string {
	return "devices"
}

func (u *Device) GetId() string {
	return u.UserId
}

func (u *Device) GetDeviceId() string {
	return *u.Id
}
