package sellers

// Repositório -> lista de produtos + id

var sr []Seller = []Seller{}

var lastID int

// Este repositório é uma interface, portanto tem alguns métodos
type Repository interface {
	GetAll() ([]Seller, error)
	Store(id int, cid int, company_name string, address string, telephone int) (Seller, error)
	LastID() (int, error)
}

// Criamos uma struct repository que irá implementar a interface

type repository struct{}

func (repository) GetAll() ([]Seller, error) {
	return sr, nil
}

func (repository) LastID() (int, error) {
	return lastID, nil
}

func (repository) Store(id int, cid int, company_name string, address string, telephone int) (Seller, error) {
	p := Seller{id, cid, company_name, address, telephone}
	sr = append(sr, p)
	lastID = p.ID
	return p, nil
}

// Método criado para instanciar o novo repositório

func NewRepository() Repository {
	return &repository{}
}
