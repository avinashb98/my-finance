// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	net_worth "github.com/avinashb98/myfin/repository/net_worth"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateNetWorth provides a mock function with given fields: _a0, _a1
func (_m *Repository) CreateNetWorth(_a0 context.Context, _a1 net_worth.NetWorth) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, net_worth.NetWorth) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetNetWorth provides a mock function with given fields: _a0, _a1
func (_m *Repository) SetNetWorth(_a0 context.Context, _a1 net_worth.NetWorth) (*net_worth.NetWorth, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *net_worth.NetWorth
	if rf, ok := ret.Get(0).(func(context.Context, net_worth.NetWorth) *net_worth.NetWorth); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*net_worth.NetWorth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, net_worth.NetWorth) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
