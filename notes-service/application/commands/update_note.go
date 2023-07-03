package commands

import (
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/domain"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/context"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/core"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/errors"
)

type UpdateNote struct {
	NoteId  core.ID
	Title   string
	Content string
}

type UpdateNoteHandler struct {
	unitOfWork interfaces.IUnitOfWork
}

func NewUpdateNoteHandler(unitOfWork interfaces.IUnitOfWork) *UpdateNoteHandler {
	return &UpdateNoteHandler{unitOfWork: unitOfWork}
}

func (h *UpdateNoteHandler) Handle(ctx *context.AppContext, request *UpdateNote) (retNote *domain.Note, retErr error) {
	note, err := h.unitOfWork.Store().
		NotesRead().
		GetById(request.NoteId)

	if err != nil {
		return nil, err
	}

	if err = h.authorize(note, ctx); err != nil {
		return nil, err
	}

	if err = note.Update(request.Title, request.Content); err != nil {
		return nil, err
	}

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

	writeRepo := h.unitOfWork.Store().NotesWrite()

	if err = writeRepo.Update(note); err != nil {
		panic(err)
	}

	if err = h.unitOfWork.Commit(); err != nil {
		panic(err)
	}

	return note, nil
}

func (h *UpdateNoteHandler) authorize(note *domain.Note, ctx *context.AppContext) error {
	if ctx.UserId != note.UserId {
		return errors.ForbiddenExecption
	}

	return nil
}
