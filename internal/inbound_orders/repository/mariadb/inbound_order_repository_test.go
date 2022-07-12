package mariadb

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
)

var rowsInboundOrderStruct = []string{
	"id",
	"order_date",
	"order_number",
	"employee_id",
	"product_batch_id",
	"warehouse_id",
}

func TestCreateNewInboundOrder(t *testing.T) {
	mockInboundOrder := utils.CreateRandomInboundOrder()

	t.Run("success to create inbound order", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).WithArgs(
			mockInboundOrder.OrderDate,
			mockInboundOrder.OrderNumber,
			mockInboundOrder.EmployeeId,
			mockInboundOrder.ProductBatchId,
			mockInboundOrder.WarehouseId,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		repository := NewMariaDBRepository(db)
		newInboundOrder, err := repository.Create(context.Background(), &mockInboundOrder)

		assert.NoError(t, err)
		assert.Equal(t, &mockInboundOrder, newInboundOrder)
	})

	t.Run("fail to create inbound order", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).WithArgs(
			"10-07-2022",
			"123",
			1,
			1,
			1,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		repository := NewMariaDBRepository(db)
		_, err = repository.Create(context.Background(), &mockInboundOrder)

		assert.Error(t, err)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("success to get all inbound orders", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mockInboundOrders := utils.CreateRandomListInboundOrders()
		rows := sqlmock.NewRows(rowsInboundOrderStruct)

		for _, mockmockInboundOrder := range mockInboundOrders {
			rows.AddRow(
				mockmockInboundOrder.ID,
				mockmockInboundOrder.OrderDate,
				mockmockInboundOrder.OrderNumber,
				mockmockInboundOrder.EmployeeId,
				mockmockInboundOrder.ProductBatchId,
				mockmockInboundOrder.WarehouseId,
			)
		}

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetAll)).WillReturnRows(rows)

		repository := NewMariaDBRepository(db)
		result, err := repository.GetAll(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, result, &mockInboundOrders)
	})

	t.Run("fail to scan inbound order", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		rows := sqlmock.NewRows(rowsInboundOrderStruct).AddRow("", "", "", "", "", "")

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetAll)).WillReturnRows(rows)

		repository := NewMariaDBRepository(db)
		_, err = repository.GetAll(context.Background())

		assert.Error(t, err)
	})

	t.Run("fail to select inbound order", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetAll)).WillReturnError(sql.ErrNoRows)

		repository := NewMariaDBRepository(db)
		_, err = repository.GetAll(context.Background())

		assert.Error(t, err)
	})
}
