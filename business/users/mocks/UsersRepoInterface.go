// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	users "github.com/daffaalex22/seleksi-deall/business/users"
	mock "github.com/stretchr/testify/mock"
)

// UsersRepoInterface is an autogenerated mock type for the UsersRepoInterface type
type UsersRepoInterface struct {
	mock.Mock
}

// UsersAdd provides a mock function with given fields: ctx, domain
func (_m *UsersRepoInterface) UsersAdd(ctx context.Context, domain users.Domain) (users.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, users.Domain) users.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, users.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UsersDelete provides a mock function with given fields: ctx, id
func (_m *UsersRepoInterface) UsersDelete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UsersGetAll provides a mock function with given fields: ctx
func (_m *UsersRepoInterface) UsersGetAll(ctx context.Context) ([]users.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []users.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []users.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UsersGetByID provides a mock function with given fields: ctx, id
func (_m *UsersRepoInterface) UsersGetByID(ctx context.Context, id string) (users.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) users.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UsersLogin provides a mock function with given fields: ctx, domain
func (_m *UsersRepoInterface) UsersLogin(ctx context.Context, domain users.Domain) (users.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, users.Domain) users.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, users.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UsersUpdate provides a mock function with given fields: ctx, domain
func (_m *UsersRepoInterface) UsersUpdate(ctx context.Context, domain users.Domain) (users.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, users.Domain) users.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, users.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
