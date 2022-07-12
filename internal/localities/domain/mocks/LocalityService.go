// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/marcoglnd/mercado-fresco-packmain/internal/localities/domain"
	mock "github.com/stretchr/testify/mock"
)

// LocalityService is an autogenerated mock type for the LocalityService type
type LocalityService struct {
	mock.Mock
}

// CreateLocality provides a mock function with given fields: ctx, local
func (_m *LocalityService) CreateLocality(ctx context.Context, local *domain.Locality) (int64, error) {
	ret := _m.Called(ctx, local)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Locality) int64); ok {
		r0 = rf(ctx, local)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Locality) error); ok {
		r1 = rf(ctx, local)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLocalityByID provides a mock function with given fields: ctx, id
func (_m *LocalityService) GetLocalityByID(ctx context.Context, id int64) (*domain.GetLocality, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.GetLocality
	if rf, ok := ret.Get(0).(func(context.Context, int64) *domain.GetLocality); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.GetLocality)
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

// GetQtyOfSellers provides a mock function with given fields: ctx
func (_m *LocalityService) GetAllQtyOfSellers(ctx context.Context) (*[]domain.QtyOfSellers, error) {
	ret := _m.Called(ctx)

	var r0 *[]domain.QtyOfSellers
	if rf, ok := ret.Get(0).(func(context.Context) *[]domain.QtyOfSellers); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.QtyOfSellers)
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

// GetQtyOfSellersByLocalityId provides a mock function with given fields: ctx, id
func (_m *LocalityService) GetQtyOfSellersByLocalityId(ctx context.Context, id int64) (*domain.QtyOfSellers, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.QtyOfSellers
	if rf, ok := ret.Get(0).(func(context.Context, int64) *domain.QtyOfSellers); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.QtyOfSellers)
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

type mockConstructorTestingTNewLocalityService interface {
	mock.TestingT
	Cleanup(func())
}

// NewLocalityService creates a new instance of LocalityService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLocalityService(t mockConstructorTestingTNewLocalityService) *LocalityService {
	mock := &LocalityService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
