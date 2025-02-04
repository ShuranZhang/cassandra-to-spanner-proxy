/*
 * Copyright (C) 2024 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy of
 * the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */
 
// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	message "github.com/datastax/go-cassandra-native-protocol/message"
	mock "github.com/stretchr/testify/mock"

	responsehandler "github.com/cloudspannerecosystem/cassandra-to-spanner-proxy/responsehandler"

	spanner "github.com/cloudspannerecosystem/cassandra-to-spanner-proxy/spanner"

	tableConfig "github.com/cloudspannerecosystem/cassandra-to-spanner-proxy/tableConfig"

)

// SpannerClientIface is an autogenerated mock type for the SpannerClientIface type
type SpannerClientIface struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *SpannerClientIface) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMapByKey provides a mock function with given fields: ctx, query
func (_m *SpannerClientIface) UpdateMapByKey(ctx context.Context, query responsehandler.QueryMetadata) (*message.RowsResult, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMapByKey")
	}

	var r0 *message.RowsResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, responsehandler.QueryMetadata) (*message.RowsResult, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, responsehandler.QueryMetadata) *message.RowsResult); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*message.RowsResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, responsehandler.QueryMetadata) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUsingMutations provides a mock function with given fields: ctx, query
func (_m *SpannerClientIface) DeleteUsingMutations(ctx context.Context, query responsehandler.QueryMetadata) (*message.RowsResult, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUsingMutations")
	}

	var r0 *message.RowsResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, responsehandler.QueryMetadata) (*message.RowsResult, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, responsehandler.QueryMetadata) *message.RowsResult); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*message.RowsResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, responsehandler.QueryMetadata) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterAndExecuteBatch provides a mock function with given fields: ctx, queries
func (_m *SpannerClientIface) FilterAndExecuteBatch(ctx context.Context, queries []*responsehandler.QueryMetadata) (*message.RowsResult, string, error) {
	ret := _m.Called(ctx, queries)

	if len(ret) == 0 {
			panic("no return value specified for FilterAndExecuteBatch")
	}

	var r0 *message.RowsResult
	var r1 string  
	var r2 error  

	if rf, ok := ret.Get(0).(func(context.Context, []*responsehandler.QueryMetadata) (*message.RowsResult, string, error)); ok {
			r0, r1, r2 = rf(ctx, queries) 
	} else {
			if rf, ok := ret.Get(0).(func(context.Context, []*responsehandler.QueryMetadata) *message.RowsResult); ok {
					r0 = rf(ctx, queries)
			} else {
					if ret.Get(0) != nil {
							r0 = ret.Get(0).(*message.RowsResult)
					}
			}

			if rf, ok := ret.Get(1).(func(context.Context, []*responsehandler.QueryMetadata) string); ok { // Add this block for string
					r1 = rf(ctx, queries)
			} else {
					if ret.Get(1) != nil {
							r1 = ret.Get(1).(string)
					}
			}
			if rf, ok := ret.Get(2).(func(context.Context, []*responsehandler.QueryMetadata) error); ok { // Adjust index to 2.
					r2 = rf(ctx, queries)
			} else {
					r2 = ret.Error(2) 
			}

	}

	return r0, r1, r2 
}

// GetTableConfigs provides a mock function with given fields: ctx, query
func (_m *SpannerClientIface) GetTableConfigs(ctx context.Context, query responsehandler.QueryMetadata) (map[string]map[string]*tableConfig.Column, map[string][]tableConfig.Column, *spanner.SystemQueryMetadataCache, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for GetTableConfigs")
	}

	var r0 map[string]map[string]*tableConfig.Column
	var r1 map[string][]tableConfig.Column
	var r2 *spanner.SystemQueryMetadataCache
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, responsehandler.QueryMetadata) (map[string]map[string]*tableConfig.Column, map[string][]tableConfig.Column, *spanner.SystemQueryMetadataCache, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, responsehandler.QueryMetadata) map[string]map[string]*tableConfig.Column); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]map[string]*tableConfig.Column)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, responsehandler.QueryMetadata) map[string][]tableConfig.Column); ok {
		r1 = rf(ctx, query)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string][]tableConfig.Column)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, responsehandler.QueryMetadata) *spanner.SystemQueryMetadataCache); ok {
		r2 = rf(ctx, query)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*spanner.SystemQueryMetadataCache)
		}
	}

	if rf, ok := ret.Get(3).(func(context.Context, responsehandler.QueryMetadata) error); ok {
		r3 = rf(ctx, query)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// InsertOrUpdateMutation provides a mock function with given fields: ctx, query
func (_m *SpannerClientIface) InsertOrUpdateMutation(ctx context.Context, query responsehandler.QueryMetadata) (*message.RowsResult, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for InsertOrUpdateMutation")
	}

	var r0 *message.RowsResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, responsehandler.QueryMetadata) (*message.RowsResult, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, responsehandler.QueryMetadata) *message.RowsResult); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*message.RowsResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, responsehandler.QueryMetadata) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertUpdateOrDeleteStatement provides a mock function with given fields: ctx, query
func (_m *SpannerClientIface) InsertUpdateOrDeleteStatement(ctx context.Context, query responsehandler.QueryMetadata) (*message.RowsResult, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for InsertUpdateOrDeleteStatement")
	}

	var r0 *message.RowsResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, responsehandler.QueryMetadata) (*message.RowsResult, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, responsehandler.QueryMetadata) *message.RowsResult); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*message.RowsResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, responsehandler.QueryMetadata) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectStatement provides a mock function with given fields: ctx, query
func (_m *SpannerClientIface) SelectStatement(ctx context.Context, query responsehandler.QueryMetadata) (*message.RowsResult, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for SelectStatement")
	}

	var r0 *message.RowsResult
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, responsehandler.QueryMetadata) (*message.RowsResult, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, responsehandler.QueryMetadata) *message.RowsResult); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*message.RowsResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, responsehandler.QueryMetadata) error); ok {
		r2 = rf(ctx, query)
	} else {
		r2 = ret.Error(1)
	}

	return r0, r2
}

// NewSpannerClientIface creates a new instance of SpannerClientIface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSpannerClientIface(t interface {
	mock.TestingT
	Cleanup(func())
}) *SpannerClientIface {
	mock := &SpannerClientIface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
