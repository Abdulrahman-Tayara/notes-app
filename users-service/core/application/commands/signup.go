package commands

import (
	"context"
	interfaces2 "github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/ports"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain/entity"
)

type SignUp struct {
	Email    string
	Name     string
	Password string
}

type SingUpHandler struct {
	unitOfWork  interfaces2.IUnitOfWork
	hashService interfaces2.IHashService
}

func NewSingUpHandler(unitOfWork interfaces2.IUnitOfWork, hashService interfaces2.IHashService) *SingUpHandler {
	return &SingUpHandler{
		unitOfWork:  unitOfWork,
		hashService: hashService,
	}
}

func (h *SingUpHandler) Handle(ctx context.Context, command SignUp, outputPort ports.IOutputPort[*entity.User]) {
	user, err := entity.NewUser(command.Name, command.Email, command.Password)

	if err != nil {
		outputPort.HandleError(err)
		return
	}

	user.HashPassword(h.hashService.HashString)

	usersRepo := h.unitOfWork.Store().UsersRead()

	if h.isEmailExists(usersRepo, command.Email) {
		outputPort.HandleError(domain.EmailAlreadyExists)
		return
	}

	if err = h.unitOfWork.Begin(); err != nil {
		outputPort.HandleError(err)
		return
	}

	defer func() {
		if err := recover(); err != nil {
			if e, ok := err.(error); ok {
				_ = h.unitOfWork.Rollback()
				outputPort.HandleError(e)
			}
		}
	}()

	usersWriteRepo := h.unitOfWork.Store().UsersWrite()

	if _, err = usersWriteRepo.Save(user); err != nil {
		panic(err)
	}

	if err = h.unitOfWork.Commit(); err != nil {
		panic(err)
	}

	outputPort.HandleResult(user)
}

func (h *SingUpHandler) isEmailExists(repo interfaces2.IUserReadRepository, email string) bool {
	usersByEmail := repo.Count(interfaces2.UsersFilter{
		Email: email,
	})

	return usersByEmail > 0
}
