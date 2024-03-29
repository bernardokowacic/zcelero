// Code generated by mockery v2.14.1. DO NOT EDIT.

package helper

import (
	os "os"

	mock "github.com/stretchr/testify/mock"
)

// HelperInterface is an autogenerated mock type for the HelperInterface type
type HelperInterface struct {
	mock.Mock
}

// CreateFile provides a mock function with given fields: filePath
func (_m *HelperInterface) CreateFile(filePath string) (*os.File, error) {
	ret := _m.Called(filePath)

	var r0 *os.File
	if rf, ok := ret.Get(0).(func(string) *os.File); ok {
		r0 = rf(filePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*os.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateUuid provides a mock function with given fields:
func (_m *HelperInterface) GenerateUuid() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ReadFile provides a mock function with given fields: filePath
func (_m *HelperInterface) ReadFile(filePath string) ([]byte, error) {
	ret := _m.Called(filePath)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(filePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WriteFile provides a mock function with given fields: file, content
func (_m *HelperInterface) WriteFile(file *os.File, content string) (int, error) {
	ret := _m.Called(file, content)

	var r0 int
	if rf, ok := ret.Get(0).(func(*os.File, string) int); ok {
		r0 = rf(file, content)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*os.File, string) error); ok {
		r1 = rf(file, content)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewHelperInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewHelperInterface creates a new instance of HelperInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHelperInterface(t mockConstructorTestingTNewHelperInterface) *HelperInterface {
	mock := &HelperInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
