package mariadb

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
)

var (
	queryInsert = regexp.QuoteMeta(sqlInsert)
)

func TestCreatePurchaseOrder(t *testing.T) {
	mockPurchaseOrder := utils.CreateRandomPurchaseOrder()

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsert).
			WithArgs(
				mockPurchaseOrder.OrderNumber,
				mockPurchaseOrder.OrderDate,
				mockPurchaseOrder.TrackingCode,
				mockPurchaseOrder.BuyerId,
				mockPurchaseOrder.CarrierId,
				mockPurchaseOrder.OrderStatusId,
				mockPurchaseOrder.WarehouseId,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewMariaDBRepository(db)

		sec, err := repo.Create(
			context.Background(),
			mockPurchaseOrder.OrderNumber,
			mockPurchaseOrder.OrderDate,
			mockPurchaseOrder.TrackingCode,
			mockPurchaseOrder.BuyerId,
			mockPurchaseOrder.CarrierId,
			mockPurchaseOrder.OrderStatusId,
			mockPurchaseOrder.WarehouseId,
		)
		assert.NoError(t, err)

		assert.Equal(t, &mockPurchaseOrder, sec)
	})

	t.Run("failed to create purchase order", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(queryInsert).
			WithArgs(0, 0, 0, 0, 0, 0, 0).
			WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewMariaDBRepository(db)
		_, err = repo.Create(
			context.Background(),
			mockPurchaseOrder.OrderNumber,
			mockPurchaseOrder.OrderDate,
			mockPurchaseOrder.TrackingCode,
			mockPurchaseOrder.BuyerId,
			mockPurchaseOrder.CarrierId,
			mockPurchaseOrder.OrderStatusId,
			mockPurchaseOrder.WarehouseId,
		)

		assert.Error(t, err)
	})
}
