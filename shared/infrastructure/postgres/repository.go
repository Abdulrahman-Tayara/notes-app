package postgres

import (
	"errors"
	"github.com/Abdulrahman-Tayara/notes-app/shared/core"
	sharederrors "github.com/Abdulrahman-Tayara/notes-app/shared/errors"
	"github.com/Abdulrahman-Tayara/notes-app/shared/infrastructure"
	"gorm.io/gorm"
)

type ReadRepository[TEntity core.Entity, TModel infrastructure.BaseModel[TEntity], TFilters any] struct {
	db *gorm.DB

	filtersMapper func(TFilters) any
}

func NewPostgresReadRepository[TEntity core.Entity, TModel infrastructure.BaseModel[TEntity], TFilters any](db *gorm.DB, filtersMapper func(TFilters) any) *ReadRepository[TEntity, TModel, TFilters] {
	return &ReadRepository[TEntity, TModel, TFilters]{
		db:            db,
		filtersMapper: filtersMapper,
	}
}

func (r ReadRepository[TEntity, TModel, TFilters]) GetById(id core.ID) (*TEntity, error) {
	var model TModel

	res := r.db.Where("id = ?", id.String()).First(&model)

	if res.Error != nil {
		return nil, normalizeGORMErrors(res.Error)
	}

	return model.To(), nil
}

func (r ReadRepository[TEntity, TModel, TFilters]) GetAll(filters TFilters) (entities []TEntity, err error) {
	filter := r.filtersMapper(filters)

	var model TModel

	var models []TModel

	res := r.db.Model(&model).Where(filter).Find(&models)

	err = normalizeGORMErrors(res.Error)

	if err != nil {
		return
	}

	entities = infrastructure.MapToEntities[TEntity, TModel](models)

	return
}

func (r ReadRepository[TEntity, TModel, TFilters]) Count(filters TFilters) int32 {
	var model TModel

	filter := r.filtersMapper(filters)

	count := int64(0)

	_ = r.db.Model(&model).Where(filter).Count(&count)

	return int32(count)
}

type WriteRepository[TEntity core.Entity, TModel infrastructure.BaseModel[TEntity]] struct {
	modelFactory func(*TEntity) *TModel
	db           *gorm.DB
}

func NewPostgresWriteRepository[TEntity core.Entity, TModel infrastructure.BaseModel[TEntity]](
	db *gorm.DB, modelFactory func(*TEntity) *TModel) *WriteRepository[TEntity, TModel] {
	return &WriteRepository[TEntity, TModel]{
		modelFactory: modelFactory,
		db:           db,
	}
}

func (w WriteRepository[TEntity, TModel]) Save(entity *TEntity) (*TEntity, error) {
	var model = w.modelFactory(entity)

	res := w.db.Save(&model)

	if res.Error != nil {
		return nil, normalizeGORMErrors(res.Error)
	}

	return entity, nil
}

func (w WriteRepository[TEntity, TModel]) Delete(entity *TEntity) (err error) {
	var model = w.modelFactory(entity)

	res := w.db.Delete(&model)

	err = normalizeGORMErrors(res.Error)

	return
}

func (w WriteRepository[TEntity, TModel]) DeleteById(id core.ID) (err error) {
	var model TModel

	res := w.db.Where("id = ?", id.String()).Delete(&model)

	err = normalizeGORMErrors(res.Error)

	return
}

func (w WriteRepository[TEntity, TModel]) Update(entity *TEntity) (err error) {
	var model = w.modelFactory(entity)

	res := w.db.Save(model)

	err = normalizeGORMErrors(res.Error)

	return
}

func normalizeGORMErrors(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return sharederrors.ErrEntityNotFound
	}

	return err
}
