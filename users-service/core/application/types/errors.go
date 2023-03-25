package types

import (
	sharederrors "github.com/Abdulrahman-Tayara/notes-app/shared/errors"
)

var (
	InvalidCredentialsException = sharederrors.NewException("invalid credentials", 401)
)
