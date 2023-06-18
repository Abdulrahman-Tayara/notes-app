package commands

import (
	"context"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/errors"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/auth"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/ports"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/mocks/application/interfaces"
	"testing"
)

func TestLogoutHandler_Handle(t *testing.T) {
	tests := map[string]struct {
		input    Logout
		err      error
		mockFunc func(refreshTokenRepo *interfaces.IRefreshTokenRepository)
	}{
		"Success": {
			input: Logout{RefreshToken: "refresh token"},
			err:   nil,
			mockFunc: func(refreshTokenRepo *interfaces.IRefreshTokenRepository) {
				refreshToken := &auth.RefreshToken{
					Token: "refresh token",
				}

				refreshTokenRepo.On("GetByToken", refreshToken.Token).Return(refreshToken, nil)
				refreshTokenRepo.On("Delete", refreshToken).Return(nil)
			},
		},
		"Not found refresh token": {
			input: Logout{RefreshToken: "refresh token"},
			err:   errors.ErrEntityNotFound,
			mockFunc: func(refreshTokenRepo *interfaces.IRefreshTokenRepository) {
				refreshTokenRepo.On("GetByToken", "refresh token").Return(nil, errors.ErrEntityNotFound)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			refreshTokenRepo := interfaces.NewIRefreshTokenRepository(t)

			handler := NewLogoutHandler(refreshTokenRepo)

			output := ports.MockOutputPort[bool]{}

			test.mockFunc(refreshTokenRepo)

			handler.Handle(context.TODO(), test.input, &output)

			if output.Err != test.err {
				t.Errorf("expected error: %e, actual: %e", test.err, output.Err)
			}
		})
	}

}
