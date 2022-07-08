// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/marcoglnd/mercado-fresco-packmain/internal/sellers/domain"
	mock "github.com/stretchr/testify/mock"
)

// SellerRepository is an autogenerated mock type for the SellerRepository type
type SellerRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, seller
func (_m *SellerRepository) Create(ctx context.Context, seller *domain.Seller) (*domain.Seller, error) {
	ret := _m.Called(ctx, seller)

	var r0 *domain.Seller
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Seller) *domain.Seller); ok {
		r0 = rf(ctx, seller)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Seller)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Seller) error); ok {
		r1 = rf(ctx, seller)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateLocality provides a mock function with given fields: ctx, record
func (_m *SellerRepository) CreateLocality(ctx context.Context, record *domain.Locality) (int64, error) {
	ret := _m.Called(ctx, record)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Locality) int64); ok {
		r0 = rf(ctx, record)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Locality) error); ok {
		r1 = rf(ctx, record)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *SellerRepository) Delete(ctx context.Context, id int64) error {
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
func (_m *SellerRepository) GetAll(ctx context.Context) (*[]domain.Seller, error) {
	ret := _m.Called(ctx)

	var r0 *[]domain.Seller
	if rf, ok := ret.Get(0).(func(context.Context) *[]domain.Seller); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.Seller)
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

// GetByID provides a mock function with given fields: ctx, id
func (_m *SellerRepository) GetByID(ctx context.Context, id int64) (*domain.Seller, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Seller
	if rf, ok := ret.Get(0).(func(context.Context, int64) *domain.Seller); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Seller)
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

// Update provides a mock function with given fields: ctx, seller
func (_m *SellerRepository) Update(ctx context.Context, seller *domain.Seller) (*domain.Seller, error) {
	ret := _m.Called(ctx, seller)

	var r0 *domain.Seller
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Seller) *domain.Seller); ok {
		r0 = rf(ctx, seller)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Seller)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Seller) error); ok {
		r1 = rf(ctx, seller)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSellerRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewSellerRepository creates a new instance of SellerRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSellerRepository(t mockConstructorTestingTNewSellerRepository) *SellerRepository {
	mock := &SellerRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
