package controllers

import (
	"github.com/Abdulrahman-Tayara/notes-app/shared/http"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/commands"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/injection"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/prsentation/api/presenters"
)

type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

func SignUpController(ctx *http.Context) {
	var request SignUpRequest

	if !ctx.BindJsonOrReturnError(&request) {
		return
	}

	command := commands.SignUp{
		Email:    request.Email,
		Name:     request.Name,
		Password: request.Password,
	}

	handler := injection.InitSignUpCommand()
	presenter := presenters.NewSingUpPresenter()

	handler.Handle(ctx, command, presenter)

	ctx.Response(presenter.Present())
}
