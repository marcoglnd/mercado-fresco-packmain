package sellers

import "fmt"

// Repositório -> lista de produtos + id

var sr []Seller = []Seller{}

var lastID int

// Este repositório é uma interface, portanto tem alguns métodos
type Repository interface {
	GetAll() ([]Seller, error)
	GetById(id int) (Seller, error)
	Store(id int, cid int, company_name string, address string, telephone string) (Seller, error)
	LastID() (int, error)
	Update(id int, cid int, company_name string, address string, telephone string) (Seller, error)
	Delete(id int) error
}

// Criamos uma struct repository que irá implementar a interface

type repository struct{}

func (repository) GetAll() ([]Seller, error) {
	return sr, nil
}

func (repository) GetById(id int) (Seller, error) {
	var s Seller
	foundEmployee := false
	for i := range sr {
		if sr[i].ID == id {
			s = sr[i]
			foundEmployee = true
		}
	}
	if !foundEmployee {
		return Seller{}, fmt.Errorf("Seller %d não encontrado", id)
	}
	return s, nil
}

func (repository) LastID() (int, error) {
	return lastID, nil
}

func (repository) Store(id int, cid int, company_name string, address string, telephone string) (Seller, error) {
	p := Seller{id, cid, company_name, address, telephone}
	sr = append(sr, p)
	lastID = p.ID
	return p, nil
}

func (repository) Update(id int, cid int, company_name string, address string, telephone string) (Seller, error) {
	p := Seller{Cid: cid, Company_name: company_name, Address: address, Telephone: telephone}
	updated := false
	for i := range sr {
		if sr[i].ID == id {
			p.ID = id
			sr[i] = p
			updated = true
		}
	}
	if !updated {
		return Seller{}, fmt.Errorf("seller %d não encontrado", id)
	}
	return p, nil
}

func (repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range sr {
		if sr[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("seller %d não encontrado", id)
	}

	sr = append(sr[:index], sr[index+1:]...)
	return nil
}

// Método criado para instanciar o novo repositório

func NewRepository() Repository {
	return &repository{}
}
