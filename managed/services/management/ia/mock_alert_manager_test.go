// Code generated by mockery v2.40.2. DO NOT EDIT.

package ia

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	ammodels "github.com/percona/pmm/api/alertmanager/ammodels"
	services "github.com/percona/pmm/managed/services"
)

// mockAlertManager is an autogenerated mock type for the alertManager type
type mockAlertManager struct {
	mock.Mock
}

// FindAlertsByID provides a mock function with given fields: ctx, params, ids
func (_m *mockAlertManager) FindAlertsByID(ctx context.Context, params *services.FilterParams, ids []string) ([]*ammodels.GettableAlert, error) {
	ret := _m.Called(ctx, params, ids)

	if len(ret) == 0 {
		panic("no return value specified for FindAlertsByID")
	}

	var r0 []*ammodels.GettableAlert
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *services.FilterParams, []string) ([]*ammodels.GettableAlert, error)); ok {
		return rf(ctx, params, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *services.FilterParams, []string) []*ammodels.GettableAlert); ok {
		r0 = rf(ctx, params, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ammodels.GettableAlert)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *services.FilterParams, []string) error); ok {
		r1 = rf(ctx, params, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAlerts provides a mock function with given fields: ctx, params
func (_m *mockAlertManager) GetAlerts(ctx context.Context, params *services.FilterParams) ([]*ammodels.GettableAlert, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for GetAlerts")
	}

	var r0 []*ammodels.GettableAlert
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *services.FilterParams) ([]*ammodels.GettableAlert, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *services.FilterParams) []*ammodels.GettableAlert); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ammodels.GettableAlert)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *services.FilterParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestConfigurationUpdate provides a mock function with given fields:
func (_m *mockAlertManager) RequestConfigurationUpdate() {
	_m.Called()
}

// SilenceAlerts provides a mock function with given fields: ctx, alerts
func (_m *mockAlertManager) SilenceAlerts(ctx context.Context, alerts []*ammodels.GettableAlert) error {
	ret := _m.Called(ctx, alerts)

	if len(ret) == 0 {
		panic("no return value specified for SilenceAlerts")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*ammodels.GettableAlert) error); ok {
		r0 = rf(ctx, alerts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnsilenceAlerts provides a mock function with given fields: ctx, alerts
func (_m *mockAlertManager) UnsilenceAlerts(ctx context.Context, alerts []*ammodels.GettableAlert) error {
	ret := _m.Called(ctx, alerts)

	if len(ret) == 0 {
		panic("no return value specified for UnsilenceAlerts")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*ammodels.GettableAlert) error); ok {
		r0 = rf(ctx, alerts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockAlertManager creates a new instance of mockAlertManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockAlertManager(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockAlertManager {
	mock := &mockAlertManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
