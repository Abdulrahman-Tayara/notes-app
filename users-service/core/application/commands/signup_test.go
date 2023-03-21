package commands

import (
	"context"
	goerrors "errors"
	"github.com/Abdulrahman-Tayara/notes-app/shared/errors"
	interfaces2 "github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/ports"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain/entity"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/internal/mocks/application/interfaces"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/internal/mocks/application/services"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func TestSingUpHandler_Handle(t *testing.T) {
	tests := map[string]struct {
		inputs      SignUp
		outputError error
		outputUser  *entity.User
		mockFunc    func(*testing.T, *interfaces.IUnitOfWork, *interfaces.IRepositoriesConstructor, *interfaces.IUserReadRepository, *interfaces.IUserWriteRepository, *services.IHashService)
	}{
		"Success": {
			inputs:      SignUp{Name: "Abdulrahman", Email: "abdulrahman@gmail.com", Password: "12345"},
			outputError: nil,
			outputUser:  &entity.User{Name: "Abdulrahman", Email: "abdulrahman@gmail.com", Password: "hashedPassword"},
			mockFunc: func(test *testing.T, unitOfWorkMock *interfaces.IUnitOfWork, storeMock *interfaces.IRepositoriesConstructor, userReadRepoMock *interfaces.IUserReadRepository, usersWriteRepoMock *interfaces.IUserWriteRepository, hashService *services.IHashService) {
				hashService.On("HashString", "12345").Return("hashedPassword")
				storeMock.On("UsersRead").Return(userReadRepoMock)
				storeMock.On("UsersWrite").Return(usersWriteRepoMock)

				userToSave := &entity.User{Name: "Abdulrahman", Email: "abdulrahman@gmail.com", Password: "hashedPassword"}

				userReadRepoMock.On("Count", interfaces2.UsersFilter{Email: "abdulrahman@gmail.com"}).Return(int32(0))

				unitOfWorkMock.On("Begin").Return(nil)
				unitOfWorkMock.On("Store").Return(storeMock)
				usersWriteRepoMock.On("Save", mock.Anything).Return(userToSave, nil)
				unitOfWorkMock.On("Commit").Return(nil)

				unitOfWorkMock.AssertNotCalled(test, "Rollback")
			},
		},
		"Commit error": {
			inputs:      SignUp{Name: "Abdulrahman", Email: "abdulrahman@gmail.com", Password: "12345"},
			outputError: goerrors.New("some commit error"),
			outputUser:  nil,
			mockFunc: func(test *testing.T, work *interfaces.IUnitOfWork, constructor *interfaces.IRepositoriesConstructor, readRepository *interfaces.IUserReadRepository, writeRepository *interfaces.IUserWriteRepository, service *services.IHashService) {
				service.On("HashString", "12345").Return("hashedPassword")
				constructor.On("UsersRead").Return(readRepository)
				constructor.On("UsersWrite").Return(writeRepository)

				readRepository.On("Count", interfaces2.UsersFilter{Email: "abdulrahman@gmail.com"}).Return(int32(0))

				work.On("Begin").Return(nil)
				work.On("Store").Return(constructor)
				writeRepository.On("Save", mock.Anything).Return(&entity.User{}, nil)
				work.On("Commit").Return(goerrors.New("some commit error"))
				work.On("Rollback").Return(nil)
			},
		},
		"InValid Email": {
			inputs:      SignUp{Name: "Abdulrahman", Email: "abdulrahman", Password: "12345"},
			outputError: errors.BadValueException("email"),
			mockFunc: func(test *testing.T, unitOfWorkMock *interfaces.IUnitOfWork, storeMock *interfaces.IRepositoriesConstructor, userReadRepoMock *interfaces.IUserReadRepository, usersWriteRepoMock *interfaces.IUserWriteRepository, hashService *services.IHashService) {

			},
		},

		"Email Already Exists": {
			inputs:      SignUp{Name: "Abdulrahman", Email: "abdulrahman@gmail.com", Password: "12345"},
			outputError: domain.EmailAlreadyExists,
			mockFunc: func(test *testing.T, unitOfWorkMock *interfaces.IUnitOfWork, storeMock *interfaces.IRepositoriesConstructor, userReadRepoMock *interfaces.IUserReadRepository, usersWriteRepoMock *interfaces.IUserWriteRepository, hashService *services.IHashService) {
				hashService.On("HashString", "12345").Return("hashedPassword")
				storeMock.On("UsersRead").Return(userReadRepoMock)

				unitOfWorkMock.On("Store").Return(storeMock)

				userReadRepoMock.On("Count", interfaces2.UsersFilter{Email: "abdulrahman@gmail.com"}).Return(int32(1))
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			usersReadRepoMock := interfaces.NewIUserReadRepository(t)
			usersWriteRepoMock := interfaces.NewIUserWriteRepository(t)
			storeMock := interfaces.NewIRepositoriesConstructor(t)
			unitOfWorkMock := interfaces.NewIUnitOfWork(t)

			hashService := services.NewIHashService(t)

			output := &ports.MockOutputPort[*entity.User]{}

			handler := NewSingUpHandler(unitOfWorkMock, hashService)

			test.mockFunc(t, unitOfWorkMock, storeMock, usersReadRepoMock, usersWriteRepoMock, hashService)

			handler.Handle(context.TODO(), SignUp{Name: test.inputs.Name, Email: test.inputs.Email, Password: test.inputs.Password}, output)

			if !reflect.DeepEqual(output.Err, test.outputError) {
				t.Errorf("expceted err: %v, actual err: %v", test.outputError, output.Err)
			}

		})
	}
}
