// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	user "github.com/avinashb98/myfin/repository/user"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: _a0, _a1, _a2
func (_m *Repository) CreateUser(_a0 context.Context, _a1 user.User, _a2 user.Auth) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, user.User, user.Auth) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserAuthByHandle provides a mock function with given fields: _a0, _a1
func (_m *Repository) GetUserAuthByHandle(_a0 context.Context, _a1 string) (*user.Auth, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *user.Auth
	if rf, ok := ret.Get(0).(func(context.Context, string) *user.Auth); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Auth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByHandle provides a mock function with given fields: _a0, _a1
func (_m *Repository) GetUserByHandle(_a0 context.Context, _a1 string) (*user.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *user.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
