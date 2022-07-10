package service

import (
	"context"
	"errors"
	"testing"

	"github.com/marcoglnd/mercado-fresco-packmain/internal/inbound_orders/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/inbound_orders/domain/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNewInboundOrder(t *testing.T) {
	mockInboundOrderRepository := mocks.NewInboundOrderRepository(t)
	mockInboundOrder := utils.CreateRandomInboundOrder()

	t.Run("In case of success", func(t *testing.T) {
		mockInboundOrderRepository.On("Create",
			mock.Anything,
			mock.Anything,
		).Return(&mockInboundOrder, nil).Once()

		service := NewInboundOrderService(mockInboundOrderRepository)
		newInboundOrder, err := service.Create(context.Background(), &mockInboundOrder)

		assert.NoError(t, err)
		assert.Equal(t, &mockInboundOrder, newInboundOrder)

		mockInboundOrderRepository.AssertExpectations(t)

	})

	t.Run("In case of error", func(t *testing.T) {
		mockInboundOrderRepository := mocks.NewInboundOrderRepository(t)
		mockInboundOrder := utils.CreateRandomInboundOrder()

		mockInboundOrderRepository.On("Create",
			mock.Anything,
			mock.Anything,
		).Return(&domain.InboundOrder{}, errors.New("failed to create inbound order")).Once()

		service := NewInboundOrderService(mockInboundOrderRepository)
		_, err := service.Create(context.Background(), &mockInboundOrder)

		assert.Error(t, err)

		mockInboundOrderRepository.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("In case of success", func(t *testing.T) {
		mockInboundOrderRepository := mocks.NewInboundOrderRepository(t)
		mockInboundOrder := utils.CreateRandomListInboundOrders()

		mockInboundOrderRepository.On("GetAll", mock.Anything).Return(&mockInboundOrder, nil).Once()

		service := NewInboundOrderService(mockInboundOrderRepository)
		newInboundOrders, err := service.GetAll(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, &mockInboundOrder, newInboundOrders)

		mockInboundOrderRepository.AssertExpectations(t)

	})

	t.Run("In case of error", func(t *testing.T) {
		mockInboundOrderRepository := mocks.NewInboundOrderRepository(t)

		mockInboundOrderRepository.On("GetAll", mock.Anything).Return(nil, errors.New("failed to retrieve inbound orders")).Once()

		service := NewInboundOrderService(mockInboundOrderRepository)
		_, err := service.GetAll(context.Background())

		assert.NotNil(t, err)

		mockInboundOrderRepository.AssertExpectations(t)

	})
}
