package interfaces

import (
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/domain"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/persistence"
)

type (
	NoteSpecification struct {
		UserId string
		Title  string
	}

	INoteReadRepository interface {
		persistence.IReadRepository[domain.Note, NoteSpecification]
	}

	INoteWriteRepository interface {
		persistence.IWriteRepository[domain.Note]
	}
)

type IRepositoriesStore interface {
	NotesRead() INoteReadRepository

	NotesWrite() INoteWriteRepository
}

type IUnitOfWork interface {
	persistence.IUnitOfWork[IRepositoriesStore]
}
