package commands

import (
	"context"
	"github.com/Abdulrahman-Tayara/notes-app/shared/errors"
	interfaces2 "github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/ports"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain/entity"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/internal/mocks/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/internal/mocks/application/services"
	"reflect"
	"testing"
)

func TestSingUpHandler_Handle(t *testing.T) {
	tests := map[string]struct {
		inputs      SignUp
		outputError error
		outputUser  *entity.User
		mockFunc    func(*interfaces.IUnitOfWork, *interfaces.IRepositoriesConstructor, *interfaces.IUserReadRepository, *interfaces.IUserWriteRepository, *services.IHashService)
	}{
		"Success": {
			inputs:      SignUp{Email: "abdulrahman@gmail.com", Password: "12345"},
			outputError: nil,
			outputUser:  &entity.User{Email: "abdulrahman@gmail.com", Password: "12345"},
			mockFunc: func(unitOfWorkMock *interfaces.IUnitOfWork, storeMock *interfaces.IRepositoriesConstructor, userReadRepoMock *interfaces.IUserReadRepository, usersWriteRepoMock *interfaces.IUserWriteRepository, hashService *services.IHashService) {
				hashService.On("HashString", "12345").Return("hashedPassword")
				storeMock.On("UsersRead").Return(userReadRepoMock)
				storeMock.On("UsersWrite").Return(usersWriteRepoMock)

				unitOfWorkMock.On("Begin").Return(nil)
				unitOfWorkMock.On("Store").Return(storeMock)
				unitOfWorkMock.On("Commit").Return(nil)

				userReadRepoMock.On("Count", interfaces2.UsersFilter{Email: "abdulrahman@gmail.com"}).Return(int32(0))

				userToSave := &entity.User{Email: "abdulrahman@gmail.com", Password: "hashedPassword"}

				usersWriteRepoMock.On("Save", userToSave).Return(userToSave, nil)
			},
		},

		"InValid Email": {
			inputs:      SignUp{Email: "abdulrahman", Password: "12345"},
			outputError: errors.BadValueException("email"),
			mockFunc: func(unitOfWorkMock *interfaces.IUnitOfWork, storeMock *interfaces.IRepositoriesConstructor, userReadRepoMock *interfaces.IUserReadRepository, usersWriteRepoMock *interfaces.IUserWriteRepository, hashService *services.IHashService) {

			},
		},

		"Email Already Exists": {
			inputs:      SignUp{Email: "abdulrahman@gmail.com", Password: "12345"},
			outputError: domain.EmailAlreadyExists,
			mockFunc: func(unitOfWorkMock *interfaces.IUnitOfWork, storeMock *interfaces.IRepositoriesConstructor, userReadRepoMock *interfaces.IUserReadRepository, usersWriteRepoMock *interfaces.IUserWriteRepository, hashService *services.IHashService) {
				hashService.On("HashString", "12345").Return("hashedPassword")
				storeMock.On("UsersRead").Return(userReadRepoMock)

				unitOfWorkMock.On("Store").Return(storeMock)

				userReadRepoMock.On("Count", interfaces2.UsersFilter{Email: "abdulrahman@gmail.com"}).Return(int32(1))
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			usersReadRepoMock := interfaces.NewIUserReadRepository(t)
			usersWriteRepoMock := interfaces.NewIUserWriteRepository(t)
			storeMock := interfaces.NewIRepositoriesConstructor(t)
			unitOfWorkMock := interfaces.NewIUnitOfWork(t)

			hashService := services.NewIHashService(t)

			output := &ports.MockOutputPort[*entity.User]{}

			handler := NewSingUpHandler(unitOfWorkMock, hashService)

			test.mockFunc(unitOfWorkMock, storeMock, usersReadRepoMock, usersWriteRepoMock, hashService)

			handler.Handle(context.TODO(), SignUp{Email: test.inputs.Email, Password: test.inputs.Password}, output)

			if !reflect.DeepEqual(output.Err, test.outputError) {
				t.Errorf("expceted err: %v, actual err: %v", test.outputError, output.Err)
			}

		})
	}
}
