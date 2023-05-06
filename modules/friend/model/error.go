package friendmodel

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrUserBeBlocked = errors.New("user is blocked")

var ErrRequestExists = errors.New("request exists")
var ErrRequestNotFound = errors.New("request not found")
var ErrHasBeenFriend = errors.New("has been friend")
