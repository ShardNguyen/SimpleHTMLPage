package consts

import "errors"

// Error messages for http messages
const (
	AuthHeaderMissing = "missing authorization header"
	CreateFailed      = "failed to create new record"
	EmailInvalid      = "invalid email"
	InputInvalid      = "invalid input format"
	JSONMarshalFailed = "JSON Marshal failed"
	OrmNotExist       = "ORM does not exist"
	PasswordInvalid   = "invalid password"
	TokenGetFailed    = "failed to get token"
	TokenStoreFailed  = "failed to store token"
	TokenInvalid      = "invalid token"
	TokenNotExists    = "token does not exist"
	TokenExpired      = "token is expired"
	UsernameInvalid   = "invalid username"
	UsernameGetFailed = "failed to get username"
	UsernameExisted   = "username is already taken"
)

// Defining errors
var (
	ErrEmailInvalid    = errors.New(EmailInvalid)
	ErrInputInvalid    = errors.New(InputInvalid)
	ErrOrmNotExist     = errors.New(OrmNotExist)
	ErrTokenInvalid    = errors.New(TokenInvalid)
	ErrPasswordInvalid = errors.New(PasswordInvalid)
	ErrUsernameExisted = errors.New(UsernameExisted)
	ErrUsernameInvalid = errors.New(UsernameInvalid)
)
