// Code generated by mockery v2.33.0. DO NOT EDIT.

package inventory

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// mockInventoryMetrics is an autogenerated mock type for the inventoryMetrics type
type mockInventoryMetrics struct {
	mock.Mock
}

// GetAgentMetrics provides a mock function with given fields: ctx
func (_m *mockInventoryMetrics) GetAgentMetrics(ctx context.Context) ([]Metric, error) {
	ret := _m.Called(ctx)

	var r0 []Metric
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]Metric, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []Metric); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Metric)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodeMetrics provides a mock function with given fields: ctx
func (_m *mockInventoryMetrics) GetNodeMetrics(ctx context.Context) ([]Metric, error) {
	ret := _m.Called(ctx)

	var r0 []Metric
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]Metric, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []Metric); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Metric)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServiceMetrics provides a mock function with given fields: ctx
func (_m *mockInventoryMetrics) GetServiceMetrics(ctx context.Context) ([]Metric, error) {
	ret := _m.Called(ctx)

	var r0 []Metric
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]Metric, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []Metric); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Metric)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newMockInventoryMetrics creates a new instance of mockInventoryMetrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockInventoryMetrics(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockInventoryMetrics {
	mock := &mockInventoryMetrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
