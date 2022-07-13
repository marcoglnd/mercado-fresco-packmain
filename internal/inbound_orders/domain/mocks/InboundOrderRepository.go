// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/marcoglnd/mercado-fresco-packmain/internal/inbound_orders/domain"
	mock "github.com/stretchr/testify/mock"
)

// InboundOrderRepository is an autogenerated mock type for the InboundOrderRepository type
type InboundOrderRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, inboundOrder
func (_m *InboundOrderRepository) Create(ctx context.Context, inboundOrder *domain.InboundOrder) (*domain.InboundOrder, error) {
	ret := _m.Called(ctx, inboundOrder)

	var r0 *domain.InboundOrder
	if rf, ok := ret.Get(0).(func(context.Context, *domain.InboundOrder) *domain.InboundOrder); ok {
		r0 = rf(ctx, inboundOrder)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.InboundOrder)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.InboundOrder) error); ok {
		r1 = rf(ctx, inboundOrder)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *InboundOrderRepository) GetAll(ctx context.Context) (*[]domain.InboundOrder, error) {
	ret := _m.Called(ctx)

	var r0 *[]domain.InboundOrder
	if rf, ok := ret.Get(0).(func(context.Context) *[]domain.InboundOrder); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.InboundOrder)
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

type mockConstructorTestingTNewInboundOrderRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewInboundOrderRepository creates a new instance of InboundOrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInboundOrderRepository(t mockConstructorTestingTNewInboundOrderRepository) *InboundOrderRepository {
	mock := &InboundOrderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
