package service

import (
	"context"
	"errors"
	"testing"

	. "github.com/marcoglnd/mercado-fresco-packmain/internal/purchase-orders/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/purchase-orders/mocks"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePurchaseOrder(t *testing.T) {
	mockPurchaseOrderRepo := mocks.NewPurchaseOrderRepository(t)
	mockPurchaseOrder := utils.CreateRandomPurchaseOrder()

	t.Run("In case of success", func(t *testing.T) {
		mockPurchaseOrderRepo.On("Create",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(&mockPurchaseOrder, nil).Once()
		mockPurchaseOrderRepo.On("GetByOrderNumber", mock.Anything, mock.Anything).Return(nil, nil)

		s := NewPurchaseOrderService(mockPurchaseOrderRepo)

		newPurchaseOrder, err := s.Create(
			context.Background(),
			mockPurchaseOrder.OrderDate,
			mockPurchaseOrder.OrderNumber,
			mockPurchaseOrder.TrackingCode,
			mockPurchaseOrder.BuyerId,
			mockPurchaseOrder.CarrierId,
			mockPurchaseOrder.OrderStatusId,
			mockPurchaseOrder.WarehouseId,
		)

		assert.NoError(t, err)
		assert.Equal(t, &mockPurchaseOrder, newPurchaseOrder)

		mockPurchaseOrderRepo.AssertExpectations(t)
	})

	t.Run("In case of error", func(t *testing.T) {
		mockPurchaseOrderRepo.On("Create",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(&PurchaseOrder{}, errors.New("failed to create buyer")).Once()

		s := NewPurchaseOrderService(mockPurchaseOrderRepo)

		_, err := s.Create(
			context.Background(),
			mockPurchaseOrder.OrderDate,
			mockPurchaseOrder.OrderNumber,
			mockPurchaseOrder.TrackingCode,
			mockPurchaseOrder.BuyerId,
			mockPurchaseOrder.CarrierId,
			mockPurchaseOrder.OrderStatusId,
			mockPurchaseOrder.WarehouseId,
		)

		assert.Error(t, err)

		mockPurchaseOrderRepo.AssertExpectations(t)
	})
}
