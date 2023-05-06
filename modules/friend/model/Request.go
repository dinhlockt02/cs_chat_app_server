package friendmodel

import (
	"cs_chat_app_server/common"
	"time"
)

type RequestUser struct {
	Id     string `bson:"id" json:"id"`
	Name   string `bson:"name" json:"name"`
	Avatar string `json:"avatar" bson:"avatar"`
}

type Request struct {
	Id                    *string     `json:"-" bson:"_id,omitempty"`
	Sender                RequestUser `json:"sender" bson:"sender"`
	Receiver              RequestUser `json:"receiver" bson:"receiver"`
	common.MongoCreatedAt `bson:",inline" json:",inline"`
}

func (Request) CollectionName() string {
	return "requests"
}

func (r *Request) Process() {
	now := time.Now()
	r.CreatedAt = &now
}
