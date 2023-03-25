package interfaces

import "github.com/Abdulrahman-Tayara/notes-app/shared/core"

type (
	IReadRepository[TEntity core.Entity, TFilters any] interface {
		GetById(id core.ID) (*TEntity, error)

		GetOne(filter *TEntity) (*TEntity, error)

		GetAll(filters TFilters) ([]TEntity, error)

		Count(filters TFilters) int32
	}
)

type (
	IWriteRepository[TEntity core.Entity] interface {
		Save(*TEntity) (*TEntity, error)

		Delete(*TEntity) error

		DeleteById(id core.ID) error

		Update(*TEntity) error
	}
)
