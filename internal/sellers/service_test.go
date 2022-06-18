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
