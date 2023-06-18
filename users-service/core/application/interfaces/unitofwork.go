package interfaces

import (
	"github.com/Abdulrahman-Tayara/notes-app/pkg/persistence"
)

type IRepositoriesConstructor interface {
	UsersRead() IUserReadRepository
	UsersWrite() IUserWriteRepository
}

type IUnitOfWork interface {
	persistence.IUnitOfWork[IRepositoriesConstructor]
}
