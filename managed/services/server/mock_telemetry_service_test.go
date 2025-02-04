// Code generated by mockery. DO NOT EDIT.

package server

import (
	mock "github.com/stretchr/testify/mock"

	serverv1 "github.com/percona/pmm/api/server/v1"
)

// mockTelemetryService is an autogenerated mock type for the telemetryService type
type mockTelemetryService struct {
	mock.Mock
}

// DistributionMethod provides a mock function with no fields
func (_m *mockTelemetryService) DistributionMethod() serverv1.DistributionMethod {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DistributionMethod")
	}

	var r0 serverv1.DistributionMethod
	if rf, ok := ret.Get(0).(func() serverv1.DistributionMethod); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(serverv1.DistributionMethod)
	}

	return r0
}

// GetSummaries provides a mock function with no fields
func (_m *mockTelemetryService) GetSummaries() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSummaries")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// newMockTelemetryService creates a new instance of mockTelemetryService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockTelemetryService(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockTelemetryService {
	mock := &mockTelemetryService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
