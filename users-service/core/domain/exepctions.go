package domain

import (
	"github.com/Abdulrahman-Tayara/notes-app/pkg/errors"
)

var EmailAlreadyExists = errors.NewException("Email already exists", 1122)
var InvalidCredentialsException = errors.NewException("invalid credentials", 401)
