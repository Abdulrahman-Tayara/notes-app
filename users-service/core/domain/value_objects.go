package domain

import (
	"github.com/Abdulrahman-Tayara/notes-app/shared/errors"
	"github.com/go-playground/validator/v10"
)

type Email string

func NewEmail(email string) (*Email, error) {
	err := validator.New().Var(email, "required,email")

	if err != nil {
		return nil, errors.BadValueException("email")
	}

	e := Email(email)

	return &e, nil
}
