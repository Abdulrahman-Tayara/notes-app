package queries

import (
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/domain"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/context"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/core"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/errors"
)

type GetNote struct {
	NoteId core.ID
}

type GetNoteHandler struct {
	repo interfaces.INoteReadRepository
}

func NewGetNoteHandler(repo interfaces.INoteReadRepository) *GetNoteHandler {
	return &GetNoteHandler{repo: repo}
}

func (h *GetNoteHandler) Handle(ctx *context.AppContext, request *GetNote) (*domain.Note, error) {
	note, err := h.repo.GetById(request.NoteId)

	if err != nil {
		return nil, err
	}

	if err = h.authorize(ctx, note); err != nil {
		return nil, err
	}

	return note, nil
}

func (h *GetNoteHandler) authorize(ctx *context.AppContext, note *domain.Note) error {
	if ctx.UserId != note.UserId {
		return errors.ForbiddenExecption
	}

	return nil
}
