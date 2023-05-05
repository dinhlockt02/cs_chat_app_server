package usermodel

import (
	"cs_chat_app_server/common"
	"errors"
	"fmt"
	"strings"
	"time"
)

type UpdateUser struct {
	UpdatedAt      *time.Time `bson:"-" json:"updated_at,omitempty"`
	Name           *string    `json:"name,omitempty" bson:"name,omitempty"`
	Avatar         *string    `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Phone          *string    `json:"phone,omitempty" bson:"phone,omitempty"`
	Gender         *string    `json:"gender,omitempty" bson:"gender,omitempty"`
	Birthday       *time.Time `json:"birthday,omitempty" bson:"birthday,omitempty"`
	ProfileUpdated bool       `json:"-" bson:"profile_updated"`
	Bio            *string    `json:"bio,omitempty" bson:"bio,omitempty"`
}

func (UpdateUser) EntityName() string {
	return "User"
}

func (UpdateUser) CollectionName() string {
	return "users"
}

func (u *UpdateUser) Process() error {
	var errs = make([]error, 0)
	if u.Name != nil && strings.TrimSpace(*u.Name) == "" {
		errs = append(errs, errors.New("name must not be empty"))
	}

	if u.Avatar != nil && !common.URLRegexp.Match([]byte(*u.Avatar)) {
		errs = append(errs, errors.New("invalid image url"))
	}

	if u.Gender != nil {
		if gender := strings.TrimSpace(string(*u.Gender)); gender == "" {
			errs = append(errs, errors.New("gender must not be empty"))
		} else if gender != common.Male && gender != common.Female && gender != common.Other {
			errs = append(errs, errors.New(fmt.Sprintf("gender must be '%s' or '%s' or '%s'", common.Male, common.Female, common.Other)))
		}
	}

	if u.Phone != nil && strings.TrimSpace(*u.Phone) == "" {
		errs = append(errs, errors.New("phone number must not be empty"))
	}

	if u.Bio != nil && strings.TrimSpace(*u.Phone) == "" {
		errs = append(errs, errors.New("bio must not be empty"))
	}

	u.ProfileUpdated = true

	now := time.Now()
	u.UpdatedAt = &now

	if len(errs) > 0 {
		return common.ValidationError(errs)
	}
	return nil
}
