package buyers

type Service interface {
	GetAll() ([]Buyer, error)
	Store(cardNumberId, firstName, lastName string) (Buyer, error)
	Update(id int, cardNumberId, firstName, lastName string) (Buyer, error)
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

func (s service) GetAll() ([]Buyer, error) {
	buyersList, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return buyersList, nil
}

func (s service) Store(cardNumberId, firstName, lastName string) (Buyer, error) {

	lastID, err := s.repository.LastID()

	if err != nil {
		return Buyer{}, err
	}

	lastID++

	buyer, err := s.repository.Store(lastID, cardNumberId, firstName, lastName)

	if err != nil {
		return Buyer{}, err
	}

	return buyer, nil

}

func (s service) Update(id int, cardNumberId, firstName, lastName string) (Buyer, error) {
	buyer, err := s.repository.Update(id, cardNumberId, firstName, lastName)
	if err != nil {
		return Buyer{}, err
	}
	return buyer, err
}

func (s service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
