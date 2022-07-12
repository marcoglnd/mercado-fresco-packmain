package utils

import "github.com/marcoglnd/mercado-fresco-packmain/internal/warehouses/domain"

func CreateRandomWarehouse() domain.Warehouse {
	warehouse := domain.Warehouse{
		ID:                 0,
		WarehouseCode:      RandomString(3),
		Address:            RandomString(10),
		Telephone:          RandomString(6),
		MinimumCapacity:    int(RandomInt(1, 100)),
		MinimumTemperature: float32(RandomFloat64()),
		LocalityId:         0,
	}
	return warehouse
}

func CreateRandomListWarehouses() []domain.Warehouse {
	var listOfWarehouses []domain.Warehouse
	for i := 1; i <= 5; i++ {
		warehouse := CreateRandomWarehouse()
		warehouse.ID = int64(i)
		listOfWarehouses = append(listOfWarehouses, warehouse)
	}
	return listOfWarehouses
}
