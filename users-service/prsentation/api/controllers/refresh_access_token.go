package controllers

import (
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/commands"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/injection"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/prsentation/api/context"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/prsentation/api/presenters"
)

type RefreshAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func RefreshAccessToken(ctx *context.Context) {

	var request RefreshAccessTokenRequest

	if !ctx.BindJsonOrReturnError(&request) {
		return
	}

	command := commands.RefreshAccessToken{
		RefreshToken: request.RefreshToken,
	}

	handler := injection.InitRefreshAccessTokenCommand()
	presenter := presenters.NewRefreshAccessTokenPresenter()

	handler.Handle(ctx, command, presenter)

	ctx.Response(presenter.Present())
}
