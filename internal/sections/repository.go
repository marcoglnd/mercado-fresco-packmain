package sections

import "fmt"

var sectionList []Section = []Section{}

type Repository interface {
	GetAll() ([]Section, error)
	GetById(id int) (Section, error)
	Create(sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId int) (Section, error)
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
	if len(sectionList) == 0 {
		return 1, nil
	}
	lastId := sectionList[len(sectionList)-1].ID + 1
	return lastId, nil
}

func (r *repository) Create(sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId int) (Section, error) {
	sectionsList, err := r.GetAll()
	if err != nil {
		return Section{}, err
	}

	for i := range sectionsList {
		if sectionsList[i].SectionNumber == sectionNumber {
			return Section{}, fmt.Errorf("sectionNumber %v do Section já existe", sectionNumber)
		}
	}

	lastID, err := r.LastID()

	if err != nil {
		return Section{}, err
	}

	for i := range sectionList {
		if sectionList[i].SectionNumber == sectionNumber {
			return Section{}, fmt.Errorf("SectionNumber %d já existe", sectionNumber)
		}
	}
	section := Section{lastID, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId}
	sectionList = append(sectionList, section)
	return section, nil
}

func (repository) Update(
	id, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity,
	maximumCapacity, warehouseId, productTypeId int) (Section, error) {
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
