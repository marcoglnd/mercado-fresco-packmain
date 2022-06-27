// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go

// Package mock_warehouses is a generated GoMock package.
package mock_warehouses

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	warehouses "github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockService) Create(warehouseCode, address, telephone string, minimumCapacity, minimumTemperature int) (*warehouses.Warehouse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", warehouseCode, address, telephone, minimumCapacity, minimumTemperature)
	ret0, _ := ret[0].(*warehouses.Warehouse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockServiceMockRecorder) Create(warehouseCode, address, telephone, minimumCapacity, minimumTemperature interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), warehouseCode, address, telephone, minimumCapacity, minimumTemperature)
}

// Delete mocks base method.
func (m *MockService) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockService)(nil).Delete), id)
}

// FindById mocks base method.
func (m *MockService) FindById(id int) (*warehouses.Warehouse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(*warehouses.Warehouse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockServiceMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockService)(nil).FindById), id)
}

// FindByWarehouseCode mocks base method.
func (m *MockService) FindByWarehouseCode(warehouseCode string) (*warehouses.Warehouse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByWarehouseCode", warehouseCode)
	ret0, _ := ret[0].(*warehouses.Warehouse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByWarehouseCode indicates an expected call of FindByWarehouseCode.
func (mr *MockServiceMockRecorder) FindByWarehouseCode(warehouseCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByWarehouseCode", reflect.TypeOf((*MockService)(nil).FindByWarehouseCode), warehouseCode)
}

// GetAll mocks base method.
func (m *MockService) GetAll() ([]warehouses.Warehouse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]warehouses.Warehouse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockService)(nil).GetAll))
}

// IsWarehouseCodeAvailable mocks base method.
func (m *MockService) IsWarehouseCodeAvailable(warehouseCode string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsWarehouseCodeAvailable", warehouseCode)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsWarehouseCodeAvailable indicates an expected call of IsWarehouseCodeAvailable.
func (mr *MockServiceMockRecorder) IsWarehouseCodeAvailable(warehouseCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsWarehouseCodeAvailable", reflect.TypeOf((*MockService)(nil).IsWarehouseCodeAvailable), warehouseCode)
}

// Update mocks base method.
func (m *MockService) Update(warehouseId int, warehouseCode, address, telephone string, minimumCapacity, minimumTemperature int) (*warehouses.Warehouse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", warehouseId, warehouseCode, address, telephone, minimumCapacity, minimumTemperature)
	ret0, _ := ret[0].(*warehouses.Warehouse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceMockRecorder) Update(warehouseId, warehouseCode, address, telephone, minimumCapacity, minimumTemperature interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockService)(nil).Update), warehouseId, warehouseCode, address, telephone, minimumCapacity, minimumTemperature)
}
