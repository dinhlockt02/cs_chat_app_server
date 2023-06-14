package gchatmdl

import "cs_chat_app_server/common"

type Group struct {
	common.MongoId `json:",inline" bson:",inline"`
	Name           string  `json:"name" bson:"name"`
	ImageUrl       *string `json:"image_url" bson:"image_url"`
}

func (Group) CollectionName() string {
	return common.GroupCollectionName
}
