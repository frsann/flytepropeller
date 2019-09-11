// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

import mock "github.com/stretchr/testify/mock"

// TaskReader is an autogenerated mock type for the TaskReader type
type TaskReader struct {
	mock.Mock
}

// GetTaskID provides a mock function with given fields:
func (_m *TaskReader) GetTaskID() *core.Identifier {
	ret := _m.Called()

	var r0 *core.Identifier
	if rf, ok := ret.Get(0).(func() *core.Identifier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Identifier)
		}
	}

	return r0
}

// GetTaskType provides a mock function with given fields:
func (_m *TaskReader) GetTaskType() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Read provides a mock function with given fields: ctx
func (_m *TaskReader) Read(ctx context.Context) (*core.TaskTemplate, error) {
	ret := _m.Called(ctx)

	var r0 *core.TaskTemplate
	if rf, ok := ret.Get(0).(func(context.Context) *core.TaskTemplate); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.TaskTemplate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
