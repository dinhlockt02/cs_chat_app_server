package requestmdl

import (
	"cs_chat_app_server/common"
)

type UpdateRequest struct {
	Sender   *RequestUser  `json:"sender,omitempty" bson:"sender,omitempty"`
	Receiver *RequestUser  `json:"receiver,omitempty" bson:"receiver,omitempty"`
	Group    *RequestGroup `bson:"group,omitempty" json:"group,omitempty"`
}

func (UpdateRequest) CollectionName() string {
	return common.RequestCollectionName
}
