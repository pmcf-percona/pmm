// Code generated by mockery v2.40.2. DO NOT EDIT.

package agentlocal

import (
	mock "github.com/stretchr/testify/mock"

	agentlocalpb "github.com/percona/pmm/api/agentlocalpb"
)

// mockSupervisor is an autogenerated mock type for the supervisor type
type mockSupervisor struct {
	mock.Mock
}

// AgentsList provides a mock function with given fields:
func (_m *mockSupervisor) AgentsList() []*agentlocalpb.AgentInfo {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AgentsList")
	}

	var r0 []*agentlocalpb.AgentInfo
	if rf, ok := ret.Get(0).(func() []*agentlocalpb.AgentInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*agentlocalpb.AgentInfo)
		}
	}

	return r0
}

// AgentsLogs provides a mock function with given fields:
func (_m *mockSupervisor) AgentsLogs() map[string][]string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AgentsLogs")
	}

	var r0 map[string][]string
	if rf, ok := ret.Get(0).(func() map[string][]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]string)
		}
	}

	return r0
}

// newMockSupervisor creates a new instance of mockSupervisor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockSupervisor(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockSupervisor {
	mock := &mockSupervisor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
