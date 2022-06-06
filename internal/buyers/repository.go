package buyers

import "fmt"

//Repositorio

var buyersList []Buyer = []Buyer{}

var lastID int

type Repository interface {
	GetAll() ([]Buyer, error)
	GetById(id int) (Buyer, error)
	Create(id int, cardNumberId, firstName, lastName string) (Buyer, error)
	LastID() (int, error)
	Update(id int, cardNumberId, firstName, lastName string) (Buyer, error)
	Delete(id int) error
}

type repository struct{}

func (repository) GetAll() ([]Buyer, error) {
	return buyersList, nil
}

func (repository) GetById(id int) (Buyer, error) {
	for i := range buyersList {
		if buyersList[i].ID == id {
			return buyersList[i], nil
		}
	}
	return Buyer{}, fmt.Errorf("Buyer %d não encontrado", id)
}

func (repository) LastID() (int, error) {
	return lastID, nil
}

func (repository) Create(id int, cardNumberId, firstName, lastName string) (Buyer, error) {
	buyer := Buyer{id, cardNumberId, firstName, lastName}
	buyersList = append(buyersList, buyer)
	lastID = buyer.ID
	return buyer, nil
}

func (repository) Update(id int, cardNumberId, firstName, lastName string) (Buyer, error) {
	buyer := Buyer{CardNumberID: cardNumberId, FirstName: firstName, LastName: lastName}
	updated := false
	for i := range buyersList {
		if buyersList[i].ID == id {
			buyer.ID = id
			buyersList[i] = buyer
			updated = true
		}
	}
	if !updated {
		return Buyer{}, fmt.Errorf("Buyer %d não encontrado", id)
	}
	return buyer, nil
}

func (repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range buyersList {
		if buyersList[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Buyer %d nao encontrado", id)
	}

	buyersList = append(buyersList[:index], buyersList[index+1:]...)
	return nil
}

func NewRepository() Repository {
	return &repository{}
}
