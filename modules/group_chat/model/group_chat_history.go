package gchatmdl

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

type GroupChatItem struct {
	common.MongoId        `bson:",inline" json:",inline,omitempty"`
	Type                  MessageType `json:"type" bson:"type"`
	SenderId              string      `bson:"sender" json:"-"`
	Sender                *User       `json:"sender" bson:"-"`
	GroupId               string      `json:"-" bson:"group"`
	Group                 *Group      `json:"group" bson:"-"`
	Message               string      `bson:"message" json:"message"`
	Optional              *string     `bson:"optional,omitempty" json:"optional,omitempty"`
	common.MongoCreatedAt `bson:",inline" json:",inline"`
	IsMe                  *bool `json:"is_me,omitempty"`
}

func (p *GroupChatItem) CollectionName() string {
	return common.GroupChatHistoryCollectionName
}

func (p *GroupChatItem) Process() error {
	now := time.Now()
	p.CreatedAt = &now
	p.Id = nil
	return nil
}
