package errors

import "errors"

var (
	ErrorUsernameRequired = errors.New("username is required")
	ErrorEmailRequired    = errors.New("email is required")
)
