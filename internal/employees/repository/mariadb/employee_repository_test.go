package mariadb

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marcoglnd/mercado-fresco-packmain/internal/employees/domain"
	"github.com/marcoglnd/mercado-fresco-packmain/utils"
	"github.com/stretchr/testify/assert"
)

var rowsEmployeeStruct = []string{
	"id",
	"card_number_id",
	"first_name",
	"last_name",
	"warehouse_id",
}

var rowsInboundOrdersStruct = []string{
	"id",
	"card_number_id",
	"first_name",
	"last_name",
	"warehouse_id",
	"inbound_orders_count",
}

func TestCreateNewEmployee(t *testing.T) {
	mockEmployee := utils.CreateRandomEmployee()

	t.Run("success to create employee", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).WithArgs(
			mockEmployee.CardNumberId,
			mockEmployee.FirstName,
			mockEmployee.LastName,
			mockEmployee.WarehouseId,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		repository := NewMariaDBRepository(db)
		newEmployee, err := repository.Create(context.Background(), &mockEmployee)

		assert.NoError(t, err)
		assert.Equal(t, &mockEmployee, newEmployee)
	})

	t.Run("fail to create employee", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).WithArgs(
			"123",
			"Liz",
			"Souza",
			42,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		repository := NewMariaDBRepository(db)
		_, err = repository.Create(context.Background(), &mockEmployee)

		assert.Error(t, err)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("success to get all employees", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mockEmployees := utils.CreateRandomListEmployees()
		rows := sqlmock.NewRows(rowsEmployeeStruct)

		for _, mockEmployee := range mockEmployees {
			rows.AddRow(
				mockEmployee.ID,
				mockEmployee.CardNumberId,
				mockEmployee.FirstName,
				mockEmployee.LastName,
				mockEmployee.WarehouseId,
			)
		}

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetAll)).WillReturnRows(rows)

		repository := NewMariaDBRepository(db)
		result, err := repository.GetAll(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, result, &mockEmployees)

	})

	t.Run("fail to scan employee", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		rows := sqlmock.NewRows(rowsEmployeeStruct).AddRow("", "", "", "", "")

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetAll)).WillReturnRows(rows)

		repository := NewMariaDBRepository(db)
		_, err = repository.GetAll(context.Background())

		assert.Error(t, err)
	})

	t.Run("fail to select employee", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetAll)).WillReturnError(sql.ErrNoRows)

		repository := NewMariaDBRepository(db)
		_, err = repository.GetAll(context.Background())

		assert.Error(t, err)
	})
}

func TestGetById(t *testing.T) {
	t.Run("success to get employee by id", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mockEmployee := utils.CreateRandomEmployee()

		rows := sqlmock.NewRows(rowsEmployeeStruct).AddRow(
			mockEmployee.ID,
			mockEmployee.CardNumberId,
			mockEmployee.FirstName,
			mockEmployee.LastName,
			mockEmployee.WarehouseId,
		)

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetById)).WillReturnRows(rows)

		repository := NewMariaDBRepository(db)
		result, err := repository.GetById(context.Background(), 0)

		assert.NoError(t, err)
		assert.Equal(t, result, &mockEmployee)
	})

	t.Run("fail to scan employee", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		rows := sqlmock.NewRows(rowsEmployeeStruct).AddRow("", "", "", "", "")

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetById)).WillReturnRows(rows)

		repository := NewMariaDBRepository(db)
		_, err = repository.GetById(context.Background(), 1)

		assert.Error(t, err)
	})

	t.Run("fail to select employee", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(sqlGetById)).WillReturnError(sql.ErrNoRows)

		repository := NewMariaDBRepository(db)
		_, err = repository.GetById(context.Background(), 0)

		assert.Error(t, err)
	})
}

func TestUpdateEmployee(t *testing.T) {
	mockEmployee := utils.CreateRandomEmployee()

	t.Run("success to update employee", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlUpdate)).WithArgs(
			mockEmployee.CardNumberId,
			mockEmployee.FirstName,
			mockEmployee.LastName,
			mockEmployee.WarehouseId,
			mockEmployee.ID,
		).WillReturnResult(sqlmock.NewResult(0, 1))

		repository := NewMariaDBRepository(db)
		result, err := repository.Update(context.Background(), &mockEmployee)

		assert.NoError(t, err)
		assert.Equal(t, &mockEmployee, result)
	})

	t.Run("fail to update employee", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlUpdate)).WithArgs("", "", "", "").
			WillReturnResult(sqlmock.NewResult(0, 1))

		repository := NewMariaDBRepository(db)
		_, err = repository.Update(context.Background(), &mockEmployee)

		assert.Error(t, err)
	})

	t.Run("employee's ID not founded", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlUpdate)).WithArgs(
			mockEmployee.CardNumberId,
			mockEmployee.FirstName,
			mockEmployee.LastName,
			mockEmployee.WarehouseId,
			mockEmployee.ID,
		).WillReturnResult(sqlmock.NewResult(0, 0))

		repository := NewMariaDBRepository(db)
		_, err = repository.Update(context.Background(), &mockEmployee)

		assert.Error(t, err)
		assert.Equal(t, domain.ErrIdNotFound, err)
	})
}

func TestDeleteEmployee(t *testing.T) {
	mockEmployee := utils.CreateRandomEmployee()

	t.Run("success to delete employee", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlDelete)).
			WithArgs(
				mockEmployee.ID,
			).WillReturnResult(sqlmock.NewResult(0, 1))

		repository := NewMariaDBRepository(db)
		err = repository.Delete(context.Background(), mockEmployee.ID)

		assert.NoError(t, err)
	})

	t.Run("fail to delete employee", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlDelete)).
			WithArgs(0).
			WillReturnResult(sqlmock.NewResult(0, 1))

		repository := NewMariaDBRepository(db)
		err = repository.Delete(context.Background(), mockEmployee.ID)

		assert.Error(t, err)
	})

	t.Run("employee's ID not founded", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(sqlDelete)).
			WithArgs(mockEmployee.ID).
			WillReturnResult(sqlmock.NewResult(0, 0))

		repository := NewMariaDBRepository(db)
		err = repository.Delete(context.Background(), mockEmployee.ID)

		assert.Error(t, err)
		assert.Equal(t, domain.ErrIdNotFound, err)
	})
}

func TestReportAllInboundOrders(t *testing.T) {
	t.Run("success to get all inbound orders", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mockInboundOrders := utils.CreateRamdomListReportInboundOrders()
		rows := sqlmock.NewRows(rowsInboundOrdersStruct)

		for _, mockInboundOrders := range mockInboundOrders {
			rows.AddRow(
				mockInboundOrders.ID,
				mockInboundOrders.CardNumberId,
				mockInboundOrders.FirstName,
				mockInboundOrders.LastName,
				mockInboundOrders.WarehouseId,
				mockInboundOrders.InboundOrdersCount,
			)
		}

		mock.ExpectQuery(regexp.QuoteMeta(sqlAllInboundOrdersCount)).WillReturnRows(rows)

		repository := NewMariaDBRepository(db)
		result, err := repository.ReportAllInboundOrders(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, result, &mockInboundOrders)

	})

	t.Run("fail to scan inbound orders", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		rows := sqlmock.NewRows(rowsInboundOrdersStruct).AddRow("", "", "", "", "", "")

		mock.ExpectQuery(regexp.QuoteMeta(sqlAllInboundOrdersCount)).WillReturnRows(rows)

		repository := NewMariaDBRepository(db)
		_, err = repository.ReportAllInboundOrders(context.Background())

		assert.Error(t, err)
	})

	t.Run("fail to select inbound orders", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(sqlAllInboundOrdersCount)).WillReturnError(sql.ErrNoRows)

		repository := NewMariaDBRepository(db)
		_, err = repository.ReportAllInboundOrders(context.Background())

		assert.Error(t, err)
	})
}

func TestReportInboundOrders(t *testing.T) {
	t.Run("success to get inbound orders by employee_id", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		mockInboundOrder := utils.CreateRandomReportInboundOrder()

		rows := sqlmock.NewRows(rowsInboundOrdersStruct).AddRow(
			mockInboundOrder.ID,
			mockInboundOrder.CardNumberId,
			mockInboundOrder.FirstName,
			mockInboundOrder.LastName,
			mockInboundOrder.WarehouseId,
			mockInboundOrder.InboundOrdersCount,
		)

		mock.ExpectQuery(regexp.QuoteMeta(sqlInboundOrdersCountByEmployeeId)).WillReturnRows(rows)

		repository := NewMariaDBRepository(db)
		result, err := repository.ReportInboundOrders(context.Background(), 0)

		assert.NoError(t, err)
		assert.Equal(t, result, &mockInboundOrder)
	})

	t.Run("fail to scan inbound order", func(t *testing.T) {
		db, mock, err := sqlmock.New()

		assert.NoError(t, err)

		defer db.Close()

		rows := sqlmock.NewRows(rowsInboundOrdersStruct).AddRow("", "", "", "", "", "")

		mock.ExpectQuery(regexp.QuoteMeta(sqlInboundOrdersCountByEmployeeId)).WillReturnRows(rows)

		repository := NewMariaDBRepository(db)
		_, err = repository.ReportInboundOrders(context.Background(), 1)

		assert.Error(t, err)
	})
}
