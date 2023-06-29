package controllers

import (
	"github.com/Abdulrahman-Tayara/notes-app/pkg/http"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/api/http/presenters"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/commands"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/injection"
)

type RefreshAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func RefreshAccessToken(ctx *http.Context) {

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
