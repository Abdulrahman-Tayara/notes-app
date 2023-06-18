package postgres

import (
	"errors"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/persistence"
	"gorm.io/gorm"
)

type StoreFactory[TStore any] interface {
	Create(db *gorm.DB) TStore
}

type UnitOfWork[TStore any] struct {
	db           *gorm.DB
	storeFactory StoreFactory[TStore]

	openedTransaction *gorm.DB
}

func NewPostgresUnitOfWork[TStore any](db *gorm.DB, factory StoreFactory[TStore]) persistence.IUnitOfWork[TStore] {
	return &UnitOfWork[TStore]{
		db:           db,
		storeFactory: factory,
	}
}

func (p UnitOfWork[TStore]) Begin() error {
	if p.openedTransaction != nil {
		return errors.New("there is an opened transaction")
	}

	p.openedTransaction = p.db.Begin()

	return nil
}

func (p UnitOfWork[TStore]) Commit() error {
	if p.openedTransaction != nil {
		p.openedTransaction.Commit()
		p.openedTransaction = nil
	}

	return nil
}

func (p UnitOfWork[TStore]) Rollback() error {
	if p.openedTransaction != nil {
		p.openedTransaction.Rollback()
		p.openedTransaction = nil
	}

	return nil
}

func (p UnitOfWork[TStore]) Store() TStore {
	if p.openedTransaction != nil {
		return p.storeFactory.Create(p.openedTransaction)
	}

	return p.storeFactory.Create(p.db)
}
