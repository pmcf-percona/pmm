// Code generated by mockery v2.33.0. DO NOT EDIT.

package checks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	ammodels "github.com/percona/pmm/api/alertmanager/ammodels"
	services "github.com/percona/pmm/managed/services"
)

// mockAlertmanagerService is an autogenerated mock type for the alertmanagerService type
type mockAlertmanagerService struct {
	mock.Mock
}

// GetAlerts provides a mock function with given fields: ctx, params
func (_m *mockAlertmanagerService) GetAlerts(ctx context.Context, params *services.FilterParams) ([]*ammodels.GettableAlert, error) {
	ret := _m.Called(ctx, params)

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

// SendAlerts provides a mock function with given fields: ctx, alerts
func (_m *mockAlertmanagerService) SendAlerts(ctx context.Context, alerts ammodels.PostableAlerts) {
	_m.Called(ctx, alerts)
}

// SilenceAlerts provides a mock function with given fields: ctx, alerts
func (_m *mockAlertmanagerService) SilenceAlerts(ctx context.Context, alerts []*ammodels.GettableAlert) error {
	ret := _m.Called(ctx, alerts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*ammodels.GettableAlert) error); ok {
		r0 = rf(ctx, alerts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnsilenceAlerts provides a mock function with given fields: ctx, alerts
func (_m *mockAlertmanagerService) UnsilenceAlerts(ctx context.Context, alerts []*ammodels.GettableAlert) error {
	ret := _m.Called(ctx, alerts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*ammodels.GettableAlert) error); ok {
		r0 = rf(ctx, alerts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockAlertmanagerService creates a new instance of mockAlertmanagerService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockAlertmanagerService(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockAlertmanagerService {
	mock := &mockAlertmanagerService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
