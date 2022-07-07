// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	employees "github.com/marcoglnd/mercado-fresco-packmain/internal/employees"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: cardNumberId, firstName, lastName, warehouseId
func (_m *Service) Create(cardNumberId string, firstName string, lastName string, warehouseId int) (employees.Employee, error) {
	ret := _m.Called(cardNumberId, firstName, lastName, warehouseId)

	var r0 employees.Employee
	if rf, ok := ret.Get(0).(func(string, string, string, int) employees.Employee); ok {
		r0 = rf(cardNumberId, firstName, lastName, warehouseId)
	} else {
		r0 = ret.Get(0).(employees.Employee)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, int) error); ok {
		r1 = rf(cardNumberId, firstName, lastName, warehouseId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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
func (_m *Service) GetAll() ([]employees.Employee, error) {
	ret := _m.Called()

	var r0 []employees.Employee
	if rf, ok := ret.Get(0).(func() []employees.Employee); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]employees.Employee)
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
func (_m *Service) GetById(id int) (employees.Employee, error) {
	ret := _m.Called(id)

	var r0 employees.Employee
	if rf, ok := ret.Get(0).(func(int) employees.Employee); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(employees.Employee)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, cardNumberId, firstName, lastName, warehouseId
func (_m *Service) Update(id int, cardNumberId string, firstName string, lastName string, warehouseId int) (employees.Employee, error) {
	ret := _m.Called(id, cardNumberId, firstName, lastName, warehouseId)

	var r0 employees.Employee
	if rf, ok := ret.Get(0).(func(int, string, string, string, int) employees.Employee); ok {
		r0 = rf(id, cardNumberId, firstName, lastName, warehouseId)
	} else {
		r0 = ret.Get(0).(employees.Employee)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string, string, string, int) error); ok {
		r1 = rf(id, cardNumberId, firstName, lastName, warehouseId)
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