package commands

import (
	"context"
	"errors"
	errors2 "github.com/Abdulrahman-Tayara/notes-app/shared/errors"
	"github.com/Abdulrahman-Tayara/notes-app/shared/helpers"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/ports"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/services"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/types"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain/entity"
	"time"
)

type Login struct {
	Email    string
	Password string
}

type LoginResult struct {
	AccessToken  services.Token `json:"access_token"`
	RefreshToken services.Token `json:"refresh_token"`
}

type LoginHandler struct {
	options                types.AuthOptions
	userRepository         interfaces.IUserReadRepository
	refreshTokenRepository interfaces.IRefreshTokenRepository
	tokenService           services.ITokenService
	hashService            services.IHashService
}

func NewLoginHandler(
	options types.AuthOptions,
	userRepository interfaces.IUserReadRepository,
	refreshTokenRepository interfaces.IRefreshTokenRepository,
	tokenService services.ITokenService,
	hashService services.IHashService,
) *LoginHandler {
	return &LoginHandler{
		options:                options,
		userRepository:         userRepository,
		refreshTokenRepository: refreshTokenRepository,
		tokenService:           tokenService,
		hashService:            hashService,
	}
}

func (h *LoginHandler) Handle(ctx context.Context, request Login, outputPort ports.IOutputPort[*LoginResult]) {
	password := h.hashService.HashString(request.Password)
	user, err := h.userRepository.GetOne(&entity.User{Email: domain.Email(request.Email), Password: password})

	if err != nil && errors.Is(errors2.ErrEntityNotFound, err) {
		outputPort.HandleError(types.InvalidCredentialsException)
		return
	} else if err != nil {
		outputPort.HandleError(err)
		return
	}

	accessToken, err := h.tokenService.Generate(&services.GenerateInput{
		Payload: types.UserClaimsPayload{
			UserId: user.Id.String(),
			Email:  string(user.Email),
		}.AsPayload(),
		ExpiresIn: time.Now().Add(h.options.AccessTokenAge),
	})

	if err != nil {
		outputPort.HandleError(err)
		return
	}

	refreshToken := types.NewRefreshToken(
		helpers.GenerateRandomString(80),
		user.Id,
		time.Now().Add(h.options.RefreshTokenAge),
	)

	err = h.refreshTokenRepository.Save(refreshToken)
	if err != nil {
		outputPort.HandleError(err)
		return
	}

	outputPort.HandleResult(&LoginResult{
		AccessToken:  accessToken,
		RefreshToken: services.Token(refreshToken.Token),
	})
}
