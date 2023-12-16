// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	auth "flamingo.me/flamingo/v3/core/auth"

	domain "flamingo.me/flamingo-commerce/v3/customer/domain"

	mock "github.com/stretchr/testify/mock"
)

// CustomerIdentityService is an autogenerated mock type for the CustomerIdentityService type
type CustomerIdentityService struct {
	mock.Mock
}

type CustomerIdentityService_Expecter struct {
	mock *mock.Mock
}

func (_m *CustomerIdentityService) EXPECT() *CustomerIdentityService_Expecter {
	return &CustomerIdentityService_Expecter{mock: &_m.Mock}
}

// GetByIdentity provides a mock function with given fields: ctx, identity
func (_m *CustomerIdentityService) GetByIdentity(ctx context.Context, identity auth.Identity) (domain.Customer, error) {
	ret := _m.Called(ctx, identity)

	var r0 domain.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, auth.Identity) (domain.Customer, error)); ok {
		return rf(ctx, identity)
	}
	if rf, ok := ret.Get(0).(func(context.Context, auth.Identity) domain.Customer); ok {
		r0 = rf(ctx, identity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, auth.Identity) error); ok {
		r1 = rf(ctx, identity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CustomerIdentityService_GetByIdentity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByIdentity'
type CustomerIdentityService_GetByIdentity_Call struct {
	*mock.Call
}

// GetByIdentity is a helper method to define mock.On call
//   - ctx context.Context
//   - identity auth.Identity
func (_e *CustomerIdentityService_Expecter) GetByIdentity(ctx interface{}, identity interface{}) *CustomerIdentityService_GetByIdentity_Call {
	return &CustomerIdentityService_GetByIdentity_Call{Call: _e.mock.On("GetByIdentity", ctx, identity)}
}

func (_c *CustomerIdentityService_GetByIdentity_Call) Run(run func(ctx context.Context, identity auth.Identity)) *CustomerIdentityService_GetByIdentity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(auth.Identity))
	})
	return _c
}

func (_c *CustomerIdentityService_GetByIdentity_Call) Return(_a0 domain.Customer, _a1 error) *CustomerIdentityService_GetByIdentity_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CustomerIdentityService_GetByIdentity_Call) RunAndReturn(run func(context.Context, auth.Identity) (domain.Customer, error)) *CustomerIdentityService_GetByIdentity_Call {
	_c.Call.Return(run)
	return _c
}

// NewCustomerIdentityService creates a new instance of CustomerIdentityService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCustomerIdentityService(t interface {
	mock.TestingT
	Cleanup(func())
}) *CustomerIdentityService {
	mock := &CustomerIdentityService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
