package commands

import (
	"context"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/core"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/errors"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/auth"
	services2 "github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/ports"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain/entity"
	interfaces2 "github.com/Abdulrahman-Tayara/notes-app/users-service/mocks/application/interfaces"
	"testing"
	"time"
)

func TestRefreshAccessTokenHandler_Handle(t *testing.T) {
	tests := map[string]struct {
		input RefreshAccessToken
		err   error

		mockFunc func(userRepo *interfaces2.IUserReadRepository, refreshTokenRepo *interfaces2.IRefreshTokenRepository, tokenService *interfaces2.ITokenService)
	}{
		"Success": {
			input: RefreshAccessToken{RefreshToken: "some-string"},
			err:   nil,
			mockFunc: func(userRepo *interfaces2.IUserReadRepository, refreshTokenRepo *interfaces2.IRefreshTokenRepository, tokenService *interfaces2.ITokenService) {
				refreshToken := &auth.RefreshToken{
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
					Payload: auth.UserClaimsPayload{
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
			mockFunc: func(userRepo *interfaces2.IUserReadRepository, refreshTokenRepo *interfaces2.IRefreshTokenRepository, tokenService *interfaces2.ITokenService) {
				refreshToken := &auth.RefreshToken{
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
			mockFunc: func(userRepo *interfaces2.IUserReadRepository, refreshTokenRepo *interfaces2.IRefreshTokenRepository, tokenService *interfaces2.ITokenService) {
				refreshTokenRepo.On("GetByToken", "some-string").Return(nil, errors.ErrEntityNotFound)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			userRepo := interfaces2.NewIUserReadRepository(t)
			tokenService := interfaces2.NewITokenService(t)
			refreshTokenRepo := interfaces2.NewIRefreshTokenRepository(t)
			authOptions := auth.AuthOptions{
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
