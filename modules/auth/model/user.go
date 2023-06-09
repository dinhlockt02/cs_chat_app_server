package authmodel

import (
	"errors"
)

type User struct {
	Id             string  `json:"id" bson:"_id"`
	Email          string  `json:"email" bson:"email"`
	Password       string  `json:"password" bson:"password,omitempty"`
	EmailVerified  bool    `json:"email_verified" bson:"email_verified"`
	ProfileUpdated bool    `json:"profile_updated" bson:"profile_updated"`
	Name           *string `json:"name,omitempty" bson:"name,omitempty"`
}

func (User) CollectionName() string {
	return "users"
}

var ErrUserNotFound = errors.New("user not found")
var ErrPasswordNotSet = errors.New("password have not set")
var ErrEmailOrPasswordNotMatch = errors.New("email or password is not match")
var ErrPasswordNotMatch = errors.New("password is not match")
