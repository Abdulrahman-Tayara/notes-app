package errors

import "fmt"

type Exception struct {
	message string
	code    int
}

func NewException(message string, code int) Exception {
	return Exception{message: message, code: code}
}

func (e Exception) Error() string {
	return e.message
}

var (
	ErrEntityNotFound = NewException("not found", 404)
	BadValueException = func(field string) Exception {
		return NewException(fmt.Sprintf("bad valud for field %s", field), 222)
	}
	UnauthorizedException = NewException("unauthorized", 401)
)
