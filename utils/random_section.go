package utils

import "github.com/marcoglnd/mercado-fresco-packmain/internal/sections/domain"

func CreateRandomSection() domain.Section {
	section := domain.Section{
		ID:                 1,
		SectionNumber:      RandomInt64(),
		CurrentTemperature: RandomFloat64(),
		MinimumTemperature: RandomFloat64(),
		CurrentCapacity:    RandomInt64(),
		MinimumCapacity:    RandomInt64(),
		MaximumCapacity:    RandomInt64(),
		WarehouseId:        RandomInt64(),
		ProductTypeId:      RandomInt64(),
	}
	return section
}

func CreateRandomListSection() []domain.Section {
	var listOfSections []domain.Section
	for i := 1; i <= 5; i++ {
		section := CreateRandomSection()
		section.ID = int64(i)
		listOfSections = append(listOfSections, section)
	}
	return listOfSections
}
