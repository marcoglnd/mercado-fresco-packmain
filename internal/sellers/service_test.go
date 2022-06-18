package sellers_test

import (
	"errors"
	"testing"

	. "github.com/marcoglnd/mercado-fresco-packmain/internal/sellers"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/sellers/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
)

func createRandomSeller() (listOfSellers []Seller) {
	listOfSellers = []Seller{
		{
			ID:           1,
			Cid:          utils.RandomCode(),
			Company_name: utils.RandomCategory(),
			Address:      utils.RandomCategory(),
			Telephone:    utils.RandomCategory(),
		},
		{
			ID:           2,
			Cid:          utils.RandomCode(),
			Company_name: utils.RandomCategory(),
			Address:      utils.RandomCategory(),
			Telephone:    utils.RandomCategory(),
		},
	}
	return
}

func TestGetAll(t *testing.T) {
	mock := new(mocks.Repository)

	sellers := createRandomSeller()

	t.Run("GetAll in case of success", func(t *testing.T) {
		mock.On("GetAll").Return(sellers, nil).Once()

		service := NewService(mock)

		list, err := service.GetAll()

		assert.NoError(t, err)
		assert.NotEmpty(t, list)

		for i := 0; i < len(list); i++ {
			assert.Equal(t, sellers[i].ID, list[i].ID)
			assert.Equal(t, sellers[i].Cid, list[i].Cid)
			assert.Equal(t, sellers[i].Company_name, list[i].Company_name)
			assert.Equal(t, sellers[i].Address, list[i].Address)
			assert.Equal(t, sellers[i].Telephone, list[i].Telephone)
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
