package pchatmdl

import "cs_chat_app_server/common"

type User struct {
	common.MongoId `json:",inline" bson:",inline"`
	Avatar         string `json:"avatar" bson:"avatar"`
	Email          string `json:"email" bson:"email"`
	Name           string `json:"name" bson:"name"`
}

func (User) CollectionName() string {
	return "users"
}
