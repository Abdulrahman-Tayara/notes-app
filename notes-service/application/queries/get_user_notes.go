package queries

import (
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/domain"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/context"
)

type GetUserNotes struct {
	Search string
}

type GetUserNotesHandler struct {
	repo interfaces.INoteReadRepository
}

func NewGetUserNotesHandler(repo interfaces.INoteReadRepository) *GetUserNotesHandler {
	return &GetUserNotesHandler{repo: repo}
}

func (h *GetUserNotesHandler) Handle(ctx *context.AppContext, request *GetUserNotes) (domain.Notes, error) {
	notes, err := h.repo.GetAll(interfaces.NoteSpecification{
		UserId: ctx.UserId,
		Title:  request.Search,
	})

	if err != nil {
		return nil, err
	}

	return notes, nil
}
