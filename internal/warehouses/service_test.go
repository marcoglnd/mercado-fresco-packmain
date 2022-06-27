package warehouses_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses"
	mock_warehouses "github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock_warehouses.NewMockRepository(ctrl)
	service := warehouses.NewService(repositoryMock)
	warehouseFake := &warehouses.Warehouse{
		WarehouseCode:      "IBC",
		Address:            "Rua Sao Paulo",
		Telephone:          "1130304040",
		MinimumCapacity:    3,
		MinimumTemperature: 10,
	}
	repositoryMock.EXPECT().FindByWarehouseCode("IBC").Return(nil, nil)
	repositoryMock.EXPECT().Create(
		warehouseFake.WarehouseCode,
		warehouseFake.Address,
		warehouseFake.Telephone,
		warehouseFake.MinimumCapacity,
		warehouseFake.MinimumTemperature).Return(warehouseFake, nil)

	warehouse, err := service.Create(
		warehouseFake.WarehouseCode,
		warehouseFake.Address,
		warehouseFake.Telephone,
		warehouseFake.MinimumCapacity,
		warehouseFake.MinimumTemperature)

	assert.Nil(t, err)
	assert.NotNil(t, warehouse)
	assert.Equal(t, warehouseFake, warehouse)
}

func TestCreateConflict(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock_warehouses.NewMockRepository(ctrl)
	service := warehouses.NewService(repositoryMock)
	warehouseFake := &warehouses.Warehouse{
		WarehouseCode:      "IBC",
		Address:            "Rua Sao Paulo",
		Telephone:          "1130304040",
		MinimumCapacity:    3,
		MinimumTemperature: 10,
	}
	repositoryMock.EXPECT().FindByWarehouseCode("IBC").Return(warehouseFake, nil)

	_, err := service.Create(
		warehouseFake.WarehouseCode,
		warehouseFake.Address,
		warehouseFake.Telephone,
		warehouseFake.MinimumCapacity,
		warehouseFake.MinimumTemperature)

	assert.NotNil(t, err)
	assert.Equal(t, "warehouseCode already exists", err.Error())
}

func TestFindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock_warehouses.NewMockRepository(ctrl)
	service := warehouses.NewService(repositoryMock)
	warehousesFake := []warehouses.Warehouse{
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
	repositoryMock.EXPECT().GetAll().Return(warehousesFake, nil)
	warehouses, err := service.GetAll()

	assert.Nil(t, err)
	assert.NotNil(t, warehouses)
	assert.Equal(t, len(warehouses), len(warehousesFake))
}

func TestFindByIdNonExistent(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock_warehouses.NewMockRepository(ctrl)
	service := warehouses.NewService(repositoryMock)
	warehouseFakeId := 10
	repositoryMock.EXPECT().FindById(warehouseFakeId).Return(nil, nil)
	warehouse, err := service.FindById(warehouseFakeId)

	assert.NotNil(t, err)
	assert.Nil(t, warehouse)
}

func TestFindByIdExistent(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock_warehouses.NewMockRepository(ctrl)
	service := warehouses.NewService(repositoryMock)
	warehouseFake := &warehouses.Warehouse{
		ID:                 1,
		WarehouseCode:      "IBC",
		Address:            "Rua Sao Paulo",
		Telephone:          "1130304040",
		MinimumCapacity:    3,
		MinimumTemperature: 10,
	}
	repositoryMock.EXPECT().FindById(warehouseFake.ID).Return(warehouseFake, nil)
	warehouse, err := service.FindById(warehouseFake.ID)

	assert.Nil(t, err)
	assert.NotNil(t, warehouse)
	assert.Equal(t, warehouse.ID, warehouseFake.ID)
}

func TestUpdateExistent(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock_warehouses.NewMockRepository(ctrl)
	currentWarehouseFake := &warehouses.Warehouse{
		ID:                 1,
		WarehouseCode:      "IBC",
		Address:            "Rua Sao Paulo",
		Telephone:          "1130304040",
		MinimumCapacity:    3,
		MinimumTemperature: 10,
	}
	newWarehouseFake := &warehouses.Warehouse{
		ID:                 1,
		WarehouseCode:      "BRU",
		Address:            "Rua Sao Paulo",
		Telephone:          "1130304040",
		MinimumCapacity:    3,
		MinimumTemperature: 10,
	}
	repositoryMock.EXPECT().Update(newWarehouseFake).Return(nil)
	repositoryMock.EXPECT().FindById(currentWarehouseFake.ID).Return(currentWarehouseFake, nil)
	repositoryMock.EXPECT().FindByWarehouseCode(newWarehouseFake.WarehouseCode).Return(nil, nil)
	service := warehouses.NewService(repositoryMock)
	warehouse, err := service.Update(
		currentWarehouseFake.ID,
		newWarehouseFake.WarehouseCode,
		newWarehouseFake.Address,
		newWarehouseFake.Telephone,
		newWarehouseFake.MinimumCapacity,
		newWarehouseFake.MinimumTemperature,
	)

	assert.Nil(t, err)
	assert.NotNil(t, warehouse)
	assert.Equal(t, newWarehouseFake.WarehouseCode, warehouse.WarehouseCode)
}

func TestUpdateNonExistent(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock_warehouses.NewMockRepository(ctrl)
	currentWarehouseFakeId := 10
	newWarehouseFake := &warehouses.Warehouse{
		ID:                 1,
		WarehouseCode:      "BRU",
		Address:            "Rua Sao Paulo",
		Telephone:          "1130304040",
		MinimumCapacity:    3,
		MinimumTemperature: 10,
	}
	repositoryMock.EXPECT().FindById(currentWarehouseFakeId).Return(nil, fmt.Errorf("id is inexistent"))
	service := warehouses.NewService(repositoryMock)
	warehouse, err := service.Update(
		currentWarehouseFakeId,
		newWarehouseFake.WarehouseCode,
		newWarehouseFake.Address,
		newWarehouseFake.Telephone,
		newWarehouseFake.MinimumCapacity,
		newWarehouseFake.MinimumTemperature,
	)

	assert.NotNil(t, err)
	assert.Nil(t, warehouse)
}

func TestDeleteNonExistent(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock_warehouses.NewMockRepository(ctrl)
	warehouseFakeId := 10
	repositoryMock.EXPECT().FindById(warehouseFakeId).Return(nil, fmt.Errorf("id is inexistent"))
	service := warehouses.NewService(repositoryMock)
	err := service.Delete(warehouseFakeId)

	assert.NotNil(t, err)
}

func TestDeleteOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock_warehouses.NewMockRepository(ctrl)
	warehouseFake := &warehouses.Warehouse{
		ID:                 1,
		WarehouseCode:      "IBC",
		Address:            "Rua Sao Paulo",
		Telephone:          "1130304040",
		MinimumCapacity:    3,
		MinimumTemperature: 10,
	}
	repositoryMock.EXPECT().FindById(warehouseFake.ID).Return(warehouseFake, nil)
	repositoryMock.EXPECT().Delete(warehouseFake.ID).Return(nil)
	service := warehouses.NewService(repositoryMock)
	err := service.Delete(warehouseFake.ID)

	assert.Nil(t, err)
}
