package sections

import "fmt"

var sectionList []Section = []Section{}

var lastID int

type Repository interface {
	GetAll() ([]Section, error)
	GetById(id int) (Section, error)
	Create(id, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId int) (Section, error)
	LastID() (int, error)
	Update(id, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId int) (Section, error)
	Delete(id int) error
}

type repository struct{}

func (repository) GetAll() ([]Section, error) {
	return sectionList, nil
}

func (repository) GetById(id int) (Section, error) {
	for i := range sectionList {
		if sectionList[i].ID == id {
			return sectionList[i], nil
		}
	}
	return Section{}, fmt.Errorf("Section %d não encontrada", id)
}

func (repository) LastID() (int, error) {
	return lastID, nil
}

func (repository) Create(id, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId int) (Section, error){
	for i := range sectionList {
		if sectionList[i].SectionNumber == sectionNumber {
			return Section{}, fmt.Errorf("SectionNumber %d já existe", sectionNumber)
		}
	}
	section := Section{id, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId}
	sectionList = append(sectionList, section)
	lastID = section.ID
	return section, nil
}

func (repository) Update(
	id, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity,
	maximumCapacity, warehouseId, productTypeId int) (Section, error){
	section := Section{
		SectionNumber:      sectionNumber,
		CurrentTemperature: currentTemperature,
		MinimumTemperature: minimumTemperature,
		CurrentCapacity:    currentCapacity,
		MinimumCapacity:    minimumCapacity,
		MaximumCapacity:    maximumCapacity,
		WarehouseId:        warehouseId,
		ProductTypeId:      productTypeId,
	}
	updated := false
	for i := range sectionList {
		if sectionList[i].ID == id {
			section.ID = id
			sectionList[i] = section
			updated = true
		}
	}
	if !updated {
		return Section{}, fmt.Errorf("Section %d não encontrada", id)
	}
	return section, nil
}

func (repository) Delete(id int) error {
	delete := false
	var index int
	for i := range sectionList {
		if sectionList[i].ID == id {
			index = i
			delete = true
		}
	}
	if !delete {
		return fmt.Errorf("Section %d não encontrada", id)
	}

	sectionList = append(sectionList[:index], sectionList[index+1:]...)
	return nil
}

func NewRepository() Repository {
	return &repository{}
}
