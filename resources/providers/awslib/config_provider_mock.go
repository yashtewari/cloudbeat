// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
// Code generated by mockery v2.10.4. DO NOT EDIT.

package awslib

import mock "github.com/stretchr/testify/mock"

// MockConfigGetter is an autogenerated mock type for the ConfigGetter type
type MockConfigGetter struct {
	mock.Mock
}

type MockConfigGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConfigGetter) EXPECT() *MockConfigGetter_Expecter {
	return &MockConfigGetter_Expecter{mock: &_m.Mock}
}

// GetConfig provides a mock function with given fields:
func (_m *MockConfigGetter) GetConfig() Config {
	ret := _m.Called()

	var r0 Config
	if rf, ok := ret.Get(0).(func() Config); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(Config)
	}

	return r0
}

// MockConfigGetter_GetConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConfig'
type MockConfigGetter_GetConfig_Call struct {
	*mock.Call
}

// GetConfig is a helper method to define mock.On call
func (_e *MockConfigGetter_Expecter) GetConfig() *MockConfigGetter_GetConfig_Call {
	return &MockConfigGetter_GetConfig_Call{Call: _e.mock.On("GetConfig")}
}

func (_c *MockConfigGetter_GetConfig_Call) Run(run func()) *MockConfigGetter_GetConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockConfigGetter_GetConfig_Call) Return(_a0 Config) *MockConfigGetter_GetConfig_Call {
	_c.Call.Return(_a0)
	return _c
}