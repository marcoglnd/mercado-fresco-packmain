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

// 		mock.On("Update", seller.ID, seller.Cid, seller.Company_name, seller.Address, seller.Telephone).Return(Seller{}, errors.New("failed to retrieve seller")).Once()

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
// 		mock.On("GetAll").Return([]Seller{sellerArg}, nil).Once()
// 		mock.On("Delete", sellerArg.ID).Return(nil).Once()
// 		mock.On("GetAll").Return([]Seller{}, nil).Once()

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
// }
