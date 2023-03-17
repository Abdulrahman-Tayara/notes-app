package interfaces

import (
	"github.com/Abdulrahman-Tayara/notes-app/shared/interfaces"
)

type IRepositoriesConstructor interface {
	UsersRead() IUserReadRepository
	UsersWrite() IUserWriteRepository
}

type IUnitOfWork interface {
	interfaces.IUnitOfWork[IRepositoriesConstructor]
}
