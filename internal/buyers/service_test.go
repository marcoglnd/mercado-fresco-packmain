package buyers_test

import (
	"errors"
	"testing"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/buyers/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
)

func createRandomBuyer() (buyer buyers.Buyer) {
	buyer = buyers.Buyer{
		ID:           1,
		CardNumberID: utils.RandomString(6),
		FirstName:    utils.RandomCategory(),
		LastName:     utils.RandomCategory(),
	}
	return
}

func createRandomBuyerList() (listOfBuyers []buyers.Buyer) {

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

		service := mocks.NewService(mock)

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
		mock.On("GetAll").Return(nil, errors.New("failed to retrieve sellers")).Once()

		service := NewService(mock)

		list, err := service.GetAll()

		assert.Error(t, err)
		assert.Empty(t, list)

		mock.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	mock := new(mocks.Repository)

	sellerArg := createRandomSeller()

	t.Run("GetById in case of success", func(t *testing.T) {
		mock.On("GetById", sellerArg.ID).Return(sellerArg, nil).Once()

		service := NewService(mock)

		seller, err := service.GetById(sellerArg.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, seller)

		assert.Equal(t, sellerArg.ID, seller.ID)
		assert.Equal(t, sellerArg.Cid, seller.Cid)
		assert.Equal(t, sellerArg.Company_name, seller.Company_name)
		assert.Equal(t, sellerArg.Address, seller.Address)
		assert.Equal(t, sellerArg.Telephone, seller.Telephone)

		mock.AssertExpectations(t)

	})

	t.Run("GetById in case of error", func(t *testing.T) {
		mock.On("GetById", 185).Return(Seller{}, errors.New("failed to retrieve seller")).Once()

		service := NewService(mock)

		seller, err := service.GetById(185)

		assert.Error(t, err)
		assert.Empty(t, seller)

		mock.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	mock := new(mocks.Repository)

	t.Run("Verify seller's ID increases when a new seller is created", func(t *testing.T) {

		sellersArg := createRandomListSeller()

		for _, seller := range sellersArg {
			mock.On("Create", seller.Cid, seller.Company_name, seller.Address, seller.Telephone).Return(seller, nil).Once()
		}

		service := NewService(mock)

		var list []Seller

		for _, sellerArg := range sellersArg {
			newSeller, err := service.Create(sellerArg.Cid, sellerArg.Company_name, sellerArg.Address, sellerArg.Telephone)

			assert.NoError(t, err)
			assert.NotEmpty(t, newSeller)

			assert.Equal(t, sellerArg.ID, newSeller.ID)
			assert.Equal(t, sellerArg.Cid, newSeller.Cid)
			assert.Equal(t, sellerArg.Company_name, newSeller.Company_name)
			assert.Equal(t, sellerArg.Address, newSeller.Address)
			assert.Equal(t, sellerArg.Telephone, newSeller.Telephone)
			list = append(list, newSeller)
		}
		assert.True(t, list[0].ID == list[1].ID-1)

		mock.AssertExpectations(t)
	})

	t.Run("Verify when a CID`s seller already exists thrown an error", func(t *testing.T) {
		seller1 := createRandomSeller()
		seller2 := createRandomSeller()

		seller2.Cid = seller1.Cid

		expectedError := errors.New("cid already used")

		mock.On("Create",
			seller1.Cid, seller1.Company_name, seller1.Address, seller1.Telephone,
		).Return(seller1, nil).Once()
		mock.On("Create",
			seller2.Cid, seller2.Company_name, seller2.Address, seller2.Telephone,
		).Return(Seller{}, expectedError).Once()

		s := NewService(mock)
		newSeller1, err := s.Create(seller1.Cid, seller1.Company_name, seller1.Address, seller1.Telephone)
		assert.NoError(t, err)
		assert.NotEmpty(t, newSeller1)

		assert.Equal(t, seller1, newSeller1)

		newSeller2, err := s.Create(seller2.Cid, seller2.Company_name, seller2.Address, seller2.Telephone)
		assert.Error(t, expectedError, err)
		assert.Empty(t, newSeller2)

		assert.NotEqual(t, seller2, newSeller2)
		mock.AssertExpectations(t)

	})
}

func TestUpdate(t *testing.T) {
	mock := new(mocks.Repository)

	t.Run("Update data in case of success", func(t *testing.T) {
		seller1 := createRandomSeller()
		seller2 := createRandomSeller()

		seller2.ID = seller1.ID

		mock.On("Create",
			seller1.Cid, seller1.Company_name, seller1.Address, seller1.Telephone,
		).Return(seller1, nil).Once()
		mock.On("Update",
			seller1.ID, seller2.Cid, seller2.Company_name, seller2.Address, seller2.Telephone,
		).Return(seller2, nil).Once()

		s := NewService(mock)
		newSeller1, err := s.Create(seller1.Cid, seller1.Company_name, seller1.Address, seller1.Telephone)
		assert.NoError(t, err)
		assert.NotEmpty(t, newSeller1)

		assert.Equal(t, seller1, newSeller1)

		newSeller2, err := s.Update(seller1.ID, seller2.Cid, seller2.Company_name, seller2.Address, seller2.Telephone)
		assert.NoError(t, err)
		assert.NotEmpty(t, newSeller2)

		assert.Equal(t, seller1.ID, newSeller2.ID)
		assert.NotEqual(t, seller1.Cid, newSeller2.Cid)
		assert.NotEqual(t, seller1.Company_name, newSeller2.Company_name)
		assert.NotEqual(t, seller1.Address, newSeller2.Address)
		assert.NotEqual(t, seller1.Telephone, newSeller2.Telephone)

		mock.AssertExpectations(t)
	})

	t.Run("Update throw an error in case of an nonexistent ID", func(t *testing.T) {
		seller := createRandomSeller()

		mock.On("Update", seller.ID, seller.Cid, seller.Company_name, seller.Address, seller.Telephone).Return(Seller{}, errors.New("failed to retrieve seller")).Once()

		service := NewService(mock)

		seller, err := service.Update(seller.ID, seller.Cid, seller.Company_name, seller.Address, seller.Telephone)

		assert.Error(t, err)
		assert.Empty(t, seller)

		mock.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mock := new(mocks.Repository)

	sellerArg := createRandomSeller()

	t.Run("Delete in case of success", func(t *testing.T) {
		mock.On("Create", sellerArg.Cid, sellerArg.Company_name, sellerArg.Address, sellerArg.Telephone).Return(sellerArg, nil).Once()
		mock.On("GetAll").Return([]Seller{sellerArg}, nil).Once()
		mock.On("Delete", sellerArg.ID).Return(nil).Once()
		mock.On("GetAll").Return([]Seller{}, nil).Once()

		service := NewService(mock)

		newSeller, err := service.Create(sellerArg.Cid, sellerArg.Company_name, sellerArg.Address, sellerArg.Telephone)
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
		mock.On("Delete", 185).Return(errors.New("seller's ID not founded")).Once()

		service := NewService(mock)

		err := service.Delete(185)

		assert.Error(t, err)

		mock.AssertExpectations(t)
	})
}
