// Code generated by mockery v2.14.0. DO NOT EDIT.

package interfaces

import (
	core "github.com/Abdulrahman-Tayara/notes-app/pkg/core"
	interfaces "github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/interfaces"
	entity "github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// IUserReadRepository is an autogenerated mock type for the IUserReadRepository type
type IUserReadRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: filters
func (_m *IUserReadRepository) Count(filters interfaces.UsersFilter) int32 {
	ret := _m.Called(filters)

	var r0 int32
	if rf, ok := ret.Get(0).(func(interfaces.UsersFilter) int32); ok {
		r0 = rf(filters)
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// GetAll provides a mock function with given fields: filters
func (_m *IUserReadRepository) GetAll(filters interfaces.UsersFilter) ([]entity.User, error) {
	ret := _m.Called(filters)

	var r0 []entity.User
	if rf, ok := ret.Get(0).(func(interfaces.UsersFilter) []entity.User); ok {
		r0 = rf(filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interfaces.UsersFilter) error); ok {
		r1 = rf(filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *IUserReadRepository) GetById(id core.ID) (*entity.User, error) {
	ret := _m.Called(id)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(core.ID) *entity.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.ID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOne provides a mock function with given fields: filter
func (_m *IUserReadRepository) GetOne(filter *entity.User) (*entity.User, error) {
	ret := _m.Called(filter)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(*entity.User) *entity.User); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.User) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIUserReadRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUserReadRepository creates a new instance of IUserReadRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUserReadRepository(t mockConstructorTestingTNewIUserReadRepository) *IUserReadRepository {
	mock := &IUserReadRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
