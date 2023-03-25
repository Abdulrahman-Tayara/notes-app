package controllers

import (
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/commands"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/injection"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/prsentation/api/context"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/prsentation/api/presenters"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func LoginController(ctx *context.Context) {

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
