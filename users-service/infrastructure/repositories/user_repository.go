package repositories

import (
	"github.com/Abdulrahman-Tayara/notes-app/pkg/persistence/postgres"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
	*postgres.ReadRepository[entity.User, interfaces.UsersFilter]
	*postgres.WriteRepository[entity.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db:              db,
		ReadRepository:  postgres.NewPostgresReadRepository[entity.User, interfaces.UsersFilter](db, filtersAsMap),
		WriteRepository: postgres.NewPostgresWriteRepository[entity.User](db),
	}
}

func filtersAsMap(filters interfaces.UsersFilter) postgres.Specification {
	return postgres.Equal("email", filters.Email)
}
