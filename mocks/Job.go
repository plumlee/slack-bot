// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import gojenkins "github.com/bndr/gojenkins"

import mock "github.com/stretchr/testify/mock"

// Job is an autogenerated mock type for the Job type
type Job struct {
	mock.Mock
}

// GetBuild provides a mock function with given fields: id
func (_m *Job) GetBuild(id int64) (*gojenkins.Build, error) {
	ret := _m.Called(id)

	var r0 *gojenkins.Build
	if rf, ok := ret.Get(0).(func(int64) *gojenkins.Build); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gojenkins.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastBuild provides a mock function with given fields:
func (_m *Job) GetLastBuild() (*gojenkins.Build, error) {
	ret := _m.Called()

	var r0 *gojenkins.Build
	if rf, ok := ret.Get(0).(func() *gojenkins.Build); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gojenkins.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Poll provides a mock function with given fields:
func (_m *Job) Poll() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
