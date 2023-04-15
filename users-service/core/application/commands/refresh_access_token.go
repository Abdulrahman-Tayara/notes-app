package commands

import (
	"context"
	"github.com/Abdulrahman-Tayara/notes-app/shared/errors"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/ports"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/services"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/types"
	"time"
)

type RefreshAccessToken struct {
	RefreshToken string
}

type RefreshAccessTokenResult struct {
	AccessToken string `json:"access_token"`
}

type RefreshAccessTokenHandler struct {
	authOptions      types.AuthOptions
	userRepo         interfaces.IUserReadRepository
	refreshTokenRepo interfaces.IRefreshTokenRepository
	tokenService     services.ITokenService
}

func NewRefreshAccessTokenHandler(authOptions types.AuthOptions, userRepo interfaces.IUserReadRepository, refreshTokenRepo interfaces.IRefreshTokenRepository,
	tokenService services.ITokenService) *RefreshAccessTokenHandler {
	return &RefreshAccessTokenHandler{
		authOptions: authOptions,
		userRepo:    userRepo, refreshTokenRepo: refreshTokenRepo, tokenService: tokenService}
}

func (h *RefreshAccessTokenHandler) Handle(ctx context.Context, request RefreshAccessToken, output ports.IOutputPort[*RefreshAccessTokenResult]) {
	refreshToken, err := h.refreshTokenRepo.GetByToken(request.RefreshToken)

	defer func() {
		if err := recover(); err != nil {
			output.HandleError(err.(error))
			return
		}
	}()

	if err != nil {
		panic(err)
	}

	if refreshToken.ExpiresIn.UTC().Before(time.Now().UTC()) {
		panic(errors.UnauthorizedException)
	}

	user, err := h.userRepo.GetById(refreshToken.UserId)

	if err != nil {
		panic(err)
	}

	token, err := h.tokenService.Generate(&services.GenerateInput{
		Payload: types.UserClaimsPayload{
			UserId: user.Id.String(),
			Email:  string(user.Email),
		}.AsPayload(),
		ExpiresIn: time.Now().Add(h.authOptions.AccessTokenAge).Truncate(time.Second),
	})

	if err != nil {
		panic(err)
	}

	output.HandleResult(&RefreshAccessTokenResult{
		AccessToken: string(token),
	})
}
