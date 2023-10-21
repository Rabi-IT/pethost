// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Database is an autogenerated mock type for the Database type
type Database struct {
	mock.Mock
}

// Start provides a mock function with given fields:
func (_m *Database) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDatabase creates a new instance of Database. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDatabase(t interface {
	mock.TestingT
	Cleanup(func())
}) *Database {
	mock := &Database{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
