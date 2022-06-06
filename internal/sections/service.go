package sections 

import "fmt"
type Service interface {
	GetAll() ([]Section, error)
	GetById(id int) (Section, error)
	Create(sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity,maximumCapacity, warehouseId, productTypeId int) (Section, error)
	Update(id, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity,maximumCapacity, warehouseId, productTypeId int) (Section, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s service) GetAll() ([]Section, error) {
	sectionsList, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return sectionsList, nil
}

func (s service) GetById(id int) (Section, error) {
	section, err := s.repository.GetById(id)
	if err != nil {
		return Section{}, err
	}
	return section, err
}

func (s service) Create(sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId int) (Section, error) {
	sectionsList, err := s.repository.GetAll()
	if err != nil {
		return Section{}, err
	}

	for i := range sectionsList {
		if sectionsList[i].SectionNumber == sectionNumber {
			return Section{}, fmt.Errorf("sectionNumber %v do Section já existe", sectionNumber)
		}
	}

	lastID, err := s.repository.LastID()

	if err != nil {
		return Section{}, err
	}

	lastID++

	section, err := s.repository.Create(lastID, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity,maximumCapacity, warehouseId, productTypeId)

	if err != nil {
		return Section{}, err
	}

	return section, nil

}

func (s service) Update(id, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId int) (Section, error) {
	section, err := s.repository.Update(id, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId)
	if err != nil {
		return Section{}, err
	}
	return section, err
}

func (s service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}