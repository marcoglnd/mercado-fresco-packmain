package buyers_test

import (
	"errors"
	"testing"

	. "github.com/marcoglnd/mercado-fresco-packmain/internal/buyers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
)

func createRandomBuyer() (buyer Buyer) {
	buyer = Buyer{
		ID:           1,
		CardNumberID: utils.RandomString(6),
		FirstName:    utils.RandomCategory(),
		LastName:     utils.RandomCategory(),
	}
	return
}

func createRandomBuyerList() (listOfBuyers []Buyer) {

	for i := 1; i <= 5; i++ {
		buyer := createRandomBuyer()
		buyer.ID = i
		listOfBuyers = append(listOfBuyers, buyer)
	}
	return
}

func TestGetAll(t *testing.T) {
	mock := new(mocks.Repository)

	buyersArg := createRandomBuyerList()

	t.Run("GetAll in case of success", func(t *testing.T) {
		mock.On("GetAll").Return(buyersArg, nil).Once()

		service := NewService(mock)

		list, err := service.GetAll()

		assert.NoError(t, err)
		assert.NotEmpty(t, list)

		for i := 0; i < len(buyersArg); i++ {
			assert.Equal(t, buyersArg[i].ID, list[i].ID)
			assert.Equal(t, buyersArg[i].CardNumberID, list[i].CardNumberID)
			assert.Equal(t, buyersArg[i].FirstName, list[i].FirstName)
			assert.Equal(t, buyersArg[i].LastName, list[i].LastName)
		}

		mock.AssertExpectations(t)
	})

	t.Run("GetAll in case of error", func(t *testing.T) {
		mock.On("GetAll").Return(nil, errors.New("failed to retrieve buyers")).Once()

		service := NewService(mock)

		list, err := service.GetAll()

		assert.Error(t, err)
		assert.Empty(t, list)

		mock.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	mock := new(mocks.Repository)

	buyersArg := createRandomBuyer()

	t.Run("GetById in case of success", func(t *testing.T) {
		mock.On("GetById", buyersArg.ID).Return(buyersArg, nil).Once()

		service := NewService(mock)

		buyer, err := service.GetById(buyersArg.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, buyer)

		assert.Equal(t, buyersArg.ID, buyer.ID)
		assert.Equal(t, buyersArg.CardNumberID, buyer.CardNumberID)
		assert.Equal(t, buyersArg.FirstName, buyer.FirstName)
		assert.Equal(t, buyersArg.LastName, buyer.LastName)

		mock.AssertExpectations(t)

	})

	t.Run("GetById in case of error", func(t *testing.T) {
		mock.On("GetById", 185).Return(Buyer{}, errors.New("failed to retrieve buyer")).Once()

		service := NewService(mock)

		buyer, err := service.GetById(185)

		assert.Error(t, err)
		assert.Empty(t, buyer)

		mock.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	mock := new(mocks.Repository)

	t.Run("Verify buyers's ID increases when a new buyer is created", func(t *testing.T) {

		buyersArg := createRandomBuyerList()

		for _, buyer := range buyersArg {
			mock.On("Create", buyer.CardNumberID, buyer.FirstName, buyer.LastName).Return(buyer, nil).Once()
		}

		service := NewService(mock)

		var list []Buyer

		for _, buyerArg := range buyersArg {
			newBuyer, err := service.Create(buyerArg.CardNumberID, buyerArg.FirstName, buyerArg.LastName)

			assert.NoError(t, err)
			assert.NotEmpty(t, newBuyer)

			assert.Equal(t, buyerArg.ID, newBuyer.ID)
			assert.Equal(t, buyerArg.CardNumberID, newBuyer.CardNumberID)
			assert.Equal(t, buyerArg.FirstName, newBuyer.FirstName)
			assert.Equal(t, buyerArg.LastName, newBuyer.LastName)
			list = append(list, newBuyer)
		}
		assert.True(t, list[0].ID == list[1].ID-1)

		mock.AssertExpectations(t)
	})

	t.Run("Verify when buyer's CardNumberID already exists throw an error", func(t *testing.T) {
		buyer1 := createRandomBuyer()
		buyer2 := createRandomBuyer()

		buyer2.CardNumberID = buyer1.CardNumberID

		expectedError := errors.New("CardNumberID already used")

		mock.On("Create",
			buyer1.CardNumberID, buyer1.FirstName, buyer1.LastName,
		).Return(buyer1, nil).Once()
		mock.On("Create",
			buyer2.CardNumberID, buyer2.FirstName, buyer2.LastName,
		).Return(Buyer{}, expectedError).Once()

		s := NewService(mock)
		newBuyer1, err := s.Create(buyer1.CardNumberID, buyer1.FirstName, buyer1.LastName)
		assert.NoError(t, err)
		assert.NotEmpty(t, newBuyer1)

		assert.Equal(t, buyer1, newBuyer1)

		newBuyer2, err := s.Create(buyer2.CardNumberID, buyer2.FirstName, buyer2.LastName)
		assert.Error(t, expectedError, err)
		assert.Empty(t, newBuyer2)

		assert.NotEqual(t, buyer2, newBuyer2)
		mock.AssertExpectations(t)

	})
}

func TestUpdate(t *testing.T) {
	mock := new(mocks.Repository)

	t.Run("Update data in case of success", func(t *testing.T) {
		buyer1 := createRandomBuyer()
		buyer2 := createRandomBuyer()

		buyer2.ID = buyer1.ID

		mock.On("Create",
			buyer1.CardNumberID, buyer1.FirstName, buyer1.LastName,
		).Return(buyer1, nil).Once()
		mock.On("Update",
			buyer1.ID, buyer2.CardNumberID, buyer2.FirstName, buyer2.LastName,
		).Return(buyer2, nil).Once()

		s := NewService(mock)
		newBuyer1, err := s.Create(buyer1.CardNumberID, buyer1.FirstName, buyer1.LastName)
		assert.NoError(t, err)
		assert.NotEmpty(t, newBuyer1)

		assert.Equal(t, buyer1, newBuyer1)

		newBuyer2, err := s.Update(buyer1.ID, buyer2.CardNumberID, buyer2.FirstName, buyer2.LastName)
		assert.NoError(t, err)
		assert.NotEmpty(t, newBuyer2)

		assert.Equal(t, buyer1.ID, newBuyer2.ID)
		assert.NotEqual(t, buyer1.CardNumberID, newBuyer2.CardNumberID)
		assert.NotEqual(t, buyer1.FirstName, newBuyer2.FirstName)
		assert.NotEqual(t, buyer1.LastName, newBuyer2.LastName)

		mock.AssertExpectations(t)
	})

	t.Run("Update throw an error in case of an nonexistent ID", func(t *testing.T) {
		buyer := createRandomBuyer()

		mock.On("Update", buyer.ID, buyer.CardNumberID, buyer.FirstName, buyer.LastName).Return(Buyer{}, errors.New("failed to retrieve buyer")).Once()

		service := NewService(mock)

		buyer, err := service.Update(buyer.ID, buyer.CardNumberID, buyer.FirstName, buyer.LastName)

		assert.Error(t, err)
		assert.Empty(t, buyer)

		mock.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mock := new(mocks.Repository)

	buyerArg := createRandomBuyer()

	t.Run("Delete in case of success", func(t *testing.T) {
		mock.On("Create", buyerArg.CardNumberID, buyerArg.FirstName, buyerArg.LastName).Return(buyerArg, nil).Once()
		mock.On("GetAll").Return([]Buyer{buyerArg}, nil).Once()
		mock.On("Delete", buyerArg.ID).Return(nil).Once()
		mock.On("GetAll").Return([]Buyer{}, nil).Once()

		service := NewService(mock)

		newSeller, err := service.Create(buyerArg.CardNumberID, buyerArg.FirstName, buyerArg.LastName)
		assert.NoError(t, err)
		list1, err := service.GetAll()
		assert.NoError(t, err)
		err = service.Delete(newSeller.ID)
		assert.NoError(t, err)
		list2, err := service.GetAll()
		assert.NoError(t, err)

		assert.NotEmpty(t, list1)
		assert.NotEqual(t, list1, list2)
		assert.Empty(t, list2)

		mock.AssertExpectations(t)

	})

	t.Run("Delete in case of error", func(t *testing.T) {
		mock.On("Delete", 185).Return(errors.New("buyer's ID not found")).Once()

		service := NewService(mock)

		err := service.Delete(185)

		assert.Error(t, err)

		mock.AssertExpectations(t)
	})
}
