package pchatmdl

import (
	"cs_chat_app_server/common"
	"time"
)

type MessageType string

const (
	image MessageType = "image"
	text  MessageType = "text"
	video MessageType = "video"
)

type PersonalChatItem struct {
	common.MongoId        `bson:",inline" json:",inline,omitempty"`
	Type                  MessageType `json:"type" bson:"type"`
	Sender                string      `bson:"sender" json:"sender"`
	Receiver              string      `json:"receiver" bson:"receiver"`
	Message               string      `bson:"message" json:"message"`
	Optional              *string     `bson:"optional,omitempty" json:"optional,omitempty"`
	common.MongoCreatedAt `bson:",inline" json:",inline"`
}

func (p *PersonalChatItem) CollectionName() string {
	return "personalChatHistory"
}

func (p *PersonalChatItem) Process() error {
	now := time.Now()
	p.CreatedAt = &now
	p.Id = nil
	return nil
}
