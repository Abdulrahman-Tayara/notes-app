package interfaces

import (
	"github.com/Abdulrahman-Tayara/notes-app/shared/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/auth"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain/entity"
)

type (
	UsersFilter struct {
		Email string
	}

	IUserReadRepository interface {
		interfaces.IReadRepository[entity.User, UsersFilter]
	}

	IUserWriteRepository interface {
		interfaces.IWriteRepository[entity.User]
	}

	IRefreshTokenRepository interface {
		Save(token *auth.RefreshToken) error
		GetByToken(token string) (*auth.RefreshToken, error)
		Delete(token *auth.RefreshToken) error
	}
)
