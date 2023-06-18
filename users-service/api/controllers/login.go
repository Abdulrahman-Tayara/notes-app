package controllers

import (
	"github.com/Abdulrahman-Tayara/notes-app/pkg/http"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/api/presenters"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/commands"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/injection"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func LoginController(ctx *http.Context) {

	var request LoginRequest

	if !ctx.BindJsonOrReturnError(&request) {
		return
	}

	command := commands.Login{
		Email:    request.Email,
		Password: request.Password,
	}

	handler := injection.InitLoginCommand()
	presenter := presenters.NewLoginPresenter()
	handler.Handle(ctx, command, presenter)

	ctx.Response(presenter.Present())
}
