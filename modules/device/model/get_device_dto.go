package devicemodel

import "cs_chat_app_server/common"

type GetDeviceDto struct {
	common.MongoId        `bson:",inline" json:",inline"`
	common.MongoCreatedAt `bson:",inline" json:",inline"`
	common.MongoUpdatedAt `bson:",inline" json:",inline"`
	Name                  string `bson:"name" json:"name"`
	UserId                string `bson:"user_id,omitempty" json:"-"`
	PushNotificationToken string `json:"-" bson:"push_notification_token"`
}

func (GetDeviceDto) CollectionName() string {
	return "devices"
}
