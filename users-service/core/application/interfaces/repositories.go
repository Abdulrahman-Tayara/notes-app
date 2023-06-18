package interfaces

import (
	"github.com/Abdulrahman-Tayara/notes-app/pkg/persistence"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/auth"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain/entity"
)

type (
	UsersFilter struct {
		Email string
	}

	IUserReadRepository interface {
		persistence.IReadRepository[entity.User, UsersFilter]
	}

	IUserWriteRepository interface {
		persistence.IWriteRepository[entity.User]
	}

	IRefreshTokenRepository interface {
		Save(token *auth.RefreshToken) error
		GetByToken(token string) (*auth.RefreshToken, error)
		Delete(token *auth.RefreshToken) error
	}
)
