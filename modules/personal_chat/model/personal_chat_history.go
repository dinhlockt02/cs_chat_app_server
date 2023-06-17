package pchatmdl

import (
	"cs_chat_app_server/common"
	"time"
)

type MessageType string

const (
	image  MessageType = "image"
	text               = "text"
	video              = "video"
	record             = "record"
)

type PersonalChatItem struct {
	common.MongoId        `bson:",inline" json:",inline,omitempty"`
	Type                  MessageType `json:"type" bson:"type"`
	SenderId              string      `bson:"sender" json:"-"`
	ReceiverId            string      `json:"receiver_id" bson:"receiver"`
	Sender                *User       `bson:"-" json:"sender"`
	Receiver              *User       `json:"receiver" bson:"-"`
	Message               string      `bson:"message" json:"message"`
	Optional              *string     `bson:"optional,omitempty" json:"optional,omitempty"`
	ThumbnailUrl          *string     `bson:"thumbnail_url,omitempty" json:"thumbnail_url,omitempty"`
	VideoUrl              *string     `bson:"video_url,omitempty" json:"video_url,omitempty"`
	ImageUrl              *string     `bson:"image_url,omitempty" json:"image_url,omitempty"`
	RecordUrl             *string     `bson:"record_url,omitempty" json:"record_url,omitempty"`
	common.MongoCreatedAt `bson:",inline" json:",inline"`
	IsMe                  *bool `json:"is_me,omitempty"`
}

func (p *PersonalChatItem) CollectionName() string {
	return "personalChatHistory"
}

func (p *PersonalChatItem) Process() error {
	now := time.Now()
	p.CreatedAt = &now
	p.Id = nil

	if p.Type != video {
		p.ThumbnailUrl = nil
		p.VideoUrl = nil
	}

	if p.Type != image {
		p.ImageUrl = nil
	}

	if p.Type != record {
		p.RecordUrl = nil
	}

	return nil
}
