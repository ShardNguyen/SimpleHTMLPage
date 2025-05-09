package consts

import "errors"

// Error messages for http messages
const (
	CreateFailed      = "failed to create new record"
	InputInvalid      = "invalid input"
	PasswordInvalid   = "invalid password"
	TokenGetFailed    = "failed to get token"
	UsernameInvalid   = "invalid username"
	UsernameGetFailed = "failed to get username"
	UsernameExisted   = "username is already taken"
)

// Defining errors
var (
	ErrUsernameExisted = errors.New(UsernameExisted)
)
