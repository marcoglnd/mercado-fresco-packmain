package sellers

type Service interface {
	GetAll() ([]Seller, error)
	Store(cid int, company_name string, address string, telephone int) (Seller, error)
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

func (s service) Store(cid int, company_name string, address string, telephone int) (Seller, error) {
	lastID, err := s.repository.LastID()

	if err != nil {
		return Seller{}, err
	}

	lastID++

	seller, err := s.repository.Store(lastID, cid, company_name, address, telephone)

	if err != nil {
		return Seller{}, err
	}

	return seller, nil

}

// Recebe a interface como parâmetro

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
