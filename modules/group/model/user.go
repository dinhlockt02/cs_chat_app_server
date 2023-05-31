package groupmdl

import "cs_chat_app_server/common"

type User struct {
	common.MongoId `json:",inline" bson:",inline"`
	Groups         []string `bson:"groups"`
	Avatar         string   `bson:"avatar" json:"avatar"`
	Name           string   `json:"name" bson:"name"`
}

func (User) CollectionName() string {
	return common.UserCollectionName
}
