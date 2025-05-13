package consts

import "errors"

// Error messages for http messages
const (
	CreateFailed      = "failed to create new record"
	InputInvalid      = "invalid input"
	OrmNotExist       = "ORM not exist"
	PasswordInvalid   = "invalid password"
	TokenGetFailed    = "failed to get token"
	TokenInvalid      = "invalid token"
	UsernameInvalid   = "invalid username"
	UsernameGetFailed = "failed to get username"
	UsernameExisted   = "username is already taken"
)

// Defining errors
var (
	ErrInputInvalid    = errors.New(InputInvalid)
	ErrOrmNotExist     = errors.New(OrmNotExist)
	ErrTokenInvalid    = errors.New(TokenInvalid)
	ErrUsernameExisted = errors.New(UsernameExisted)
)
