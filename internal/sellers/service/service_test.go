package service

import (
	"context"
	"errors"
	"testing"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/sellers/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sellers/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// func createRandomSeller() (seller domain.Seller) {
// 	seller = domain.Seller{
// 		ID:           1,
// 		Cid:          utils.RandomCode(),
// 		Company_name: utils.RandomCategory(),
// 		Address:      utils.RandomCategory(),
// 		Telephone:    utils.RandomCategory(),
// 	}
// 	return
// }

// func createRandomListSeller() (listOfSellers []domain.Seller) {

// 	for i := 1; i <= 5; i++ {
// 		seller := createRandomSeller()
// 		seller.ID = i
// 		listOfSellers = append(listOfSellers, seller)
// 	}
// 	return
// }

// func TestGetAll(t *testing.T) {
// 	mock := new(mocks.Repository)

// 	sellersArg := createRandomListSeller()

// 	t.Run("GetAll in case of success", func(t *testing.T) {
// 		mock.On("GetAll").Return(sellersArg, nil).Once()

// 		service := NewService(mock)

// 		list, err := service.GetAll()

// 		assert.NoError(t, err)
// 		assert.NotEmpty(t, list)

// 		for i := 0; i < len(sellersArg); i++ {
// 			assert.Equal(t, sellersArg[i].ID, list[i].ID)
// 			assert.Equal(t, sellersArg[i].Cid, list[i].Cid)
// 			assert.Equal(t, sellersArg[i].Company_name, list[i].Company_name)
// 			assert.Equal(t, sellersArg[i].Address, list[i].Address)
// 			assert.Equal(t, sellersArg[i].Telephone, list[i].Telephone)
// 		}

// 		mock.AssertExpectations(t)
// 	})

// 	t.Run("GetAll in case of error", func(t *testing.T) {
// 		mock.On("GetAll").Return(nil, errors.New("failed to retrieve sellers")).Once()

// 		service := NewService(mock)

// 		list, err := service.GetAll()

// 		assert.Error(t, err)
// 		assert.Empty(t, list)

// 		mock.AssertExpectations(t)
// 	})
// }

// func TestGetById(t *testing.T) {
// 	mock := new(mocks.Repository)

// 	sellerArg := createRandomSeller()

// 	t.Run("GetById in case of success", func(t *testing.T) {
// 		mock.On("GetById", sellerArg.ID).Return(sellerArg, nil).Once()

// 		service := NewService(mock)

// 		seller, err := service.GetById(sellerArg.ID)

// 		assert.NoError(t, err)
// 		assert.NotEmpty(t, seller)

// 		assert.Equal(t, sellerArg.ID, seller.ID)
// 		assert.Equal(t, sellerArg.Cid, seller.Cid)
// 		assert.Equal(t, sellerArg.Company_name, seller.Company_name)
// 		assert.Equal(t, sellerArg.Address, seller.Address)
// 		assert.Equal(t, sellerArg.Telephone, seller.Telephone)

// 		mock.AssertExpectations(t)

// 	})

// 	t.Run("GetById in case of error", func(t *testing.T) {
// 		mock.On("GetById", 185).Return(domain.Seller{}, errors.New("failed to retrieve seller")).Once()

// 		service := NewService(mock)

// 		seller, err := service.GetById(185)

// 		assert.Error(t, err)
// 		assert.Empty(t, seller)

// 		mock.AssertExpectations(t)
// 	})
// }

func TestCreate(t *testing.T) {
	// chamada do mock
	mockSellersRepo := new(mocks.SellerRepository)
	// chamada do que o mock retorna
	mockSeller := &domain.Seller{
		ID:           1,
		Cid:          123,
		Company_name: "Meli",
		Address:      "Avenida Naçoes Unidas",
		Telephone:    "12345678",
	}

	t.Run("it should create a new seller", func(t *testing.T) {
		mockSellersRepo.On("Create",
			mock.Anything,
			mock.Anything,
		).Return(mockSeller, nil).Once()

		service := NewService(mockSellersRepo)

		sel, err := service.Create(context.TODO(), mockSeller)

		expectedCompany_name := "Meli"

		assert.NoError(t, err)
		assert.Equal(t, expectedCompany_name, sel.Company_name)

		mockSellersRepo.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockSellersRepo.On("Create",
			mock.Anything,
			mock.Anything,
		).Return(&domain.Seller{}, errors.New("failed to create")).Once()

		service := NewService(mockSellersRepo)

		_, err := service.Create(context.TODO(), mockSeller)

		assert.Error(t, err)

		mockSellersRepo.AssertExpectations(t)
	})

	// func TestCreate(t *testing.T) {
	// 	mock := new(mocks.SellerRepository)

	// 	t.Run("Verify seller's ID increases when a new seller is created", func(t *testing.T) {

	// 		sellersArg := createRandomListSeller()

	// 		for _, seller := range sellersArg {
	// 			mock.On("Create", seller.Cid, seller.Company_name, seller.Address, seller.Telephone).Return(seller, nil).Once()
	// 		}

	// 		service := NewService(mock)

	// 		var list []domain.Seller

	// 		for _, sellerArg := range sellersArg {
	// 			newSeller, err := service.Create(sellerArg.Cid, sellerArg.Company_name, sellerArg.Address, sellerArg.Telephone)

	// 			assert.NoError(t, err)
	// 			assert.NotEmpty(t, newSeller)

	// 			assert.Equal(t, sellerArg.ID, newSeller.ID)
	// 			assert.Equal(t, sellerArg.Cid, newSeller.Cid)
	// 			assert.Equal(t, sellerArg.Company_name, newSeller.Company_name)
	// 			assert.Equal(t, sellerArg.Address, newSeller.Address)
	// 			assert.Equal(t, sellerArg.Telephone, newSeller.Telephone)
	// 			list = append(list, newSeller)
	// 		}
	// 		assert.True(t, list[0].ID == list[1].ID-1)

	// 		mock.AssertExpectations(t)
	// 	})

	// 	t.Run("Verify when a CID`s seller already exists thrown an error", func(t *testing.T) {
	// 		seller1 := createRandomSeller()
	// 		seller2 := createRandomSeller()

	// 		seller2.Cid = seller1.Cid

	// 		expectedError := errors.New("cid already used")

	// 		mock.On("Create",
	// 			seller1.Cid, seller1.Company_name, seller1.Address, seller1.Telephone,
	// 		).Return(seller1, nil).Once()
	// 		mock.On("Create",
	// 			seller2.Cid, seller2.Company_name, seller2.Address, seller2.Telephone,
	// 		).Return(domain.Seller{}, expectedError).Once()

	// 		s := NewService(mock)
	// 		newSeller1, err := s.Create(seller1.Cid, seller1.Company_name, seller1.Address, seller1.Telephone)
	// 		assert.NoError(t, err)
	// 		assert.NotEmpty(t, newSeller1)

	// 		assert.Equal(t, seller1, newSeller1)

	// 		newSeller2, err := s.Create(seller2.Cid, seller2.Company_name, seller2.Address, seller2.Telephone)
	// 		assert.Error(t, expectedError, err)
	// 		assert.Empty(t, newSeller2)

	// 		assert.NotEqual(t, seller2, newSeller2)
	// 		mock.AssertExpectations(t)

	// 	})
	// }

	// func TestUpdate(t *testing.T) {
	// 	mock := new(mocks.Repository)

	// 	t.Run("Update data in case of success", func(t *testing.T) {
	// 		seller1 := createRandomSeller()
	// 		seller2 := createRandomSeller()

	// 		seller2.ID = seller1.ID

	// 		mock.On("Create",
	// 			seller1.Cid, seller1.Company_name, seller1.Address, seller1.Telephone,
	// 		).Return(seller1, nil).Once()
	// 		mock.On("Update",
	// 			seller1.ID, seller2.Cid, seller2.Company_name, seller2.Address, seller2.Telephone,
	// 		).Return(seller2, nil).Once()

	// 		s := NewService(mock)
	// 		newSeller1, err := s.Create(seller1.Cid, seller1.Company_name, seller1.Address, seller1.Telephone)
	// 		assert.NoError(t, err)
	// 		assert.NotEmpty(t, newSeller1)

	// 		assert.Equal(t, seller1, newSeller1)

	// 		newSeller2, err := s.Update(seller1.ID, seller2.Cid, seller2.Company_name, seller2.Address, seller2.Telephone)
	// 		assert.NoError(t, err)
	// 		assert.NotEmpty(t, newSeller2)

	// 		assert.Equal(t, seller1.ID, newSeller2.ID)
	// 		assert.NotEqual(t, seller1.Cid, newSeller2.Cid)
	// 		assert.NotEqual(t, seller1.Company_name, newSeller2.Company_name)
	// 		assert.NotEqual(t, seller1.Address, newSeller2.Address)
	// 		assert.NotEqual(t, seller1.Telephone, newSeller2.Telephone)

	// 		mock.AssertExpectations(t)
	// 	})

	// 	t.Run("Update throw an error in case of an nonexistent ID", func(t *testing.T) {
	// 		seller := createRandomSeller()

	// 		mock.On("Update", seller.ID, seller.Cid, seller.Company_name, seller.Address, seller.Telephone).Return(domain.Seller{}, errors.New("failed to retrieve seller")).Once()

	// 		service := NewService(mock)

	// 		seller, err := service.Update(seller.ID, seller.Cid, seller.Company_name, seller.Address, seller.Telephone)

	// 		assert.Error(t, err)
	// 		assert.Empty(t, seller)

	// 		mock.AssertExpectations(t)
	// 	})
	// }

	// func TestDelete(t *testing.T) {
	// 	mock := new(mocks.Repository)

	// 	sellerArg := createRandomSeller()

	// 	t.Run("Delete in case of success", func(t *testing.T) {
	// 		mock.On("Create", sellerArg.Cid, sellerArg.Company_name, sellerArg.Address, sellerArg.Telephone).Return(sellerArg, nil).Once()
	// 		mock.On("GetAll").Return([]domain.Seller{sellerArg}, nil).Once()
	// 		mock.On("Delete", sellerArg.ID).Return(nil).Once()
	// 		mock.On("GetAll").Return([]domain.Seller{}, nil).Once()

	// 		service := NewService(mock)

	// 		newSeller, err := service.Create(sellerArg.Cid, sellerArg.Company_name, sellerArg.Address, sellerArg.Telephone)
	// 		assert.NoError(t, err)
	// 		list1, err := service.GetAll()
	// 		assert.NoError(t, err)
	// 		err = service.Delete(newSeller.ID)
	// 		assert.NoError(t, err)
	// 		list2, err := service.GetAll()
	// 		assert.NoError(t, err)

	// 		assert.NotEmpty(t, list1)
	// 		assert.NotEqual(t, list1, list2)
	// 		assert.Empty(t, list2)

	// 		mock.AssertExpectations(t)

	// 	})

	// 	t.Run("Delete in case of error", func(t *testing.T) {
	// 		mock.On("Delete", 185).Return(errors.New("seller's ID not founded")).Once()

	// 		service := NewService(mock)

	// 		err := service.Delete(185)

	// 		assert.Error(t, err)

	// 		mock.AssertExpectations(t)
	// 	})
}
