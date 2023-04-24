package devicemodel

import (
	"cs_chat_app_server/common"
	"errors"
	"strings"
	"time"
)

type Device struct {
	Id        *string    `json:"id" bson:"_id,omitempty"`
	CreatedAt *time.Time `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time `bson:"updated_at" json:"updated_at,omitempty"`
	Name      string     `bson:"name" json:"name"`
	UserId    string     `bson:"user_id" json:"-"`
}

func (d *Device) Process() error {
	var errs = make([]error, 0)

	if len(strings.TrimSpace(d.Name)) == 0 {
		errs = append(errs, errors.New("device name must not be empty"))
	}

	if len(errs) > 0 {
		return common.ValidationError(errs)
	}

	now := time.Now()
	d.CreatedAt = &now
	d.UpdatedAt = &now

	return nil
}

func (Device) CollectionName() string {
	return "devices"
}
