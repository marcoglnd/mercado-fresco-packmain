// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/marcoglnd/mercado-fresco-packmain/internal/products/domain"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateNewProduct provides a mock function with given fields: ctx, product
func (_m *Repository) CreateNewProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	ret := _m.Called(ctx, product)

	var r0 *domain.Product
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Product) *domain.Product); ok {
		r0 = rf(ctx, product)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Product) error); ok {
		r1 = rf(ctx, product)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProductBatches provides a mock function with given fields: ctx, batch
func (_m *Repository) CreateProductBatches(ctx context.Context, batch *domain.ProductBatches) (int64, error) {
	ret := _m.Called(ctx, batch)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ProductBatches) int64); ok {
		r0 = rf(ctx, batch)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.ProductBatches) error); ok {
		r1 = rf(ctx, batch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProductRecords provides a mock function with given fields: ctx, record
func (_m *Repository) CreateProductRecords(ctx context.Context, record *domain.ProductRecords) (int64, error) {
	ret := _m.Called(ctx, record)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ProductRecords) int64); ok {
		r0 = rf(ctx, record)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.ProductRecords) error); ok {
		r1 = rf(ctx, record)
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
func (_m *Repository) GetAll(ctx context.Context) (*[]domain.Product, error) {
	ret := _m.Called(ctx)

	var r0 *[]domain.Product
	if rf, ok := ret.Get(0).(func(context.Context) *[]domain.Product); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.Product)
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
func (_m *Repository) GetById(ctx context.Context, id int64) (*domain.Product, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Product
	if rf, ok := ret.Get(0).(func(context.Context, int64) *domain.Product); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
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

// GetProductBatchesById provides a mock function with given fields: ctx, id
func (_m *Repository) GetProductBatchesById(ctx context.Context, id int64) (*domain.ProductBatches, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.ProductBatches
	if rf, ok := ret.Get(0).(func(context.Context, int64) *domain.ProductBatches); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ProductBatches)
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

// GetProductRecordsById provides a mock function with given fields: ctx, id
func (_m *Repository) GetProductRecordsById(ctx context.Context, id int64) (*domain.ProductRecords, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.ProductRecords
	if rf, ok := ret.Get(0).(func(context.Context, int64) *domain.ProductRecords); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ProductRecords)
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

// GetQtyOfRecordsById provides a mock function with given fields: ctx, id
func (_m *Repository) GetQtyOfRecordsById(ctx context.Context, id int64) (*domain.QtyOfRecords, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.QtyOfRecords
	if rf, ok := ret.Get(0).(func(context.Context, int64) *domain.QtyOfRecords); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.QtyOfRecords)
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

// Update provides a mock function with given fields: ctx, product
func (_m *Repository) Update(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	ret := _m.Called(ctx, product)

	var r0 *domain.Product
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Product) *domain.Product); ok {
		r0 = rf(ctx, product)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Product) error); ok {
		r1 = rf(ctx, product)
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
