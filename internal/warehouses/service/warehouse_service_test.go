package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses/domain"
	mock "github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses/service"
	"github.com/stretchr/testify/assert"
)

func TestCreateOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockWarehouseRepository(ctrl)
	service := service.NewWarehouseService(repositoryMock)
	ctx := context.TODO()
	warehouseFake := &domain.Warehouse{
		WarehouseCode:      "IBC",
		Address:            "Rua Sao Paulo",
		Telephone:          "1130304040",
		MinimumCapacity:    3,
		MinimumTemperature: 10,
	}
	repositoryMock.EXPECT().FindByWarehouseCode(ctx, "IBC").Return(nil, nil)
	repositoryMock.EXPECT().Create(ctx, warehouseFake).Return(warehouseFake, nil)

	warehouse, err := service.Create(ctx, warehouseFake)

	assert.Nil(t, err)
	assert.NotNil(t, warehouse)
	assert.Equal(t, warehouseFake, warehouse)
}

func TestCreateConflict(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockWarehouseRepository(ctrl)
	service := service.NewWarehouseService(repositoryMock)
	ctx := context.TODO()
	warehouseFake := &domain.Warehouse{
		WarehouseCode:      "IBC",
		Address:            "Rua Sao Paulo",
		Telephone:          "1130304040",
		MinimumCapacity:    3,
		MinimumTemperature: 10,
	}
	repositoryMock.EXPECT().FindByWarehouseCode(ctx, "IBC").Return(warehouseFake, nil)

	_, err := service.Create(ctx, warehouseFake)

	assert.NotNil(t, err)
	assert.Equal(t, "warehouseCode already exists", err.Error())
}

func TestFindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockWarehouseRepository(ctrl)
	service := service.NewWarehouseService(repositoryMock)
	ctx := context.TODO()
	warehousesFake := []domain.Warehouse{
		{
			WarehouseCode:      "IBC",
			Address:            "Rua Sao Paulo",
			Telephone:          "1130304040",
			MinimumCapacity:    3,
			MinimumTemperature: 10,
		},
		{
			WarehouseCode:      "FRE",
			Address:            "Rua dos Mercado Fresco",
			Telephone:          "1130305040",
			MinimumCapacity:    5,
			MinimumTemperature: 10,
		},
	}
	repositoryMock.EXPECT().GetAll(ctx).Return(&warehousesFake, nil)
	warehouses, err := service.GetAll(ctx)

	assert.Nil(t, err)
	assert.NotNil(t, warehouses)
	assert.Equal(t, len(*warehouses), len(warehousesFake))
}

func TestFindByIdNonExistent(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockWarehouseRepository(ctrl)
	service := service.NewWarehouseService(repositoryMock)
	ctx := context.TODO()
	warehouseFakeId := 10
	repositoryMock.EXPECT().FindById(ctx, warehouseFakeId).Return(nil, nil)
	warehouse, err := service.FindById(ctx, warehouseFakeId)

	assert.NotNil(t, err)
	assert.Nil(t, warehouse)
}

func TestFindByIdExistent(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockWarehouseRepository(ctrl)
	service := service.NewWarehouseService(repositoryMock)
	ctx := context.TODO()
	warehouseFake := domain.Warehouse{
		ID:                 1,
		WarehouseCode:      "IBC",
		Address:            "Rua Sao Paulo",
		Telephone:          "1130304040",
		MinimumCapacity:    3,
		MinimumTemperature: 10,
	}
	repositoryMock.EXPECT().FindById(ctx, warehouseFake.ID).Return(&warehouseFake, nil)
	warehouse, err := service.FindById(ctx, warehouseFake.ID)

	assert.Nil(t, err)
	assert.NotNil(t, warehouse)
	assert.Equal(t, warehouse.ID, warehouseFake.ID)
}

func TestUpdateExistent(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockWarehouseRepository(ctrl)
	service := service.NewWarehouseService(repositoryMock)
	ctx := context.TODO()
	currentWarehouseFake := &domain.Warehouse{
		ID:                 1,
		WarehouseCode:      "BRU",
		Address:            "Rua Sao Paulo",
		Telephone:          "1130304040",
		MinimumCapacity:    3,
		MinimumTemperature: 10,
	}
	updatedWarehouseFake := &domain.Warehouse{
		ID:                 1,
		WarehouseCode:      "PRE",
		Address:            "Rua Sao Paulo 2",
		Telephone:          "1130304042",
		MinimumCapacity:    2,
		MinimumTemperature: 12,
	}
	repositoryMock.EXPECT().Update(ctx, currentWarehouseFake).Return(nil)
	repositoryMock.EXPECT().FindById(ctx, updatedWarehouseFake.ID).Return(currentWarehouseFake, nil)
	repositoryMock.EXPECT().FindByWarehouseCode(ctx, updatedWarehouseFake.WarehouseCode).Return(nil, nil)
	warehouse, err := service.Update(ctx, updatedWarehouseFake)

	assert.Nil(t, err)
	assert.NotNil(t, warehouse)
	assert.Equal(t, updatedWarehouseFake.WarehouseCode, warehouse.WarehouseCode)
	assert.Equal(t, updatedWarehouseFake.Address, warehouse.Address)
	assert.Equal(t, updatedWarehouseFake.Telephone, warehouse.Telephone)
	assert.Equal(t, updatedWarehouseFake.MinimumCapacity, warehouse.MinimumCapacity)
	assert.Equal(t, updatedWarehouseFake.MinimumTemperature, warehouse.MinimumTemperature)
}

func TestUpdateNonExistent(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockWarehouseRepository(ctrl)
	service := service.NewWarehouseService(repositoryMock)
	ctx := context.TODO()
	warehouseFake := domain.Warehouse{
		ID:                 1,
		WarehouseCode:      "BRU",
		Address:            "Rua Sao Paulo",
		Telephone:          "1130304040",
		MinimumCapacity:    3,
		MinimumTemperature: 10,
	}
	repositoryMock.EXPECT().FindById(ctx, warehouseFake.ID).Return(nil, fmt.Errorf("id is inexistent"))
	warehouse, err := service.Update(ctx, &warehouseFake)

	assert.NotNil(t, err)
	assert.Nil(t, warehouse)
}

func TestDeleteNonExistent(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockWarehouseRepository(ctrl)
	service := service.NewWarehouseService(repositoryMock)
	ctx := context.TODO()
	warehouseFakeId := 10
	repositoryMock.EXPECT().FindById(ctx, warehouseFakeId).Return(nil, fmt.Errorf("id is inexistent"))
	err := service.Delete(ctx, warehouseFakeId)

	assert.NotNil(t, err)
}

func TestDeleteOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockWarehouseRepository(ctrl)
	service := service.NewWarehouseService(repositoryMock)
	ctx := context.TODO()
	warehouseFake := domain.Warehouse{
		ID:                 1,
		WarehouseCode:      "IBC",
		Address:            "Rua Sao Paulo",
		Telephone:          "1130304040",
		MinimumCapacity:    3,
		MinimumTemperature: 10,
	}
	repositoryMock.EXPECT().FindById(ctx, warehouseFake.ID).Return(&warehouseFake, nil)
	repositoryMock.EXPECT().Delete(ctx, warehouseFake.ID).Return(nil)

	err := service.Delete(ctx, warehouseFake.ID)

	assert.Nil(t, err)
}

func TestFindByWarehouseCodeOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockWarehouseRepository(ctrl)
	service := service.NewWarehouseService(repositoryMock)
	ctx := context.TODO()
	warehouseFake := domain.Warehouse{
		ID:                 1,
		WarehouseCode:      "IBC",
		Address:            "Rua Sao Paulo",
		Telephone:          "1130304040",
		MinimumCapacity:    3,
		MinimumTemperature: 10,
	}
	repositoryMock.EXPECT().FindByWarehouseCode(ctx, warehouseFake.WarehouseCode).Return(&warehouseFake, nil)

	warehouse, err := service.FindByWarehouseCode(ctx, warehouseFake.WarehouseCode)

	assert.Nil(t, err)
	assert.NotNil(t, warehouse)
}
