package groupmdl

import (
	"cs_chat_app_server/common"
	"errors"
)

type UpdateGroup struct {
	Name          *string       `bson:"name,omitempty" json:"name,omitempty"`
	ImageUrl      *string       `json:"image_url,omitempty" bson:"image_url,omitempty"`
	Active        *bool         `json:"active,omitempty" bson:"active,omitempty"`
	LatestMessage *GroupMessage `json:"latest_message,omitempty" bson:"latest_message,omitempty"`
}

func (UpdateGroup) CollectionName() string {
	return common.GroupCollectionName
}

func (g *UpdateGroup) Process() error {
	errs := common.ValidationError{}
	if !common.URLRegexp.Match([]byte(*g.ImageUrl)) {
		errs = append(errs, errors.New("invalid group image url"))
	}
	if g.Name != nil && len(*g.Name) <= 0 {
		errs = append(errs, errors.New("invalid group name"))
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}
