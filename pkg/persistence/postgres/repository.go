package postgres

import (
	"errors"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/core"
	sharederrors "github.com/Abdulrahman-Tayara/notes-app/pkg/errors"
	"gorm.io/gorm"
)

type ReadRepository[TEntity core.Entity, TFilters any] struct {
	db *gorm.DB

	filtersMapper func(TFilters) any
}

func NewPostgresReadRepository[TEntity core.Entity, TFilters any](db *gorm.DB, filtersMapper func(TFilters) any) *ReadRepository[TEntity, TFilters] {
	return &ReadRepository[TEntity, TFilters]{
		db:            db,
		filtersMapper: filtersMapper,
	}
}

func (r ReadRepository[TEntity, TFilters]) GetById(id core.ID) (*TEntity, error) {
	var model TEntity

	res := r.db.Where("id = ?", id.String()).First(&model)

	if res.Error != nil {
		return nil, normalizeGORMErrors(res.Error)
	}

	return &model, nil
}

func (r ReadRepository[TEntity, TFilters]) GetOne(filter *TEntity) (*TEntity, error) {
	var model TEntity
	res := r.db.Where(filter).First(&model)

	if res.Error != nil {
		return nil, normalizeGORMErrors(res.Error)
	}

	return &model, nil
}

func (r ReadRepository[TEntity, TFilters]) GetAll(filters TFilters) (entities []TEntity, err error) {
	filter := r.filtersMapper(filters)

	var model TEntity

	var models []TEntity

	res := r.db.Model(&model).Where(filter).Find(&models)

	err = normalizeGORMErrors(res.Error)

	if err != nil {
		return
	}

	entities = models

	return
}

func (r ReadRepository[TEntity, TFilters]) Count(filters TFilters) int32 {
	var model TEntity

	filter := r.filtersMapper(filters)

	count := int64(0)

	_ = r.db.Model(&model).Where(filter).Count(&count)

	return int32(count)
}

type WriteRepository[TEntity core.Entity] struct {
	db *gorm.DB
}

func NewPostgresWriteRepository[TEntity core.Entity](db *gorm.DB) *WriteRepository[TEntity] {
	return &WriteRepository[TEntity]{
		db: db,
	}
}

func (w WriteRepository[TEntity]) Save(entity *TEntity) (*TEntity, error) {
	res := w.db.Save(&entity)

	if res.Error != nil {
		return nil, normalizeGORMErrors(res.Error)
	}

	return entity, nil
}

func (w WriteRepository[TEntity]) Delete(entity *TEntity) (err error) {
	res := w.db.Delete(&entity)

	err = normalizeGORMErrors(res.Error)

	return
}

func (w WriteRepository[TEntity]) DeleteById(id core.ID) (err error) {
	var model TEntity

	res := w.db.Where("id = ?", id.String()).Delete(&model)

	err = normalizeGORMErrors(res.Error)

	return
}

func (w WriteRepository[TEntity]) Update(entity *TEntity) (err error) {
	res := w.db.Save(entity)

	err = normalizeGORMErrors(res.Error)

	return
}

func normalizeGORMErrors(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return sharederrors.ErrEntityNotFound
	}

	return err
}
