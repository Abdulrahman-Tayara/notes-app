package entity

import (
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain"
	"time"

	"github.com/Abdulrahman-Tayara/notes-app/shared/core"
)

type User struct {
	Id        core.ID
	Email     domain.Email `json:"email"`
	Password  string       `json:"password"`
	CreatedAt time.Time    `json:"created_at"`
}

func NewUser(email string, password string) (*User, error) {
	emailObject, err := domain.NewEmail(email)

	if err != nil {
		return nil, err
	}

	user := User{
		Id:        core.NewID(),
		Email:     *emailObject,
		Password:  password,
		CreatedAt: time.Now(),
	}

	return &user, nil
}

func (u User) GetId() core.ID {
	return u.Id
}

func (u *User) HashPassword(hasher func(s string) string) {
	u.Password = hasher(u.Password)
}
