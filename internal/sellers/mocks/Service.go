// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	sellers "github.com/marcoglnd/mercado-fresco-packmain/internal/sellers"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *Service) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Service) GetAll() ([]sellers.Seller, error) {
	ret := _m.Called()

	var r0 []sellers.Seller
	if rf, ok := ret.Get(0).(func() []sellers.Seller); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sellers.Seller)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *Service) GetById(id int) (sellers.Seller, error) {
	ret := _m.Called(id)

	var r0 sellers.Seller
	if rf, ok := ret.Get(0).(func(int) sellers.Seller); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(sellers.Seller)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: cid, company_name, address, telephone
func (_m *Service) Store(cid int, company_name string, address string, telephone string) (sellers.Seller, error) {
	ret := _m.Called(cid, company_name, address, telephone)

	var r0 sellers.Seller
	if rf, ok := ret.Get(0).(func(int, string, string, string) sellers.Seller); ok {
		r0 = rf(cid, company_name, address, telephone)
	} else {
		r0 = ret.Get(0).(sellers.Seller)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string, string, string) error); ok {
		r1 = rf(cid, company_name, address, telephone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, cid, company_name, address, telephone
func (_m *Service) Update(id int, cid int, company_name string, address string, telephone string) (sellers.Seller, error) {
	ret := _m.Called(id, cid, company_name, address, telephone)

	var r0 sellers.Seller
	if rf, ok := ret.Get(0).(func(int, int, string, string, string) sellers.Seller); ok {
		r0 = rf(id, cid, company_name, address, telephone)
	} else {
		r0 = ret.Get(0).(sellers.Seller)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, string, string, string) error); ok {
		r1 = rf(id, cid, company_name, address, telephone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
