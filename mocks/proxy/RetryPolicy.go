// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	message "github.com/datastax/go-cassandra-native-protocol/message"
	mock "github.com/stretchr/testify/mock"

	proxy "github.com/ollionorg/cassandra-to-spanner-proxy/third_party/datastax/proxy"
)

// RetryPolicy is an autogenerated mock type for the RetryPolicy type
type RetryPolicy struct {
	mock.Mock
}

// OnErrorResponse provides a mock function with given fields: msg, retryCount
func (_m *RetryPolicy) OnErrorResponse(msg message.Error, retryCount int) proxy.RetryDecision {
	ret := _m.Called(msg, retryCount)

	if len(ret) == 0 {
		panic("no return value specified for OnErrorResponse")
	}

	var r0 proxy.RetryDecision
	if rf, ok := ret.Get(0).(func(message.Error, int) proxy.RetryDecision); ok {
		r0 = rf(msg, retryCount)
	} else {
		r0 = ret.Get(0).(proxy.RetryDecision)
	}

	return r0
}

// OnReadTimeout provides a mock function with given fields: msg, retryCount
func (_m *RetryPolicy) OnReadTimeout(msg *message.ReadTimeout, retryCount int) proxy.RetryDecision {
	ret := _m.Called(msg, retryCount)

	if len(ret) == 0 {
		panic("no return value specified for OnReadTimeout")
	}

	var r0 proxy.RetryDecision
	if rf, ok := ret.Get(0).(func(*message.ReadTimeout, int) proxy.RetryDecision); ok {
		r0 = rf(msg, retryCount)
	} else {
		r0 = ret.Get(0).(proxy.RetryDecision)
	}

	return r0
}

// OnUnavailable provides a mock function with given fields: msg, retryCount
func (_m *RetryPolicy) OnUnavailable(msg *message.Unavailable, retryCount int) proxy.RetryDecision {
	ret := _m.Called(msg, retryCount)

	if len(ret) == 0 {
		panic("no return value specified for OnUnavailable")
	}

	var r0 proxy.RetryDecision
	if rf, ok := ret.Get(0).(func(*message.Unavailable, int) proxy.RetryDecision); ok {
		r0 = rf(msg, retryCount)
	} else {
		r0 = ret.Get(0).(proxy.RetryDecision)
	}

	return r0
}

// OnWriteTimeout provides a mock function with given fields: msg, retryCount
func (_m *RetryPolicy) OnWriteTimeout(msg *message.WriteTimeout, retryCount int) proxy.RetryDecision {
	ret := _m.Called(msg, retryCount)

	if len(ret) == 0 {
		panic("no return value specified for OnWriteTimeout")
	}

	var r0 proxy.RetryDecision
	if rf, ok := ret.Get(0).(func(*message.WriteTimeout, int) proxy.RetryDecision); ok {
		r0 = rf(msg, retryCount)
	} else {
		r0 = ret.Get(0).(proxy.RetryDecision)
	}

	return r0
}

// NewRetryPolicy creates a new instance of RetryPolicy. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRetryPolicy(t interface {
	mock.TestingT
	Cleanup(func())
}) *RetryPolicy {
	mock := &RetryPolicy{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
