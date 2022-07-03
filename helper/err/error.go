package err

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")
	ErrUnauthorized   = errors.New("error unauthorized")
	ErrNotFound       = errors.New("data not found")
	ErrIDEmpty        = errors.New("id is empty")
	ErrEmailEmpty     = errors.New("email is empty")
	ErrNameEmpty      = errors.New("name is empty")
	ErrPasswordEmpty  = errors.New("password is empty")
	ErrNoUser         = errors.New("there is no such user")
)
