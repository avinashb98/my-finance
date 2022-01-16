// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mongo "github.com/avinashb98/myfin/datasources/mongo"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Connect provides a mock function with given fields:
func (_m *Client) Connect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Database provides a mock function with given fields: _a0
func (_m *Client) Database(_a0 string) mongo.Database {
	ret := _m.Called(_a0)

	var r0 mongo.Database
	if rf, ok := ret.Get(0).(func(string) mongo.Database); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongo.Database)
		}
	}

	return r0
}
