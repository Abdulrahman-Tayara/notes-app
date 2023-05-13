package viewmodels

import (
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain/entity"
	"time"
)

type UserViewModel struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func UserToViewModel(user *entity.User) *UserViewModel {
	return &UserViewModel{
		Id:        user.Id.String(),
		Name:      user.Name,
		Email:     string(user.Email),
		CreatedAt: user.CreatedAt,
	}
}
