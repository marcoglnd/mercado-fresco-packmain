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

var rowsEmployeeStruct = []string{
	"id",
	"card_number_id",
	"first_name",
	"last_name",
	"warehouse_id",
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
