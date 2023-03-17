package infrastructure

import (
	"fmt"
	"github.com/Abdulrahman-Tayara/notes-app/shared/infrastructure/postgres"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	*postgres.ReadRepository[entity.User, User, interfaces.UsersFilter]
	*postgres.WriteRepository[entity.User, User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		ReadRepository:  postgres.NewPostgresReadRepository[entity.User, User, interfaces.UsersFilter](db, filtersAsMap),
		WriteRepository: postgres.NewPostgresWriteRepository[entity.User, User](db, From),
	}
}

func filtersAsMap(filters interfaces.UsersFilter) any {
	return fmt.Sprintf("email = '%s'", filters.Email)
}
