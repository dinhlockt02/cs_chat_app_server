package gchatmdl

import (
	"cs_chat_app_server/common"
	groupmdl "cs_chat_app_server/modules/group/model"
)

type Group struct {
	common.MongoId `json:",inline" bson:",inline"`
	Name           string             `json:"name" bson:"name"`
	ImageUrl       *string            `json:"image_url" bson:"image_url"`
	Type           groupmdl.GroupType `json:"type"`
}

func (Group) CollectionName() string {
	return common.GroupCollectionName
}
