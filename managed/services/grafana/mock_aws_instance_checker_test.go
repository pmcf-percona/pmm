// Code generated by mockery v2.36.0. DO NOT EDIT.

package grafana

import mock "github.com/stretchr/testify/mock"

// mockAwsInstanceChecker is an autogenerated mock type for the awsInstanceChecker type
type mockAwsInstanceChecker struct {
	mock.Mock
}

// MustCheck provides a mock function with given fields:
func (_m *mockAwsInstanceChecker) MustCheck() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// newMockAwsInstanceChecker creates a new instance of mockAwsInstanceChecker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockAwsInstanceChecker(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockAwsInstanceChecker {
	mock := &mockAwsInstanceChecker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
