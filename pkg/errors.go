package pkg

import "errors"

var ErrorsEnum = struct {
	ErrEmailNotMatch    error
	ErrPasswordNotMatch error
	ErrAccountNotExist  error
}{
	ErrEmailNotMatch:    errors.New("email not match"),
	ErrPasswordNotMatch: errors.New("password not match"),
	ErrAccountNotExist:  errors.New("account not exist"),
}
