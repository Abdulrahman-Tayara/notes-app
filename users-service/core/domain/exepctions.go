package domain

import (
	"github.com/Abdulrahman-Tayara/notes-app/shared/errors"
)

var EmailAlreadyExists = errors.NewException("Email already exists", 1122)
