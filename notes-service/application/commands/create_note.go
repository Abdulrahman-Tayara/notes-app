package commands

import (
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/domain"
	context2 "github.com/Abdulrahman-Tayara/notes-app/pkg/context"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/core"
)

type CreateNote struct {
	UserId  core.ID
	Title   string
	Content string
}

type CreateNoteHandler struct {
	unitOfWork interfaces.IUnitOfWork
}

func NewCreateNoteHandler(unitOfWork interfaces.IUnitOfWork) *CreateNoteHandler {
	return &CreateNoteHandler{unitOfWork: unitOfWork}
}

func (h *CreateNoteHandler) Handle(ctx *context2.AppContext, request *CreateNote) (retNote *domain.Note, retErr error) {
	note, err := domain.NewNote(request.UserId, request.Title, request.Content)

	if err != nil {
		return nil, err
	}

	writeRepo := h.unitOfWork.Store().NotesWrite()

	if err = h.unitOfWork.Begin(); err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			if e, ok := err.(error); ok {
				_ = h.unitOfWork.Rollback()
				retErr = e
			}
		}
	}()

	note, err = writeRepo.Save(note)

	if err != nil {
		panic(err)
	}

	if err = h.unitOfWork.Commit(); err != nil {
		panic(err)
	}

	return note, nil
}
