package repositories

import (
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/domain"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/persistence/postgres"
	"gorm.io/gorm"
)

type NoteRepository struct {
	db *gorm.DB
	*postgres.ReadRepository[domain.Note, interfaces.NoteSpecification]
	*postgres.WriteRepository[domain.Note]
}

func NewNoteRepository(db *gorm.DB) *NoteRepository {
	return &NoteRepository{
		db:              db,
		ReadRepository:  postgres.NewPostgresReadRepository[domain.Note, interfaces.NoteSpecification](db, nil),
		WriteRepository: postgres.NewPostgresWriteRepository[domain.Note](db),
	}
}
