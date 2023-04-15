package commands

import (
	"context"
	"github.com/Abdulrahman-Tayara/notes-app/shared/core"
	"github.com/Abdulrahman-Tayara/notes-app/shared/errors"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/ports"
	services2 "github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/services"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/types"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain/entity"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/internal/mocks/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/internal/mocks/application/services"
	"testing"
	"time"
)

func TestRefreshAccessTokenHandler_Handle(t *testing.T) {
	tests := map[string]struct {
		input RefreshAccessToken
		err   error

		mockFunc func(userRepo *interfaces.IUserReadRepository, refreshTokenRepo *interfaces.IRefreshTokenRepository, tokenService *services.ITokenService)
	}{
		"Success": {
			input: RefreshAccessToken{RefreshToken: "some-string"},
			err:   nil,
			mockFunc: func(userRepo *interfaces.IUserReadRepository, refreshTokenRepo *interfaces.IRefreshTokenRepository, tokenService *services.ITokenService) {
				refreshToken := &types.RefreshToken{
					Token:     "some-token",
					UserId:    core.ID("id"),
					ExpiresIn: time.Now().Add(time.Hour),
				}

				refreshTokenRepo.On("GetByToken", "some-string").Return(refreshToken, nil)

				user := entity.User{
					Id:    core.ID("id"),
					Email: "test@test.com",
				}

				userRepo.On("GetById", refreshToken.UserId).Return(&user, nil)
				tokenService.On("Generate", &services2.GenerateInput{
					Payload: types.UserClaimsPayload{
						UserId: user.Id.String(),
						Email:  string(user.Email),
					}.AsPayload(),
					ExpiresIn: time.Now().Add(time.Hour).Truncate(time.Second),
				}).Return(services2.Token("new token"), nil)
			},
		},
		"Expired refresh token": {
			input: RefreshAccessToken{RefreshToken: "some-string"},
			err:   errors.UnauthorizedException,
			mockFunc: func(userRepo *interfaces.IUserReadRepository, refreshTokenRepo *interfaces.IRefreshTokenRepository, tokenService *services.ITokenService) {
				refreshToken := &types.RefreshToken{
					Token:     "some-token",
					UserId:    core.ID("id"),
					ExpiresIn: time.Now().Add(-time.Hour * 10),
				}

				refreshTokenRepo.On("GetByToken", "some-string").Return(refreshToken, nil)
			},
		},
		"Not found refresh token": {
			input: RefreshAccessToken{RefreshToken: "some-string"},
			err:   errors.ErrEntityNotFound,
			mockFunc: func(userRepo *interfaces.IUserReadRepository, refreshTokenRepo *interfaces.IRefreshTokenRepository, tokenService *services.ITokenService) {
				refreshTokenRepo.On("GetByToken", "some-string").Return(nil, errors.ErrEntityNotFound)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			userRepo := interfaces.NewIUserReadRepository(t)
			tokenService := services.NewITokenService(t)
			refreshTokenRepo := interfaces.NewIRefreshTokenRepository(t)
			authOptions := types.AuthOptions{
				AccessTokenAge:  time.Hour,
				RefreshTokenAge: time.Hour,
			}

			handler := NewRefreshAccessTokenHandler(authOptions, userRepo, refreshTokenRepo, tokenService)

			test.mockFunc(userRepo, refreshTokenRepo, tokenService)

			output := ports.MockOutputPort[*RefreshAccessTokenResult]{}

			handler.Handle(context.TODO(), test.input, &output)

			if test.err != output.Err {
				t.Errorf("expected error: %v, actual error: %v", test.err, output.Err)
			}
		})
	}
}