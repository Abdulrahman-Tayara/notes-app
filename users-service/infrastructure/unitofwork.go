package infrastructure

import (
	"github.com/Abdulrahman-Tayara/notes-app/pkg/persistence/postgres"
	interfaces2 "github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/infrastructure/repositories"
	"gorm.io/gorm"
)

type RepositoriesConstructor struct {
	db *gorm.DB
}

func (r RepositoriesConstructor) UsersRead() interfaces2.IUserReadRepository {
	return repositories.NewUserRepository(r.db)
}

func (r RepositoriesConstructor) UsersWrite() interfaces2.IUserWriteRepository {
	return repositories.NewUserRepository(r.db)
}

type storeFactory struct {
}

func (f storeFactory) Create(db *gorm.DB) interfaces2.IRepositoriesConstructor {
	c := RepositoriesConstructor{db: db}

	return &c
}

func NewAppUnitOfWork(db *gorm.DB) interfaces2.IUnitOfWork {
	return postgres.NewPostgresUnitOfWork[interfaces2.IRepositoriesConstructor](db, storeFactory{})
}
