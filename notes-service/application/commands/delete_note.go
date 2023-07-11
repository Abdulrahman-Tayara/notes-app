package commands

import (
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/domain"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/context"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/core"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/errors"
)

type DeleteNote struct {
	NoteId core.ID
}

type DeleteNoteHandler struct {
	rRepo interfaces.INoteReadRepository
	wRepo interfaces.INoteWriteRepository
}

func NewDeleteNoteHandler(rRepo interfaces.INoteReadRepository, wRepo interfaces.INoteWriteRepository) *DeleteNoteHandler {
	return &DeleteNoteHandler{rRepo: rRepo, wRepo: wRepo}
}

func (h *DeleteNoteHandler) Handle(ctx *context.AppContext, request *DeleteNote) error {
	note, err := h.rRepo.GetById(request.NoteId)

	if err != nil {
		return err
	}

	if err = h.authorize(ctx, note); err != nil {
		return err
	}

	return h.wRepo.DeleteById(note.Id)
}

func (h *DeleteNoteHandler) authorize(ctx *context.AppContext, note *domain.Note) error {
	if ctx.UserId != note.UserId {
		return errors.ForbiddenExecption
	}
	return nil
}
