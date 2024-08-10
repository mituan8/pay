// Code generated by mockery v2.32.0. DO NOT EDIT.

package mock

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	wallet "github.com/mituan8/pay/pkg/api-kms/v1/client/wallet"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// CreateBSCTransaction provides a mock function with given fields: params, opts
func (_m *ClientService) CreateBSCTransaction(params *wallet.CreateBSCTransactionParams, opts ...wallet.ClientOption) (*wallet.CreateBSCTransactionCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *wallet.CreateBSCTransactionCreated
	var r1 error
	if rf, ok := ret.Get(0).(func(*wallet.CreateBSCTransactionParams, ...wallet.ClientOption) (*wallet.CreateBSCTransactionCreated, error)); ok {
		return rf(params, opts...)
	}
	if rf, ok := ret.Get(0).(func(*wallet.CreateBSCTransactionParams, ...wallet.ClientOption) *wallet.CreateBSCTransactionCreated); ok {
		r0 = rf(params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wallet.CreateBSCTransactionCreated)
		}
	}

	if rf, ok := ret.Get(1).(func(*wallet.CreateBSCTransactionParams, ...wallet.ClientOption) error); ok {
		r1 = rf(params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateEthereumTransaction provides a mock function with given fields: params, opts
func (_m *ClientService) CreateEthereumTransaction(params *wallet.CreateEthereumTransactionParams, opts ...wallet.ClientOption) (*wallet.CreateEthereumTransactionCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *wallet.CreateEthereumTransactionCreated
	var r1 error
	if rf, ok := ret.Get(0).(func(*wallet.CreateEthereumTransactionParams, ...wallet.ClientOption) (*wallet.CreateEthereumTransactionCreated, error)); ok {
		return rf(params, opts...)
	}
	if rf, ok := ret.Get(0).(func(*wallet.CreateEthereumTransactionParams, ...wallet.ClientOption) *wallet.CreateEthereumTransactionCreated); ok {
		r0 = rf(params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wallet.CreateEthereumTransactionCreated)
		}
	}

	if rf, ok := ret.Get(1).(func(*wallet.CreateEthereumTransactionParams, ...wallet.ClientOption) error); ok {
		r1 = rf(params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateMaticTransaction provides a mock function with given fields: params, opts
func (_m *ClientService) CreateMaticTransaction(params *wallet.CreateMaticTransactionParams, opts ...wallet.ClientOption) (*wallet.CreateMaticTransactionCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *wallet.CreateMaticTransactionCreated
	var r1 error
	if rf, ok := ret.Get(0).(func(*wallet.CreateMaticTransactionParams, ...wallet.ClientOption) (*wallet.CreateMaticTransactionCreated, error)); ok {
		return rf(params, opts...)
	}
	if rf, ok := ret.Get(0).(func(*wallet.CreateMaticTransactionParams, ...wallet.ClientOption) *wallet.CreateMaticTransactionCreated); ok {
		r0 = rf(params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wallet.CreateMaticTransactionCreated)
		}
	}

	if rf, ok := ret.Get(1).(func(*wallet.CreateMaticTransactionParams, ...wallet.ClientOption) error); ok {
		r1 = rf(params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTronTransaction provides a mock function with given fields: params, opts
func (_m *ClientService) CreateTronTransaction(params *wallet.CreateTronTransactionParams, opts ...wallet.ClientOption) (*wallet.CreateTronTransactionCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *wallet.CreateTronTransactionCreated
	var r1 error
	if rf, ok := ret.Get(0).(func(*wallet.CreateTronTransactionParams, ...wallet.ClientOption) (*wallet.CreateTronTransactionCreated, error)); ok {
		return rf(params, opts...)
	}
	if rf, ok := ret.Get(0).(func(*wallet.CreateTronTransactionParams, ...wallet.ClientOption) *wallet.CreateTronTransactionCreated); ok {
		r0 = rf(params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wallet.CreateTronTransactionCreated)
		}
	}

	if rf, ok := ret.Get(1).(func(*wallet.CreateTronTransactionParams, ...wallet.ClientOption) error); ok {
		r1 = rf(params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateWallet provides a mock function with given fields: params, opts
func (_m *ClientService) CreateWallet(params *wallet.CreateWalletParams, opts ...wallet.ClientOption) (*wallet.CreateWalletCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *wallet.CreateWalletCreated
	var r1 error
	if rf, ok := ret.Get(0).(func(*wallet.CreateWalletParams, ...wallet.ClientOption) (*wallet.CreateWalletCreated, error)); ok {
		return rf(params, opts...)
	}
	if rf, ok := ret.Get(0).(func(*wallet.CreateWalletParams, ...wallet.ClientOption) *wallet.CreateWalletCreated); ok {
		r0 = rf(params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wallet.CreateWalletCreated)
		}
	}

	if rf, ok := ret.Get(1).(func(*wallet.CreateWalletParams, ...wallet.ClientOption) error); ok {
		r1 = rf(params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteWallet provides a mock function with given fields: params, opts
func (_m *ClientService) DeleteWallet(params *wallet.DeleteWalletParams, opts ...wallet.ClientOption) (*wallet.DeleteWalletNoContent, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *wallet.DeleteWalletNoContent
	var r1 error
	if rf, ok := ret.Get(0).(func(*wallet.DeleteWalletParams, ...wallet.ClientOption) (*wallet.DeleteWalletNoContent, error)); ok {
		return rf(params, opts...)
	}
	if rf, ok := ret.Get(0).(func(*wallet.DeleteWalletParams, ...wallet.ClientOption) *wallet.DeleteWalletNoContent); ok {
		r0 = rf(params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wallet.DeleteWalletNoContent)
		}
	}

	if rf, ok := ret.Get(1).(func(*wallet.DeleteWalletParams, ...wallet.ClientOption) error); ok {
		r1 = rf(params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWallet provides a mock function with given fields: params, opts
func (_m *ClientService) GetWallet(params *wallet.GetWalletParams, opts ...wallet.ClientOption) (*wallet.GetWalletOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *wallet.GetWalletOK
	var r1 error
	if rf, ok := ret.Get(0).(func(*wallet.GetWalletParams, ...wallet.ClientOption) (*wallet.GetWalletOK, error)); ok {
		return rf(params, opts...)
	}
	if rf, ok := ret.Get(0).(func(*wallet.GetWalletParams, ...wallet.ClientOption) *wallet.GetWalletOK); ok {
		r0 = rf(params, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wallet.GetWalletOK)
		}
	}

	if rf, ok := ret.Get(1).(func(*wallet.GetWalletParams, ...wallet.ClientOption) error); ok {
		r1 = rf(params, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *ClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// NewClientService creates a new instance of ClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientService {
	mock := &ClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
