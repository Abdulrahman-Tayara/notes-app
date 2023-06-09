package persistence

type IUnitOfWork[TStore any] interface {
	Begin() error
	Commit() error
	Rollback() error

	Store() TStore
}
