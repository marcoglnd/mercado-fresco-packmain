// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/marcoglnd/mercado-fresco-packmain/internal/sections/domain"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, section
func (_m *Repository) Create(ctx context.Context, section *domain.Section) (*domain.Section, error) {
	ret := _m.Called(ctx, section)

	var r0 *domain.Section
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Section) *domain.Section); ok {
		r0 = rf(ctx, section)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Section)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Section) error); ok {
		r1 = rf(ctx, section)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *Repository) GetAll(ctx context.Context) (*[]domain.Section, error) {
	ret := _m.Called(ctx)

	var r0 *[]domain.Section
	if rf, ok := ret.Get(0).(func(context.Context) *[]domain.Section); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.Section)
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

// GetById provides a mock function with given fields: ctx, id
func (_m *Repository) GetById(ctx context.Context, id int64) (*domain.Section, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Section
	if rf, ok := ret.Get(0).(func(context.Context, int64) *domain.Section); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Section)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, section
func (_m *Repository) Update(ctx context.Context, section *domain.Section) (*domain.Section, error) {
	ret := _m.Called(ctx, section)

	var r0 *domain.Section
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Section) *domain.Section); ok {
		r0 = rf(ctx, section)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Section)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Section) error); ok {
		r1 = rf(ctx, section)
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
