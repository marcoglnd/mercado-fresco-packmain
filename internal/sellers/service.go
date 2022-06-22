package sellers

type Service interface {
	GetAll() ([]Seller, error)
	GetById(id int) (Seller, error)
	Create(cid int, company_name string, address string, telephone string) (Seller, error)
	Update(id int, cid int, company_name string, address string, telephone string) (Seller, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func (s service) GetAll() ([]Seller, error) {
	sr, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return sr, nil
}

func (s service) GetById(id int) (Seller, error) {
	sr, err := s.repository.GetById(id)
	if err != nil {
		return Seller{}, err
	}
	return sr, nil
}

func (s service) Create(cid int, company_name string, address string, telephone string) (Seller, error) {

	seller, err := s.repository.Create(cid, company_name, address, telephone)

	if err != nil {
		return Seller{}, err
	}

	return seller, nil

}

func (s service) Update(id int, cid int, company_name string, address string, telephone string) (Seller, error) {
	seller, err := s.repository.Update(id, cid, company_name, address, telephone)
	if err != nil {
		return Seller{}, err
	}
	return seller, err
}

func (s service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}

// Recebe a interface como parâmetro

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
