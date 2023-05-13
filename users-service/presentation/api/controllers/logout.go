package controllers

import (
	"github.com/Abdulrahman-Tayara/notes-app/shared/http"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/commands"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/injection"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/presentation/api/presenters"
)

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func Logout(ctx *http.Context) {
	var request LogoutRequest

	if !ctx.BindJsonOrReturnError(&request) {
		return
	}

	command := injection.InitLogoutCommand()
	presenter := presenters.NewLogoutPresenter()

	command.Handle(ctx, commands.Logout{RefreshToken: request.RefreshToken}, presenter)

	ctx.Response(presenter.Present())
}
