package devicemodel

import (
	"cs_chat_app_server/common"
	"errors"
	"strings"
	"time"
)

type UpdateDevice struct {
	common.MongoUpdatedAt `bson:",inline" json:",inline"`
	Name                  string `bson:"name" json:"name"`
	PushNotificationToken string `json:"push_notification_token" bson:"push_notification_token"`
}

func (d *UpdateDevice) Process() error {
	var errs = make([]error, 0)

	if len(strings.TrimSpace(d.Name)) == 0 {
		errs = append(errs, errors.New("device name must not be empty"))
	}

	if len(errs) > 0 {
		return common.ValidationError(errs)
	}

	now := time.Now()
	d.UpdatedAt = &now

	return nil
}

func (UpdateDevice) CollectionName() string {
	return "devices"
}
