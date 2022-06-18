package sellers_test

import (
	"errors"
	"testing"

	. "github.com/marcoglnd/mercado-fresco-packmain/internal/sellers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sellers/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
)

func createRandomSeller() (seller Seller) {
	seller = Seller{
		ID:           1,
		Cid:          utils.RandomCode(),
		Company_name: utils.RandomCategory(),
		Address:      utils.RandomCategory(),
		Telephone:    utils.RandomCategory(),
	}
	return
}

func createRandomListSeller() (listOfSellers []Seller) {

	for i := 1; i <= 5; i++ {
		seller := createRandomSeller()
		seller.ID = i
		listOfSellers = append(listOfSellers, seller)
	}
	return
}

func TestGetAll(t *testing.T) {
	mock := new(mocks.Repository)

	sellersArg := createRandomListSeller()

	t.Run("GetAll in case of success", func(t *testing.T) {
		mock.On("GetAll").Return(sellersArg, nil).Once()

		service := NewService(mock)

		list, err := service.GetAll()

		assert.NoError(t, err)
		assert.NotEmpty(t, list)

		for i := 0; i < len(sellersArg); i++ {
			assert.Equal(t, sellersArg[i].ID, list[i].ID)
			assert.Equal(t, sellersArg[i].Cid, list[i].Cid)
			assert.Equal(t, sellersArg[i].Company_name, list[i].Company_name)
			assert.Equal(t, sellersArg[i].Address, list[i].Address)
			assert.Equal(t, sellersArg[i].Telephone, list[i].Telephone)
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
