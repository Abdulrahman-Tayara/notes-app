package infrastructure

import (
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/infrastructure/repositories"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/persistence/postgres"
	"gorm.io/gorm"
)

type RepositoriesStore struct {
	db *gorm.DB
}

func (r RepositoriesStore) NotesRead() interfaces.INoteReadRepository {
	return repositories.NewNoteRepository(r.db)
}

func (r RepositoriesStore) NotesWrite() interfaces.INoteWriteRepository {
	return repositories.NewNoteRepository(r.db)
}

type storeFactory struct {
}

func (f storeFactory) Create(db *gorm.DB) interfaces.IRepositoriesStore {
	c := RepositoriesStore{db: db}

	return &c
}

func NewAppUnitOfWork(db *gorm.DB) interfaces.IUnitOfWork {
	return postgres.NewPostgresUnitOfWork[interfaces.IRepositoriesStore](db, storeFactory{})
}
