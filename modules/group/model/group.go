package groupmdl

import (
	"cs_chat_app_server/common"
	"errors"
)

type GroupType string

const (
	TypePersonal GroupType = "personal"
	TypeGroup              = "group"
)

type Group struct {
	common.MongoId `json:",inline,omitempty" bson:",inline,omitempty"`
	Name           string    `bson:"name" json:"name"`
	Members        []string  `bson:"members,omitempty" json:"members,omitempty"`
	ImageUrl       *string   `json:"image_url" bson:"image_url"`
	Type           GroupType `bson:"type" json:"type"`
	Active         *bool     `json:"active,omitempty" bson:"active,omitempty"`
}

func (Group) CollectionName() string {
	return common.GroupCollectionName
}

func (g *Group) Process() error {
	errs := common.ValidationError{}
	if !common.URLRegexp.Match([]byte(*g.ImageUrl)) {
		errs = append(errs, errors.New("invalid group image url"))
	}
	if len(g.Name) <= 0 {
		errs = append(errs, errors.New("invalid group name"))
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}
