// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	buyers "github.com/marcoglnd/mercado-fresco-packmain/internal/buyers"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: cardNumberId, firstName, lastName
func (_m *Repository) Create(cardNumberId string, firstName string, lastName string) (buyers.Buyer, error) {
	ret := _m.Called(cardNumberId, firstName, lastName)

	var r0 buyers.Buyer
	if rf, ok := ret.Get(0).(func(string, string, string) buyers.Buyer); ok {
		r0 = rf(cardNumberId, firstName, lastName)
	} else {
		r0 = ret.Get(0).(buyers.Buyer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(cardNumberId, firstName, lastName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id int) error {
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
func (_m *Repository) GetAll() ([]buyers.Buyer, error) {
	ret := _m.Called()

	var r0 []buyers.Buyer
	if rf, ok := ret.Get(0).(func() []buyers.Buyer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]buyers.Buyer)
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
func (_m *Repository) GetById(id int) (buyers.Buyer, error) {
	ret := _m.Called(id)

	var r0 buyers.Buyer
	if rf, ok := ret.Get(0).(func(int) buyers.Buyer); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(buyers.Buyer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LastID provides a mock function with given fields:
func (_m *Repository) LastID() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, cardNumberId, firstName, lastName
func (_m *Repository) Update(id int, cardNumberId string, firstName string, lastName string) (buyers.Buyer, error) {
	ret := _m.Called(id, cardNumberId, firstName, lastName)

	var r0 buyers.Buyer
	if rf, ok := ret.Get(0).(func(int, string, string, string) buyers.Buyer); ok {
		r0 = rf(id, cardNumberId, firstName, lastName)
	} else {
		r0 = ret.Get(0).(buyers.Buyer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string, string, string) error); ok {
		r1 = rf(id, cardNumberId, firstName, lastName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}