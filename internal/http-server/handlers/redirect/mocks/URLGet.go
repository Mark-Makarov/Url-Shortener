// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// URLGet is an autogenerated mock type for the URLGet type
type URLGet struct {
	mock.Mock
}

// GetURL provides a mock function with given fields: alias
func (_m *URLGet) GetURL(alias string) (string, error) {
	ret := _m.Called(alias)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(alias)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(alias)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(alias)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewURLGet interface {
	mock.TestingT
	Cleanup(func())
}

// NewURLGet creates a new instance of URLGet. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewURLGet(t mockConstructorTestingTNewURLGet) *URLGet {
	mock := &URLGet{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
