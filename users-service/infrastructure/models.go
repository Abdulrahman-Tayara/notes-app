package infrastructure

import (
	"github.com/Abdulrahman-Tayara/notes-app/shared/core"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain/entity"
	"time"
)

type User struct {
	ID        string `gorm:"type:uuid;primary_key;"`
	Email     string
	Password  string
	CreatedAt time.Time
}

func From(entity *entity.User) *User {
	return &User{
		ID:        entity.Id.String(),
		Email:     string(entity.Email),
		Password:  entity.Password,
		CreatedAt: entity.CreatedAt,
	}

}

func (u User) To() *entity.User {
	return &entity.User{
		Id:        core.Parse(u.ID),
		Email:     domain.Email(u.Email),
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
	}
}
