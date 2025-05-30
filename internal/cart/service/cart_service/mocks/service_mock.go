// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mocks

import (
	"cart/internal/cart/service/cart_service"
	"context"

	mock "github.com/stretchr/testify/mock"
)

// NewServiceMock creates a new instance of ServiceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceMock {
	mock := &ServiceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// ServiceMock is an autogenerated mock type for the Service type
type ServiceMock struct {
	mock.Mock
}

type ServiceMock_Expecter struct {
	mock *mock.Mock
}

func (_m *ServiceMock) EXPECT() *ServiceMock_Expecter {
	return &ServiceMock_Expecter{mock: &_m.Mock}
}

// AddItem provides a mock function for the type ServiceMock
func (_mock *ServiceMock) AddItem(ctx context.Context, userID int64, skuID int64, count uint16) error {
	ret := _mock.Called(ctx, userID, skuID, count)

	if len(ret) == 0 {
		panic("no return value specified for AddItem")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, int64, int64, uint16) error); ok {
		r0 = returnFunc(ctx, userID, skuID, count)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// ServiceMock_AddItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddItem'
type ServiceMock_AddItem_Call struct {
	*mock.Call
}

// AddItem is a helper method to define mock.On call
//   - ctx
//   - userID
//   - skuID
//   - count
func (_e *ServiceMock_Expecter) AddItem(ctx interface{}, userID interface{}, skuID interface{}, count interface{}) *ServiceMock_AddItem_Call {
	return &ServiceMock_AddItem_Call{Call: _e.mock.On("AddItem", ctx, userID, skuID, count)}
}

func (_c *ServiceMock_AddItem_Call) Run(run func(ctx context.Context, userID int64, skuID int64, count uint16)) *ServiceMock_AddItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(int64), args[3].(uint16))
	})
	return _c
}

func (_c *ServiceMock_AddItem_Call) Return(err error) *ServiceMock_AddItem_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *ServiceMock_AddItem_Call) RunAndReturn(run func(ctx context.Context, userID int64, skuID int64, count uint16) error) *ServiceMock_AddItem_Call {
	_c.Call.Return(run)
	return _c
}

// Checkout provides a mock function for the type ServiceMock
func (_mock *ServiceMock) Checkout(ctx context.Context, userID int64) (int64, error) {
	ret := _mock.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for Checkout")
	}

	var r0 int64
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, int64) (int64, error)); ok {
		return returnFunc(ctx, userID)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, int64) int64); ok {
		r0 = returnFunc(ctx, userID)
	} else {
		r0 = ret.Get(0).(int64)
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = returnFunc(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// ServiceMock_Checkout_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Checkout'
type ServiceMock_Checkout_Call struct {
	*mock.Call
}

// Checkout is a helper method to define mock.On call
//   - ctx
//   - userID
func (_e *ServiceMock_Expecter) Checkout(ctx interface{}, userID interface{}) *ServiceMock_Checkout_Call {
	return &ServiceMock_Checkout_Call{Call: _e.mock.On("Checkout", ctx, userID)}
}

func (_c *ServiceMock_Checkout_Call) Run(run func(ctx context.Context, userID int64)) *ServiceMock_Checkout_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *ServiceMock_Checkout_Call) Return(orderID int64, err error) *ServiceMock_Checkout_Call {
	_c.Call.Return(orderID, err)
	return _c
}

func (_c *ServiceMock_Checkout_Call) RunAndReturn(run func(ctx context.Context, userID int64) (int64, error)) *ServiceMock_Checkout_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCartByUserId provides a mock function for the type ServiceMock
func (_mock *ServiceMock) DeleteCartByUserId(ctx context.Context, userID int64) error {
	ret := _mock.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCartByUserId")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = returnFunc(ctx, userID)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// ServiceMock_DeleteCartByUserId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCartByUserId'
type ServiceMock_DeleteCartByUserId_Call struct {
	*mock.Call
}

// DeleteCartByUserId is a helper method to define mock.On call
//   - ctx
//   - userID
func (_e *ServiceMock_Expecter) DeleteCartByUserId(ctx interface{}, userID interface{}) *ServiceMock_DeleteCartByUserId_Call {
	return &ServiceMock_DeleteCartByUserId_Call{Call: _e.mock.On("DeleteCartByUserId", ctx, userID)}
}

func (_c *ServiceMock_DeleteCartByUserId_Call) Run(run func(ctx context.Context, userID int64)) *ServiceMock_DeleteCartByUserId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *ServiceMock_DeleteCartByUserId_Call) Return(err error) *ServiceMock_DeleteCartByUserId_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *ServiceMock_DeleteCartByUserId_Call) RunAndReturn(run func(ctx context.Context, userID int64) error) *ServiceMock_DeleteCartByUserId_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteItem provides a mock function for the type ServiceMock
func (_mock *ServiceMock) DeleteItem(ctx context.Context, userID int64, skuID int64) error {
	ret := _mock.Called(ctx, userID, skuID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteItem")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = returnFunc(ctx, userID, skuID)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// ServiceMock_DeleteItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteItem'
type ServiceMock_DeleteItem_Call struct {
	*mock.Call
}

// DeleteItem is a helper method to define mock.On call
//   - ctx
//   - userID
//   - skuID
func (_e *ServiceMock_Expecter) DeleteItem(ctx interface{}, userID interface{}, skuID interface{}) *ServiceMock_DeleteItem_Call {
	return &ServiceMock_DeleteItem_Call{Call: _e.mock.On("DeleteItem", ctx, userID, skuID)}
}

func (_c *ServiceMock_DeleteItem_Call) Run(run func(ctx context.Context, userID int64, skuID int64)) *ServiceMock_DeleteItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(int64))
	})
	return _c
}

func (_c *ServiceMock_DeleteItem_Call) Return(err error) *ServiceMock_DeleteItem_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *ServiceMock_DeleteItem_Call) RunAndReturn(run func(ctx context.Context, userID int64, skuID int64) error) *ServiceMock_DeleteItem_Call {
	_c.Call.Return(run)
	return _c
}

// GetCartByUserID provides a mock function for the type ServiceMock
func (_mock *ServiceMock) GetCartByUserID(ctx context.Context, userID int64) (*cart_service.Cart, error) {
	ret := _mock.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetCartByUserID")
	}

	var r0 *cart_service.Cart
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, int64) (*cart_service.Cart, error)); ok {
		return returnFunc(ctx, userID)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, int64) *cart_service.Cart); ok {
		r0 = returnFunc(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart_service.Cart)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = returnFunc(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// ServiceMock_GetCartByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCartByUserID'
type ServiceMock_GetCartByUserID_Call struct {
	*mock.Call
}

// GetCartByUserID is a helper method to define mock.On call
//   - ctx
//   - userID
func (_e *ServiceMock_Expecter) GetCartByUserID(ctx interface{}, userID interface{}) *ServiceMock_GetCartByUserID_Call {
	return &ServiceMock_GetCartByUserID_Call{Call: _e.mock.On("GetCartByUserID", ctx, userID)}
}

func (_c *ServiceMock_GetCartByUserID_Call) Run(run func(ctx context.Context, userID int64)) *ServiceMock_GetCartByUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *ServiceMock_GetCartByUserID_Call) Return(cart *cart_service.Cart, err error) *ServiceMock_GetCartByUserID_Call {
	_c.Call.Return(cart, err)
	return _c
}

func (_c *ServiceMock_GetCartByUserID_Call) RunAndReturn(run func(ctx context.Context, userID int64) (*cart_service.Cart, error)) *ServiceMock_GetCartByUserID_Call {
	_c.Call.Return(run)
	return _c
}
