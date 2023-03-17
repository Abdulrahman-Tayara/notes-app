package infrastructure

import (
	"github.com/Abdulrahman-Tayara/notes-app/shared/core"
)

type BaseModel[TEntity core.Entity] interface {
	To() *TEntity
}

func MapToEntities[TEntity core.Entity, V BaseModel[TEntity]](records []V) []TEntity {
	var res []TEntity

	for _, record := range records {
		res = append(res, *record.To())
	}

	return res
}
