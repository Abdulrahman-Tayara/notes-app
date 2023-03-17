package ports

type IOutputPort[TResult any] interface {
	HandleError(error)

	HandleResult(TResult)
}

type MockOutputPort[TResult any] struct {
	Result TResult
	Err    error
}

func (m *MockOutputPort[TResult]) HandleError(err error) {
	m.Err = err
}

func (m *MockOutputPort[TResult]) HandleResult(result TResult) {
	m.Result = result
}
