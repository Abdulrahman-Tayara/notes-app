package commands

import (
	"context"
	"errors"
	errors2 "github.com/Abdulrahman-Tayara/notes-app/pkg/errors"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/auth"
	interfaces2 "github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/ports"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain/entity"
	interfaces3 "github.com/Abdulrahman-Tayara/notes-app/users-service/mocks/application/interfaces"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
	"time"
)

func TestLoginHandler_Handle(t *testing.T) {
	tests := map[string]struct {
		input  Login
		err    error
		result *LoginResult

		mockFunc func(hashService *interfaces3.IHashService, tokenService *interfaces3.ITokenService,
			userRepository *interfaces3.IUserReadRepository, refreshTokenRepo *interfaces3.IRefreshTokenRepository)
	}{
		"Invalid creds": {
			input:  Login{Email: "test@email.com", Password: "12345"},
			err:    domain.InvalidCredentialsException,
			result: nil,
			mockFunc: func(hashService *interfaces3.IHashService, tokenService *interfaces3.ITokenService, userRepository *interfaces3.IUserReadRepository, refreshTokenRepo *interfaces3.IRefreshTokenRepository) {
				filter := &entity.User{Email: domain.Email("test@email.com"), Password: "hashedString"}
				hashService.On("HashString", mock.Anything).Return("hashedString")
				userRepository.On("GetOne", filter).
					Return(nil, errors2.ErrEntityNotFound)
			},
		},
		"Success login": {
			input:  Login{Email: "test@email.com", Password: "12345"},
			err:    nil,
			result: &LoginResult{AccessToken: "access", RefreshToken: "refresh"},
			mockFunc: func(hashService *interfaces3.IHashService, tokenService *interfaces3.ITokenService, userRepository *interfaces3.IUserReadRepository, refreshTokenRepo *interfaces3.IRefreshTokenRepository) {
				user := entity.User{
					Id:       "user-id",
					Email:    "test@email.com",
					Password: "hashedString",
				}

				filter := &entity.User{Email: domain.Email("test@email.com"), Password: "hashedString"}

				hashService.On("HashString", mock.Anything).Return("hashedString")
				userRepository.On("GetOne", filter).
					Return(&user, nil)

				tokenService.On("Generate", mock.Anything).Return(interfaces2.Token("access"), nil)
				refreshTokenRepo.On("Save", mock.Anything).Return(nil)
			},
		},
		"Save refresh token error": {
			input:  Login{Email: "test@email.com", Password: "12345"},
			err:    errors.New("some db error"),
			result: nil,
			mockFunc: func(hashService *interfaces3.IHashService, tokenService *interfaces3.ITokenService, userRepository *interfaces3.IUserReadRepository, refreshTokenRepo *interfaces3.IRefreshTokenRepository) {
				user := entity.User{
					Id:       "user-id",
					Email:    "test@email.com",
					Password: "hashedString",
				}
				filter := &entity.User{Email: domain.Email("test@email.com"), Password: "hashedString"}

				hashService.On("HashString", mock.Anything).Return("hashedString")
				userRepository.On("GetOne", filter).
					Return(&user, nil)
				tokenService.On("Generate", mock.Anything).Return(interfaces2.Token("access"), nil)
				refreshTokenRepo.On("Save", mock.Anything).Return(errors.New("some db error"))
			},
		},
	}

	options := auth.AuthOptions{
		AccessTokenAge:  time.Hour,
		RefreshTokenAge: time.Hour,
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			hashService := interfaces3.NewIHashService(t)
			tokenService := interfaces3.NewITokenService(t)
			refreshTokenRepo := interfaces3.NewIRefreshTokenRepository(t)
			userRepo := interfaces3.NewIUserReadRepository(t)

			handler := NewLoginHandler(options, userRepo, refreshTokenRepo, tokenService, hashService)

			test.mockFunc(hashService, tokenService, userRepo, refreshTokenRepo)

			output := ports.MockOutputPort[*LoginResult]{}

			handler.Handle(context.TODO(), test.input, &output)

			if test.err != nil {
				if !reflect.DeepEqual(test.err, output.Err) {
					t.Errorf("expected error: %e, actual error: %e", test.err, output.Err)
					return
				}
			}

		})
	}
}
