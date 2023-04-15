package commands

import (
	"context"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/ports"
)

type Logout struct {
	RefreshToken string
}

type LogoutHandler struct {
	refreshTokenRepo interfaces.IRefreshTokenRepository
}

func NewLogoutHandler(refreshTokenRepo interfaces.IRefreshTokenRepository) *LogoutHandler {
	return &LogoutHandler{
		refreshTokenRepo: refreshTokenRepo,
	}
}

func (h *LogoutHandler) Handle(ctx context.Context, request Logout, output ports.IOutputPort[bool]) {
	refreshToken, err := h.refreshTokenRepo.GetByToken(request.RefreshToken)

	if err != nil {
		output.HandleError(err)
		return
	}

	err = h.refreshTokenRepo.Delete(refreshToken)
	if err != nil {
		output.HandleError(err)
		return
	}

	output.HandleResult(true)
}
