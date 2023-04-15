// Code generated by mockery v2.14.0. DO NOT EDIT.

package interfaces

import (
	interfaces "github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/interfaces"
	mock "github.com/stretchr/testify/mock"
)

// IUnitOfWork is an autogenerated mock type for the IUnitOfWork type
type IUnitOfWork struct {
	mock.Mock
}

// Begin provides a mock function with given fields:
func (_m *IUnitOfWork) Begin() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Commit provides a mock function with given fields:
func (_m *IUnitOfWork) Commit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Rollback provides a mock function with given fields:
func (_m *IUnitOfWork) Rollback() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store provides a mock function with given fields:
func (_m *IUnitOfWork) Store() interfaces.IRepositoriesConstructor {
	ret := _m.Called()

	var r0 interfaces.IRepositoriesConstructor
	if rf, ok := ret.Get(0).(func() interfaces.IRepositoriesConstructor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.IRepositoriesConstructor)
		}
	}

	return r0
}

type mockConstructorTestingTNewIUnitOfWork interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUnitOfWork creates a new instance of IUnitOfWork. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUnitOfWork(t mockConstructorTestingTNewIUnitOfWork) *IUnitOfWork {
	mock := &IUnitOfWork{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}